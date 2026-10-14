package tavern

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
)

// AbnormalOrders 异常订单：先付制下「钱已收但流程没走完」的单子
func (s *sTavernAdmin) AbnormalOrders(ctx context.Context) (list []*tavernin.AdminAbnormalOrderModel, err error) {
	cols := dao.TavernOrder.Columns()
	list = make([]*tavernin.AdminAbnormalOrderModel, 0, 32)

	// 1) 支付成功但未进入制作（回调丢失 / 出票失败）
	rows, err := dao.TavernOrder.Ctx(ctx).
		Where(cols.PayStatus, 1).
		Where(cols.MakingStatus, 0).
		WhereLT(cols.PayTime, gtime.Now().Add(-1*gtime.S*60)).
		OrderDesc(cols.Id).
		Limit(50).
		All()
	if err != nil {
		return
	}
	for _, row := range rows {
		list = append(list, toAbnormalModel(row, "支付成功但未进入制作，请补出票"))
	}

	// 2) 退款待审核
	rows, err = dao.TavernOrder.Ctx(ctx).
		Where(cols.RefundStatus, 1).
		OrderDesc(cols.Id).
		Limit(50).
		All()
	if err != nil {
		return
	}
	for _, row := range rows {
		list = append(list, toAbnormalModel(row, "退款待审核"))
	}

	// 3) 已超时但未被取消（定时任务可能未启用）
	rows, err = dao.TavernOrder.Ctx(ctx).
		Where(cols.PayStatus, 0).
		WhereLT(cols.ExpireAt, gtime.Now()).
		WhereNotLike(cols.Remark, "%已超时取消%").
		OrderDesc(cols.Id).
		Limit(50).
		All()
	if err != nil {
		return
	}
	for _, row := range rows {
		list = append(list, toAbnormalModel(row, "已超时未取消，请检查定时任务"))
	}

	return
}

func toAbnormalModel(row gdb.Record, issue string) *tavernin.AdminAbnormalOrderModel {
	cols := dao.TavernOrder.Columns()
	return &tavernin.AdminAbnormalOrderModel{
		Id:           row[cols.Id].Int64(),
		OrderNo:      row[cols.OrderNo].String(),
		TableName:    row[cols.TableName].String(),
		MemberName:   row[cols.NicknameSnapshot].String(),
		PayStatus:    row[cols.PayStatus].Int(),
		MakingStatus: row[cols.MakingStatus].Int(),
		RefundStatus: row[cols.RefundStatus].Int(),
		PayAmount:    row[cols.PayAmount].Float64(),
		PayTime:      row[cols.PayTime].GTime(),
		ExpireAt:     row[cols.ExpireAt].GTime(),
		CreatedAt:    row[cols.CreatedAt].GTime(),
		Issue:        issue,
	}
}

// Reissue 手动补出票
func (s *sTavernAdmin) Reissue(ctx context.Context, in *tavernin.AdminReissueInp) (err error) {
	cols := dao.TavernOrder.Columns()
	row, err := dao.TavernOrder.Ctx(ctx).Where(cols.OrderNo, in.OrderNo).One()
	if err != nil {
		return err
	}
	if row.IsEmpty() {
		return gerror.New("订单不存在")
	}
	if row[cols.PayStatus].Int() != 1 {
		return gerror.New("仅已支付订单可补出票")
	}

	orderId := row[cols.Id].Int64()
	itemCols := dao.TavernOrderItem.Columns()
	if _, err = dao.TavernOrderItem.Ctx(ctx).
		Where(itemCols.OrderId, orderId).
		Where(itemCols.MakerStatus, 0).
		Data(g.Map{itemCols.MakerStatus: 1, itemCols.UpdatedAt: gtime.Now()}).
		Update(); err != nil {
		return err
	}
	_, err = dao.TavernOrder.Ctx(ctx).WherePri(orderId).Data(g.Map{
		cols.MakingStatus: 1,
		cols.Remark:       gstr.Trim(row[cols.Remark].String() + " 【后台补出票】"),
		cols.UpdatedAt:    gtime.Now(),
	}).Update()

	g.Log().Infof(ctx, "tavern order reissued orderNo:%s", in.OrderNo)
	return err
}

// Reconcile 支付对账：订单实收 vs 支付流水
func (s *sTavernAdmin) Reconcile(ctx context.Context, in *tavernin.AdminReconcileInp) (res *tavernin.AdminReconcileModel, err error) {
	date := in.Date
	if date == "" {
		date = gtime.Now().Format("Y-m-d")
	}
	start, e := gtime.StrToTime(date + " 00:00:00")
	if e != nil {
		return nil, gerror.New("日期格式应为 YYYY-MM-DD")
	}
	end := start.AddDate(0, 0, 1)

	cols := dao.TavernOrder.Columns()
	orders, err := dao.TavernOrder.Ctx(ctx).
		WhereBetween(cols.CreatedAt, start, end).
		Where(cols.PayStatus, 1).
		OrderAsc(cols.Id).
		All()
	if err != nil {
		return nil, err
	}

	payCols := dao.TavernOrderPay.Columns()
	payRows, err := dao.TavernOrderPay.Ctx(ctx).
		WhereBetween(payCols.CreatedAt, start, end).
		All()
	if err != nil {
		return nil, err
	}
	recordAmount := map[string]float64{}
	for _, r := range payRows {
		recordAmount[r[payCols.OutTradeNo].String()] += r[payCols.Amount].Float64()
	}

	res = &tavernin.AdminReconcileModel{
		Date:  date,
		Items: make([]*tavernin.AdminReconcileItem, 0, len(orders)),
	}

	for _, row := range orders {
		orderNo := row[cols.OrderNo].String()
		amount := row[cols.PayAmount].Float64()
		refund := row[cols.RefundAmount].Float64()

		res.OrderCount++
		res.OrderAmount += amount
		res.RefundAmount += refund

		item := &tavernin.AdminReconcileItem{
			OrderNo:    orderNo,
			PayAmount:  amount,
			PayChannel: row[cols.PayChannel].String(),
			PayTime:    row[cols.PayTime].GTime(),
		}
		if amt, ok := recordAmount[orderNo]; ok {
			item.HasRecord = true
			item.RecordAmt = amt
			res.PayRecordAmount += amt
		}
		res.Items = append(res.Items, item)
	}

	res.Diff = round2(res.OrderAmount - res.PayRecordAmount)
	res.OrderAmount = round2(res.OrderAmount)
	res.PayRecordAmount = round2(res.PayRecordAmount)
	res.RefundAmount = round2(res.RefundAmount)
	return
}

func round2(v float64) float64 {
	return float64(int64(v*100+0.5)) / 100
}
