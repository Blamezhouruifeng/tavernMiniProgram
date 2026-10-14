package tavern

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

type sTavernStaff struct{}

func init() {
	service.RegisterTavernStaff(NewTavernStaff())
}

func NewTavernStaff() *sTavernStaff {
	return &sTavernStaff{}
}

// requireStaff 员工身份校验（与客人共用同一套会员体系，凭会员 is_staff 判定）
func (s *sTavernStaff) requireStaff(ctx context.Context) (int64, error) {
	memberId := service.TavernMember().CurrentId(ctx)
	if memberId == 0 {
		return 0, gerror.New("请先登录")
	}
	profile, err := service.TavernMember().Profile(ctx)
	if err != nil {
		return 0, err
	}
	if !profile.IsStaff {
		return 0, gerror.New("当前会员未开通员工身份")
	}
	return memberId, nil
}

// Tables 桌台看板
func (s *sTavernStaff) Tables(ctx context.Context) (res []*tavernin.StaffTableModel, err error) {
	if _, err = s.requireStaff(ctx); err != nil {
		return nil, err
	}

	tableCols := dao.TavernTable.Columns()
	tables, err := dao.TavernTable.Ctx(ctx).
		Where(tableCols.Status, 1).
		OrderAsc(tableCols.Sort).
		OrderAsc(tableCols.Id).
		All()
	if err != nil {
		return nil, err
	}

	sessionCols := dao.TavernTableSession.Columns()
	orderCols := dao.TavernOrder.Columns()

	res = make([]*tavernin.StaffTableModel, 0, len(tables))
	for _, t := range tables {
		tableId := t[tableCols.Id].Int64()
		item := &tavernin.StaffTableModel{
			TableId:    tableId,
			TableName:  t[tableCols.Name].String(),
			SeatStatus: t[tableCols.SeatStatus].Int(),
		}

		session, _ := getOpenSession(ctx, tableId)
		if session.IsEmpty() {
			res = append(res, item)
			continue
		}

		sessionId := session[sessionCols.Id].Int64()
		item.SessionId = sessionId
		item.GuestCount = session[sessionCols.MemberCount].Int()
		if openAt := session[sessionCols.OpenAt].GTime(); openAt != nil {
			item.OpenAt = openAt.Format("Y-m-d H:i:s")
		}

		// 已付金额与订单数
		paidRows, _ := dao.TavernOrder.Ctx(ctx).
			Fields(orderCols.PayAmount).
			Where(orderCols.SessionId, sessionId).
			Where(orderCols.PayStatus, 1).
			All()
		for _, r := range paidRows {
			item.PaidAmount += r[orderCols.PayAmount].Float64()
		}
		item.OrderCount = len(paidRows)

		pending, _ := dao.TavernOrder.Ctx(ctx).
			Where(orderCols.SessionId, sessionId).
			Where(orderCols.PayStatus, 0).
			Count()
		item.PendingPayCount = pending

		// 待上齐数量（行级）
		making, _ := dao.TavernOrderItem.Ctx(ctx).
			InnerJoin("hg_tavern_order o", "o.id = hg_tavern_order_item.order_id").
			Where("o.session_id", sessionId).
			Where("o.pay_status", 1).
			Where("hg_tavern_order_item.maker_status < ?", 2).
			Count()
		item.MakingCount = making

		// 呼叫服务员 / 催单（写入会话备注，员工响应后清除）
		item.Called = gstr.Contains(session[sessionCols.Remark].String(), CallMark)

		res = append(res, item)
	}
	return
}

// Tickets 出票区：已支付且未全部上齐的订单
func (s *sTavernStaff) Tickets(ctx context.Context) (res []*tavernin.OrderModel, err error) {
	if _, err = s.requireStaff(ctx); err != nil {
		return nil, err
	}

	orderCols := dao.TavernOrder.Columns()
	rows, err := dao.TavernOrder.Ctx(ctx).
		Where(orderCols.PayStatus, 1).
		Where("making_status < ?", 3).
		OrderAsc(orderCols.PayTime).
		Limit(50).
		All()
	if err != nil {
		return nil, err
	}

	res = make([]*tavernin.OrderModel, 0, len(rows))
	for _, row := range rows {
		m, err := toOrderModel(ctx, row)
		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return
}

// Serve 行级 / 整单上齐
func (s *sTavernStaff) Serve(ctx context.Context, in *tavernin.StaffServeInp) (res *tavernin.OrderModel, err error) {
	if _, err = s.requireStaff(ctx); err != nil {
		return nil, err
	}

	orderCols := dao.TavernOrder.Columns()
	order, err := dao.TavernOrder.Ctx(ctx).Where(orderCols.OrderNo, in.OrderNo).One()
	if err != nil {
		return nil, err
	}
	if order.IsEmpty() {
		return nil, gerror.New("订单不存在")
	}
	orderId := order[orderCols.Id].Int64()

	itemCols := dao.TavernOrderItem.Columns()
	m := dao.TavernOrderItem.Ctx(ctx).Where(itemCols.OrderId, orderId)
	if len(in.ItemIds) > 0 {
		m = m.WhereIn(itemCols.Id, in.ItemIds)
	}
	_, err = m.Data(g.Map{
		itemCols.MakerStatus: 2,
		itemCols.ServedAt:    gtime.Now(),
		itemCols.UpdatedAt:   gtime.Now(),
	}).Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新上菜状态失败")
	}

	// 整单状态汇总
	total, _ := dao.TavernOrderItem.Ctx(ctx).Where(itemCols.OrderId, orderId).Count()
	served, _ := dao.TavernOrderItem.Ctx(ctx).
		Where(itemCols.OrderId, orderId).
		Where(itemCols.MakerStatus, 2).
		Count()

	makingStatus := 1
	if total > 0 && total == served {
		makingStatus = 2
	}
	_, _ = dao.TavernOrder.Ctx(ctx).WherePri(orderId).Data(g.Map{
		orderCols.MakingStatus: makingStatus,
		orderCols.UpdatedAt:    gtime.Now(),
	}).Update()

	updated, err := dao.TavernOrder.Ctx(ctx).WherePri(orderId).One()
	if err != nil {
		return nil, err
	}
	return toOrderModel(ctx, updated)
}

