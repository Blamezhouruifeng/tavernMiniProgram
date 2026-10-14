package tavern

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

type sTavernSession struct{}

func init() {
	service.RegisterTavernSession(NewTavernSession())
}

func NewTavernSession() *sTavernSession {
	return &sTavernSession{}
}

// Open 扫码开台：二次扫码回到原会话，同桌成员共享同一会话
func (s *sTavernSession) Open(ctx context.Context, in *tavernin.SessionOpenInp) (res *tavernin.TableSessionModel, err error) {
	memberId := service.TavernMember().CurrentId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	table, err := getTableByName(ctx, in.TableNo)
	if err != nil {
		return nil, err
	}
	if table.IsEmpty() {
		return nil, gerror.New("桌号不存在，请确认桌码是否正确")
	}
	tableId := table[dao.TavernTable.Columns().Id].Int64()

	exist, err := getOpenSession(ctx, tableId)
	if err != nil {
		return nil, err
	}

	var sessionId int64
	if !exist.IsEmpty() {
		sessionId = exist[dao.TavernTableSession.Columns().Id].Int64()
	} else {
		sessionId, err = dao.TavernTableSession.Ctx(ctx).Data(g.Map{
			dao.TavernTableSession.Columns().StoreId:       1,
			dao.TavernTableSession.Columns().TableId:       tableId,
			dao.TavernTableSession.Columns().SessionNo:     genSessionNo(),
			dao.TavernTableSession.Columns().OwnerMemberId: memberId,
			dao.TavernTableSession.Columns().GuestCount:    0,
			dao.TavernTableSession.Columns().Status:        1,
			dao.TavernTableSession.Columns().OpenAt:        gtime.Now(),
			dao.TavernTableSession.Columns().CreatedAt:     gtime.Now(),
			dao.TavernTableSession.Columns().UpdatedAt:     gtime.Now(),
		}).InsertAndGetId()
		if err != nil {
			return nil, gerror.Wrap(err, "开台失败")
		}
		// 桌位进入用餐中
		_, _ = dao.TavernTable.Ctx(ctx).WherePri(tableId).Data(g.Map{
			dao.TavernTable.Columns().SeatStatus: 2,
			dao.TavernTable.Columns().UpdatedAt:  gtime.Now(),
		}).Update()
	}

	// 加入会话成员（已存在则跳过）
	if err = s.joinMember(ctx, sessionId, memberId); err != nil {
		return nil, err
	}

	row, err := getSessionById(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	return toSessionModel(ctx, row, true)
}

// joinMember 加入会话（幂等）
func (s *sTavernSession) joinMember(ctx context.Context, sessionId, memberId int64) (err error) {
	cols := dao.TavernSessionMember.Columns()
	exist, err := dao.TavernSessionMember.Ctx(ctx).
		Where(cols.SessionId, sessionId).
		Where(cols.MemberId, memberId).
		One()
	if err != nil {
		return
	}
	if !exist.IsEmpty() {
		return
	}

	profile, err := service.TavernMember().Profile(ctx)
	if err != nil {
		return
	}

	// 是否开台人
	session, err := getSessionById(ctx, sessionId)
	if err != nil {
		return
	}
	isOwner := 0
	if !session.IsEmpty() && session[dao.TavernTableSession.Columns().OwnerMemberId].Int64() == memberId {
		isOwner = 1
	}

	_, err = dao.TavernSessionMember.Ctx(ctx).Data(g.Map{
		cols.StoreId:   1,
		cols.SessionId: sessionId,
		cols.MemberId:  memberId,
		cols.Nickname:  profile.Nickname,
		cols.Avatar:    profile.Avatar,
		cols.IsOwner:   isOwner,
		cols.JoinAt:    gtime.Now(),
		cols.CreatedAt: gtime.Now(),
		cols.UpdatedAt: gtime.Now(),
	}).Insert()
	if err != nil {
		return gerror.Wrap(err, "加入桌台失败")
	}

	// 维护会话人数
	count, _ := dao.TavernSessionMember.Ctx(ctx).
		Where(cols.SessionId, sessionId).
		Count()
	_, _ = dao.TavernTableSession.Ctx(ctx).WherePri(sessionId).Data(g.Map{
		dao.TavernTableSession.Columns().MemberCount: count,
		dao.TavernTableSession.Columns().UpdatedAt:   gtime.Now(),
	}).Update()
	return
}

// Snapshot 会话快照：成员 + 购物车 + 订单 + 全桌已点聚合
func (s *sTavernSession) Snapshot(ctx context.Context, sessionId int64) (res *tavernin.SessionSnapshotModel, err error) {
	if sessionId == 0 {
		return nil, gerror.New("请先扫码开台")
	}
	row, err := getSessionById(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	if row.IsEmpty() {
		return nil, gerror.New("桌台会话不存在")
	}

	sessionModel, err := toSessionModel(ctx, row, true)
	if err != nil {
		return nil, err
	}

	cart, err := loadCartItems(ctx, sessionId)
	if err != nil {
		return nil, err
	}

	orders, err := loadOrders(ctx, sessionId, 0, false)
	if err != nil {
		return nil, err
	}

	res = &tavernin.SessionSnapshotModel{
		Session:    sessionModel,
		Members:    sessionModel.Members,
		Cart:       cart,
		Orders:     orders,
		OrderedAgg: buildOrderedAgg(orders),
	}
	return
}

// Close 清台
func (s *sTavernSession) Close(ctx context.Context, in *tavernin.SessionCloseInp) (err error) {
	row, err := getSessionById(ctx, in.SessionId)
	if err != nil {
		return err
	}
	if row.IsEmpty() {
		return gerror.New("桌台会话不存在")
	}

	sessionCols := dao.TavernTableSession.Columns()
	_, err = dao.TavernTableSession.Ctx(ctx).WherePri(in.SessionId).Data(g.Map{
		sessionCols.Status:    2,
		sessionCols.CloseAt:   gtime.Now(),
		sessionCols.UpdatedAt: gtime.Now(),
	}).Update()
	if err != nil {
		return gerror.Wrap(err, "清台失败")
	}

	tableId := row[sessionCols.TableId].Int64()
	_, _ = dao.TavernTable.Ctx(ctx).WherePri(tableId).Data(g.Map{
		dao.TavernTable.Columns().SeatStatus: 1,
		dao.TavernTable.Columns().UpdatedAt:  gtime.Now(),
	}).Update()

	// 清空该会话剩余购物车（已下单的行在提交时已清除）
	_, _ = dao.TavernCartItem.Ctx(ctx).
		Where(dao.TavernCartItem.Columns().SessionId, in.SessionId).
		Delete()

	g.Log().Infof(ctx, "tavern session closed sessionId:%d tableId:%d", in.SessionId, tableId)
	return
}

// CurrentSessionId 校验并返回会话ID
func (s *sTavernSession) CurrentSessionId(ctx context.Context, sessionId int64) (id int64, err error) {
	if sessionId == 0 {
		return 0, gerror.New("请先扫码开台")
	}
	row, err := getSessionById(ctx, sessionId)
	if err != nil {
		return 0, err
	}
	if row.IsEmpty() {
		return 0, gerror.New("桌台会话不存在")
	}
	if row[dao.TavernTableSession.Columns().Status].Int() != 1 {
		return 0, gerror.New("该桌台已结束用餐，请重新扫码")
	}
	return sessionId, nil
}

func genSessionNo() string {
	return "SS" + gtime.Now().Format("YmdHis") + gconv.String(gtime.TimestampMilli()%1000)
}
