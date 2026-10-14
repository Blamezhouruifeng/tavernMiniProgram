package tavern

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
)

// BindPhone 绑定手机号（存酒取酒凭证的前置条件）
// 说明：小程序端应通过 getPhoneNumber 拿到加密数据后由后端解密；
//
//	MVP 先接收明文手机号校验入库，接入微信后替换为解密逻辑。
func (s *sTavernMember) BindPhone(ctx context.Context, in *tavernin.BindPhoneInp) (res *tavernin.MemberInfoModel, err error) {
	memberId := s.CurrentId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}
	if len(in.Phone) != 11 {
		return nil, gerror.New("手机号格式不正确")
	}

	cols := dao.TavernMember.Columns()
	if _, err = dao.TavernMember.Ctx(ctx).WherePri(memberId).Data(g.Map{
		cols.Phone:     in.Phone,
		cols.UpdatedAt: gtime.Now(),
	}).Update(); err != nil {
		return nil, gerror.Wrap(err, "绑定手机号失败")
	}
	return s.ProfileById(ctx, memberId)
}
