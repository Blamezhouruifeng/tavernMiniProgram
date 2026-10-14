package tavern

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
)

// 通用说明：
//  - 本文件集中放置会话 / 订单装配所需的查询与转换helper，避免各模块重复写 SQL。
//  - 先付制：payStatus=1 的订单才会进入出票与制作；makingStatus 由支付回调推进。

// getTableByName 按桌号取桌位
func getTableByName(ctx context.Context, tableNo string) (gdb.Record, error) {
	cols := dao.TavernTable.Columns()
	return dao.TavernTable.Ctx(ctx).
		Where(cols.Name, tableNo).
		Where(cols.Status, 1).
		One()
}

// getOpenSession 取某桌当前进行中的会话
func getOpenSession(ctx context.Context, tableId int64) (gdb.Record, error) {
	cols := dao.TavernTableSession.Columns()
	return dao.TavernTableSession.Ctx(ctx).
		Where(cols.TableId, tableId).
		Where(cols.Status, 1).
		OrderDesc(cols.Id).
		One()
}

// getSessionById 按ID取会话
func getSessionById(ctx context.Context, sessionId int64) (gdb.Record, error) {
	return dao.TavernTableSession.Ctx(ctx).WherePri(sessionId).One()
}

// loadSessionMembers 会话成员
func loadSessionMembers(ctx context.Context, sessionId int64) (res []*tavernin.SessionMemberModel, err error) {
	cols := dao.TavernSessionMember.Columns()
	rows, err := dao.TavernSessionMember.Ctx(ctx).
		Where(cols.SessionId, sessionId).
		OrderAsc(cols.Id).
		All()
	if err != nil {
		return
	}
	res = make([]*tavernin.SessionMemberModel, 0, len(rows))
	for _, row := range rows {
		res = append(res, &tavernin.SessionMemberModel{
			MemberId: row[cols.MemberId].Int64(),
			Nickname: row[cols.Nickname].String(),
			Avatar:   row[cols.Avatar].String(),
			IsOwner:  row[cols.IsOwner].Int() == 1,
			JoinAt:   row[cols.JoinAt].GTime(),
		})
	}
	return
}

// countSessionOrders 统计会话已支付订单数与金额
func countSessionOrders(ctx context.Context, sessionId int64) (count int, amount float64, err error) {
	cols := dao.TavernOrder.Columns()
	rows, err := dao.TavernOrder.Ctx(ctx).
		Fields(cols.PayAmount).
		Where(cols.SessionId, sessionId).
		Where(cols.PayStatus, 1).
		All()
	if err != nil {
		return
	}
	count = len(rows)
	for _, row := range rows {
		amount += row[cols.PayAmount].Float64()
	}
	return
}

// toSessionModel 会话记录 → 出参模型
func toSessionModel(ctx context.Context, row gdb.Record, withMembers bool) (res *tavernin.TableSessionModel, err error) {
	cols := dao.TavernTableSession.Columns()
	res = &tavernin.TableSessionModel{
		Id:              row[cols.Id].Int64(),
		SessionNo:       row[cols.SessionNo].String(),
		TableId:         row[cols.TableId].Int64(),
		Status:          row[cols.Status].Int(),
		GuestCount:      row[cols.GuestCount].Int(),
		OwnerMemberId:   row[cols.OwnerMemberId].Int64(),
		OpenAt:          row[cols.OpenAt].GTime(),
		TotalPaidAmount: row[cols.TotalPaidAmount].Float64(),
	}

	if table, _ := dao.TavernTable.Ctx(ctx).WherePri(res.TableId).One(); !table.IsEmpty() {
		res.TableName = table[dao.TavernTable.Columns().Name].String()
	}

	count, amount, err := countSessionOrders(ctx, res.Id)
	if err != nil {
		return
	}
	res.OrderCount = count
	res.TotalPaidAmount = amount
	res.Members = []*tavernin.SessionMemberModel{}
	if withMembers {
		res.Members, err = loadSessionMembers(ctx, res.Id)
		if err != nil {
			return
		}
	}
	return
}

// loadOrders 会话订单（含明细）
func loadOrders(ctx context.Context, sessionId int64, memberId int64, onlyPaid bool) (res []*tavernin.OrderModel, err error) {
	cols := dao.TavernOrder.Columns()
	m := dao.TavernOrder.Ctx(ctx)
	if sessionId > 0 {
		m = m.Where(cols.SessionId, sessionId)
	}
	if memberId > 0 {
		m = m.Where(cols.MemberId, memberId)
	}
	if onlyPaid {
		m = m.Where(cols.PayStatus, 1)
	}
	rows, err := m.OrderDesc(cols.Id).Limit(100).All()
	if err != nil {
		return
	}
	res = make([]*tavernin.OrderModel, 0, len(rows))
	for _, row := range rows {
		om, err := toOrderModel(ctx, row)
		if err != nil {
			return nil, err
		}
		res = append(res, om)
	}
	return
}