// Urgent 加急
func (s *sTavernStaff) Urgent(ctx context.Context, in *tavernin.StaffUrgentInp) (err error) {
	if _, err = s.requireStaff(ctx); err != nil {
		return err
	}
	orderCols := dao.TavernOrder.Columns()
	_, err = dao.TavernOrder.Ctx(ctx).
		Where(orderCols.OrderNo, in.OrderNo).
		Data(g.Map{orderCols.Remark: "【加急】", orderCols.UpdatedAt: gtime.Now()}).
		Update()
	return err
}

// ClearGoods 沽清（即时对客人端生效）
func (s *sTavernStaff) ClearGoods(ctx context.Context, in *tavernin.StaffClearGoodsInp) (err error) {
	if _, err = s.requireStaff(ctx); err != nil {
		return err
	}
	goodsCols := dao.TavernGoods.Columns()
	isCleared := 0
	if in.IsCleared {
		isCleared = 1
	}
	_, err = dao.TavernGoods.Ctx(ctx).WherePri(in.GoodsId).Data(g.Map{
		goodsCols.IsCleared: isCleared,
		goodsCols.UpdatedAt: gtime.Now(),
	}).Update()
	return err
}

// VerifyBottle 存酒核销：取酒码 + 店员确认（手机号后四位由店员现场核对）
func (s *sTavernStaff) VerifyBottle(ctx context.Context, in *tavernin.StaffVerifyBottleInp) (res *tavernin.StaffVerifyBottleModel, err error) {
	operatorId, err := s.requireStaff(ctx)
	if err != nil {
		return nil, err
	}

	memberId, depositIds, err := getTakeCodePayload(ctx, in.Code)
	if err != nil {
		return nil, err
	}

	depositId := in.DepositId
	if depositId == 0 && len(depositIds) > 0 {
		depositId = depositIds[0]
	}
	if depositId == 0 {
		return nil, gerror.New("请指定要核销的存酒")
	}

	cols := dao.TavernBottleDeposit.Columns()
	row, err := dao.TavernBottleDeposit.Ctx(ctx).WherePri(depositId).One()
	if err != nil {
		return nil, err
	}
	if row.IsEmpty() {
		return nil, gerror.New("存酒记录不存在")
	}
	if row[cols.Status].Int() != 1 {
		return nil, gerror.New("该存酒已不在库")
	}

	remain := row[cols.RemainRatio].Float64()
	takeRatio := in.TakeRatio
	if takeRatio <= 0 {
		takeRatio = remain
	}
	if takeRatio > remain {
		takeRatio = remain
	}
	newRemain := remain - takeRatio
	if newRemain < 0 {
		newRemain = 0
	}

	status := 1
	if newRemain <= 0 {
		status = 2
	}

	_, err = dao.TavernBottleDeposit.Ctx(ctx).WherePri(depositId).Data(g.Map{
		cols.RemainRatio: newRemain,
		cols.Status:      status,
		cols.UpdatedAt:   gtime.Now(),
		cols.UpdatedBy:   operatorId,
	}).Update()
	if err != nil {
		return nil, gerror.Wrap(err, "核销失败")
	}

	// 取酒流水
	logCols := dao.TavernBottleTakeLog.Columns()
	_, _ = dao.TavernBottleTakeLog.Ctx(ctx).Data(g.Map{
		logCols.StoreId:     1,
		logCols.DepositId:   depositId,
		logCols.MemberId:    row[cols.MemberId].Int64(),
		logCols.TakeType:    1,
		logCols.TakeRatio:   takeRatio,
		logCols.RemainRatio: newRemain,
		logCols.OperatorId:  operatorId,
		logCols.TakeAt:      gtime.Now(),
		logCols.CreatedAt:   gtime.Now(),
	}).Insert()

	updated, err := dao.TavernBottleDeposit.Ctx(ctx).WherePri(depositId).One()
	if err != nil {
		return nil, err
	}

	res = &tavernin.StaffVerifyBottleModel{
		Deposit:   toBottleModel(updated),
		MemberId:  memberId,
		TakeRatio: takeRatio,
	}
	return
}
