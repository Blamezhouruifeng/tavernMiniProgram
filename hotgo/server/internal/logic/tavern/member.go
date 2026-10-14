package tavern

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/token"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

type sTavernMember struct{}

func init() {
	service.RegisterTavernMember(NewTavernMember())
}

func NewTavernMember() *sTavernMember {
	return &sTavernMember{}
}

// CurrentId 当前登录会员ID
func (s *sTavernMember) CurrentId(ctx context.Context) int64 {
	if user := contexts.GetUser(ctx); user != nil {
		return user.Id
	}
	return 0
}

// Login 会员登录
// 说明：MVP 阶段以 code 派生 openid（可通过配置接入微信 code2session）；
// 员工与客人共用同一套会员体系，员工资格由后台在会员上勾选「是员工」决定。
func (s *sTavernMember) Login(ctx context.Context, in *tavernin.MemberLoginInp) (res *tavernin.MemberLoginModel, err error) {
	if in.Code == "" {
		return nil, gerror.New("登录凭证不能为空")
	}

	openid := "dev_" + gmd5.MustEncryptString(in.Code)

	cols := dao.TavernMember.Columns()
	member, err := dao.TavernMember.Ctx(ctx).Where(cols.Openid, openid).One()
	if err != nil {
		return nil, err
	}

	var (
		memberId int64
		nickname string
	)
	if member.IsEmpty() {
		nickname = in.Nickname
		if nickname == "" {
			nickname = "酒友" + gconv.String(time.Now().Unix()%10000)
		}
		memberId, err = dao.TavernMember.Ctx(ctx).Data(g.Map{
			cols.StoreId:     1,
			cols.Openid:      openid,
			cols.Nickname:    nickname,
			cols.Avatar:      in.Avatar,
			cols.Phone:       in.Phone,
			cols.Status:      1,
			cols.LastLoginAt: gtime.Now(),
			cols.CreatedAt:   gtime.Now(),
			cols.UpdatedAt:   gtime.Now(),
		}).InsertAndGetId()
		if err != nil {
			return nil, gerror.Wrap(err, "创建会员失败")
		}
	} else {
		var m *entity.TavernMember
		if err = member.Struct(&m); err != nil {
			return nil, err
		}
		memberId = m.Id
		nickname = m.Nickname
		_, err = dao.TavernMember.Ctx(ctx).WherePri(memberId).Data(g.Map{
			cols.LastLoginAt: gtime.Now(),
			cols.UpdatedAt:   gtime.Now(),
		}).Update()
		if err != nil {
			return nil, err
		}
	}

	identity := &model.Identity{
		Id:       memberId,
		App:      consts.AppApi,
		Username: nickname,
		LoginAt:  gtime.Now(),
	}
	tokenStr, expires, err := token.Login(ctx, identity)
	if err != nil {
		return nil, gerror.Wrap(err, "签发登录凭证失败")
	}

	profile, err := s.ProfileById(ctx, memberId)
	if err != nil {
		return nil, err
	}

	res = &tavernin.MemberLoginModel{
		Token:     tokenStr,
		ExpiresIn: expires,
		Member:    profile,
	}
	return
}

// Profile 当前会员信息
func (s *sTavernMember) Profile(ctx context.Context) (res *tavernin.MemberInfoModel, err error) {
	memberId := s.CurrentId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}
	return s.ProfileById(ctx, memberId)
}

// ProfileById 按ID取会员信息
func (s *sTavernMember) ProfileById(ctx context.Context, memberId int64) (res *tavernin.MemberInfoModel, err error) {
	member, err := dao.TavernMember.Ctx(ctx).WherePri(memberId).One()
	if err != nil {
		return nil, err
	}
	if member.IsEmpty() {
		return nil, gerror.New("会员不存在")
	}

	var m *entity.TavernMember
	if err = member.Struct(&m); err != nil {
		return nil, err
	}

	res = &tavernin.MemberInfoModel{
		Id:        m.Id,
		Nickname:  m.Nickname,
		Avatar:    m.Avatar,
		Phone:     m.Phone,
		IsStaff:   m.IsStaff == 1,
		StaffRole: m.StaffRole,
		Points:    m.Points,
		Balance:   m.Balance,
	}
	return
}

// SwitchStaff 演示用：切换员工身份（真实环境由后台在会员上标记）
func (s *sTavernMember) SwitchStaff(ctx context.Context, in *tavernin.MemberSwitchStaffInp) (res *tavernin.MemberInfoModel, err error) {
	memberId := s.CurrentId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	cols := dao.TavernMember.Columns()
	isStaff := 0
	role := ""
	if in.IsStaff {
		isStaff = 1
		role = in.Role
		if role == "" {
			role = "waiter"
		}
	}

	_, err = dao.TavernMember.Ctx(ctx).WherePri(memberId).Data(g.Map{
		cols.IsStaff:     isStaff,
		cols.StaffRole:   role,
		cols.StaffStatus: isStaff,
		cols.UpdatedAt:   gtime.Now(),
	}).Update()
	if err != nil {
		return nil, err
	}
	return s.ProfileById(ctx, memberId)
}
