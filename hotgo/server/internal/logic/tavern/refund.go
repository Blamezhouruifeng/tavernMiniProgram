package tavern

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

// ApplyRefund 申请退款
// 先付制规则：制作前（无任何行进入制作）可自助通过并原路退回；已进入制作则进入待审核。
func (s *sTavernOrder) ApplyRefund(ctx context.Context, in *tavernin.RefundApplyInp) (res *tavernin.OrderModel, err error) {
	memberId := service.TavernMember().CurrentId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	orderCols := dao.TavernOrder.Columns()
	order, err := dao.TavernOrder.Ctx(ctx).Where(orderCols.OrderNo, in.OrderNo).One()
	if err != nil {
		return nil, err
	}
	if order.IsEmpty() {
		return nil, gerror.New("订单不存在")
	}
	if order[orderCols.MemberId].Int64() != memberId {
		return nil, gerror.New("只能对本人订单申请退款")
	}
	if order[orderCols.PayStatus].Int() != 1 {
		return nil, gerror.New("仅已支付订单可申请退款")
	}
	if order[orderCols.RefundStatus].Int() != 0 {
		return nil, gerror.New("该订单已在退款流程中")
	}

	orderId := order[orderCols.Id].Int64()

	// 计算退款金额（按明细行；未指定则整单退）
	itemCols := dao.TavernOrderItem.Columns()
	m := dao.TavernOrderItem.Ctx(ctx).Where(itemCols.OrderId, orderId)
	if len(in.ItemIds) > 0 {
		m = m.WhereIn(itemCols.Id, in.ItemIds)
	}
	items, err := m.All()
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, gerror.New("没有可退的明细")
	}
	var refundAmount float64
	for _, it := range items {
		refundAmount += it[itemCols.Amount].Float64()
	}

	// 是否还能自助退款：先付制下「支付成功即制作」，因此以支付后的宽限期判定
	// （Q2 决策：宽限期内可自助原路退，超出后进入员工审核）
	payAt := order[orderCols.PayTime].GTime()
	selfService := true // 支付时间缺失时按未制作处理
	if payAt != nil {
		selfService = gtime.Now().Timestamp()-payAt.Timestamp() < int64(selfRefundGrace.Seconds())
	}
	alreadyMaking := !selfService

	data := g.Map{
		orderCols.RefundStatus: 1,
		orderCols.RefundAmount: refundAmount,
		orderCols.RefundReason: in.Reason,
		orderCols.UpdatedAt:    gtime.Now(),
	}
	if !alreadyMaking {
		// 制作前：自助通过
		data[orderCols.RefundStatus] = 4
		data[orderCols.RefundAt] = gtime.Now()
		data[orderCols.PayStatus] = refundStatusToPayStatus(refundAmount, order[orderCols.PayAmount].Float64())
	}

	if _, err = dao.TavernOrder.Ctx(ctx).WherePri(orderId).Data(data).Update(); err != nil {
		return nil, gerror.Wrap(err, "退款申请失败")
	}

	// 退款时退回优惠券
	if rerr := service.TavernCoupon().ReturnCoupon(ctx, orderId); rerr != nil {
		g.Log().Warningf(ctx, "退回优惠券失败 orderId:%d err:%+v", orderId, rerr)
	}

	if !alreadyMaking {
		// 回滚预占库存（幂等）
		if rerr := restoreStockByOrder(ctx, orderId); rerr != nil {
			g.Log().Warningf(ctx, "退款回滚库存失败 orderId:%d err:%+v", orderId, rerr)
		}

		// 原路退回流水（真实环境调用微信退款接口，此处记录流水用于对账）
		payCols := dao.TavernOrderPay.Columns()
		_, _ = dao.TavernOrderPay.Ctx(ctx).Data(g.Map{
			payCols.StoreId:     1,
			payCols.OrderId:     orderId,
			payCols.OutTradeNo:  in.OrderNo + "-R" + gtime.Now().Format("His"),
			payCols.PayChannel:  "refund",
			payCols.Amount:      refundAmount,
			payCols.PayStatus:   1,
			payCols.PaidAt:      gtime.Now(),
			payCols.RawCallback: "self-service refund before making",
			payCols.CreatedAt:   gtime.Now(),
			payCols.UpdatedAt:   gtime.Now(),
		}).Insert()
	}

	updated, err := dao.TavernOrder.Ctx(ctx).WherePri(orderId).One()
	if err != nil {
		return nil, err
	}
	return toOrderModel(ctx, updated)
}

// AuditRefund 退款审核（员工）
func (s *sTavernOrder) AuditRefund(ctx context.Context, in *tavernin.RefundAuditInp) (res *tavernin.OrderModel, err error) {
	memberId := service.TavernMember().CurrentId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	orderCols := dao.TavernOrder.Columns()
	order, err := dao.TavernOrder.Ctx(ctx).Where(orderCols.OrderNo, in.OrderNo).One()
	if err != nil {
		return nil, err
	}
	if order.IsEmpty() {
		return nil, gerror.New("订单不存在")
	}
	if order[orderCols.RefundStatus].Int() != 1 {
		return nil, gerror.New("该订单当前不在待审核状态")
	}

	orderId := order[orderCols.Id].Int64()
	refundAmount := order[orderCols.RefundAmount].Float64()

	data := g.Map{
		orderCols.RefundAuditBy: memberId,
		orderCols.UpdatedAt:     gtime.Now(),
	}
	if in.Approve {
		data[orderCols.RefundStatus] = 4
		data[orderCols.RefundAt] = gtime.Now()
		data[orderCols.PayStatus] = refundStatusToPayStatus(refundAmount, order[orderCols.PayAmount].Float64())
	} else {
		data[orderCols.RefundStatus] = 3
		if in.Reason != "" {
			data[orderCols.RefundReason] = in.Reason
		}
	}

	if _, err = dao.TavernOrder.Ctx(ctx).WherePri(orderId).Data(data).Update(); err != nil {
		return nil, gerror.Wrap(err, "退款审核失败")
	}

	if in.Approve {
		payCols := dao.TavernOrderPay.Columns()
		_, _ = dao.TavernOrderPay.Ctx(ctx).Data(g.Map{
			payCols.StoreId:     1,
			payCols.OrderId:     orderId,
			payCols.OutTradeNo:  in.OrderNo + "-R" + gtime.Now().Format("His"),
			payCols.PayChannel:  "refund",
			payCols.Amount:      refundAmount,
			payCols.PayStatus:   1,
			payCols.PaidAt:      gtime.Now(),
			payCols.RawCallback: g.Map{"auditBy": memberId, "reason": in.Reason},
			payCols.CreatedAt:   gtime.Now(),
			payCols.UpdatedAt:   gtime.Now(),
		}).Insert()
	}

	updated, err := dao.TavernOrder.Ctx(ctx).WherePri(orderId).One()
	if err != nil {
		return nil, err
	}
	return toOrderModel(ctx, updated)
}

// selfRefundGrace 自助退款宽限期（先付制：支付后在此期间内可自助原路退）
const selfRefundGrace = 2 * time.Minute

// refundStatusToPayStatus 退款金额 → 支付状态（2 部分退款 / 3 已退款）
func refundStatusToPayStatus(refundAmount, payAmount float64) int {
	if refundAmount >= payAmount {
		return 3
	}
	return 2
}