// toOrderModel 订单记录 → 出参模型（含明细）
func toOrderModel(ctx context.Context, row gdb.Record) (res *tavernin.OrderModel, err error) {
	cols := dao.TavernOrder.Columns()
	res = &tavernin.OrderModel{
		Id:             row[cols.Id].Int64(),
		OrderNo:        row[cols.OrderNo].String(),
		SessionId:      row[cols.SessionId].Int64(),
		TableId:        row[cols.TableId].Int64(),
		TableName:      row[cols.TableName].String(),
		OrderSeq:       row[cols.OrderSeq].Int(),
		MemberId:       row[cols.MemberId].Int64(),
		MemberName:     row[cols.NicknameSnapshot].String(),
		PayerMemberId:  row[cols.PayerMemberId].Int64(),
		PayStatus:      row[cols.PayStatus].Int(),
		MakingStatus:   row[cols.MakingStatus].Int(),
		RefundStatus:   row[cols.RefundStatus].Int(),
		PayChannel:     row[cols.PayChannel].String(),
		TotalAmount:    row[cols.TotalAmount].Float64(),
		PayAmount:      row[cols.PayAmount].Float64(),
		DiscountAmount: row[cols.DiscountAmount].Float64(),
		ExpireAt:       row[cols.ExpireAt].GTime(),
		PayTime:        row[cols.PayTime].GTime(),
		CreatedAt:      row[cols.CreatedAt].GTime(),
	}

	items, err := loadOrderItems(ctx, res.Id)
	if err != nil {
		return nil, err
	}
	res.Items = items
	return
}

// loadOrderItems 订单明细
func loadOrderItems(ctx context.Context, orderId int64) (res []*tavernin.OrderItemModel, err error) {
	cols := dao.TavernOrderItem.Columns()
	rows, err := dao.TavernOrderItem.Ctx(ctx).Where(cols.OrderId, orderId).OrderAsc(cols.Id).All()
	if err != nil {
		return
	}
	res = make([]*tavernin.OrderItemModel, 0, len(rows))
	for _, row := range rows {
		res = append(res, &tavernin.OrderItemModel{
			Id:          row[cols.Id].Int64(),
			OrderId:     row[cols.OrderId].Int64(),
			GoodsId:     row[cols.GoodsId].Int64(),
			SkuId:       row[cols.SkuId].Int64(),
			GoodsName:   row[cols.GoodsName].String(),
			SkuName:     row[cols.SkuName].String(),
			UnitPrice:   row[cols.UnitPrice].Float64(),
			Qty:         row[cols.Qty].Int(),
			Amount:      row[cols.Amount].Float64(),
			IsWine:      row[cols.IsWine].Int() == 1,
			MakerStatus: row[cols.MakerStatus].Int(),
			ServedAt:    row[cols.ServedAt].GTime(),
			Remark:      row[cols.Remark].String(),
		})
	}
	return
}

// buildOrderedAgg 全桌已点聚合（按菜合并，可追溯是谁点的）
func buildOrderedAgg(orders []*tavernin.OrderModel) (res []*tavernin.OrderedAggModel) {
	type key struct {
		goodsId int64
		skuId   int64
	}
	index := map[key]*tavernin.OrderedAggModel{}
	order := make([]key, 0)
	for _, o := range orders {
		for _, it := range o.Items {
			k := key{it.GoodsId, it.SkuId}
			cur, ok := index[k]
			if !ok {
				cur = &tavernin.OrderedAggModel{
					GoodsId:   it.GoodsId,
					SkuId:     it.SkuId,
					GoodsName: it.GoodsName,
					SkuName:   it.SkuName,
					ByMembers: []string{},
				}
				index[k] = cur
				order = append(order, k)
			}
			cur.Qty += it.Qty
			if it.MakerStatus >= 2 {
				cur.ServedQty += it.Qty
			}
			found := false
			for _, n := range cur.ByMembers {
				if n == o.MemberName {
					found = true
					break
				}
			}
			if !found {
				cur.ByMembers = append(cur.ByMembers, o.MemberName)
			}
		}
	}
	res = make([]*tavernin.OrderedAggModel, 0, len(order))
	for _, k := range order {
		res = append(res, index[k])
	}
	return
}

// nowTime 当前时间（统一入口，便于测试替换）
func nowTime() *gtime.Time {
	return gtime.Now()
}
