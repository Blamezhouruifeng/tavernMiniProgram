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

type sTavernOrder struct{}

func init() {
	service.RegisterTavernOrder(NewTavernOrder())
}

func NewTavernOrder() *sTavernOrder {
	return &sTavernOrder{}
}

// Submit 提交订单（先付制：先进待支付，支付成功后才进入制作）
func (s *sTavernOrder) Submit(ctx context.Context, in *tavernin.OrderSubmitInp) (res *tavernin.OrderModel, err error) {
	memberId := service.TavernMember().CurrentId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}
	if _, err = service.TavernSession().CurrentSessionId(ctx, in.SessionId); err != nil {
		return nil, err
	}
	if len(in.CartIds) == 0 {
		return nil, gerror.New("请选择要下单的商品")
	}

	profile, err := service.TavernMember().Profile(ctx)
	if err != nil {
		return nil, err
	}

	// 取自己的购物车行（归属校验：只能下单自己的篮子）
	cartCols := dao.TavernCartItem.Columns()
	rows, err := dao.TavernCartItem.Ctx(ctx).
		Where(cartCols.SessionId, in.SessionId).
		WhereIn(cartCols.Id, in.CartIds).
		WhereNull(cartCols.DeletedAt).
		All()
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, gerror.New("购物车行不存在或已失效")
	}
	for _, row := range rows {
		if row[cartCols.MemberId].Int64() != memberId {
			return nil, gerror.New("只能提交自己添加的商品")
		}
	}

	// 行级校验：沽清 / 下架只让该行失败（不整单回滚）
	goodsCols := dao.TavernGoods.Columns()
	skuCols := dao.TavernGoodsSku.Columns()
	type line struct {
		goodsId   int64
		skuId     int64
		goodsName string
		skuName   string
		price     float64
		qty       int
		isWine    bool
		remark    string
		rowId     int64
	}
	lines := make([]*line, 0, len(rows))
	var total float64
	for _, row := range rows {
		goodsId := row[cartCols.GoodsId].Int64()
		skuId := row[cartCols.SkuId].Int64()

		goods, err := dao.TavernGoods.Ctx(ctx).WherePri(goodsId).One()
		if err != nil {
			return nil, err
		}
		if goods.IsEmpty() || goods[goodsCols.Status].Int() != 1 {
			return nil, gerror.New("部分商品已下架，请从购物车中移除后重试")
		}
		if goods[goodsCols.IsCleared].Int() == 1 {
			return nil, gerror.New(goods[goodsCols.Name].String() + " 已沽清，请从购物车中移除后再下单")
		}

		sku, err := dao.TavernGoodsSku.Ctx(ctx).WherePri(skuId).One()
		if err != nil {
			return nil, err
		}
		if sku.IsEmpty() {
			return nil, gerror.New("部分规格已失效，请重新选择")
		}

		qty := row[cartCols.Qty].Int()

		// 行级预占库存（限量规格）：扣不到即报该行缺货，避免超卖
		if err = deductStock(ctx, skuId, qty); err != nil {
			return nil, err
		}

		price := sku[skuCols.Price].Float64()
		total += price * float64(qty)
		lines = append(lines, &line{
			goodsId:   goodsId,
			skuId:     skuId,
			goodsName: goods[goodsCols.Name].String(),
			skuName:   sku[skuCols.Name].String(),
			price:     price,
			qty:       qty,
			isWine:    goods[goodsCols.IsWine].Int() == 1,
			remark:    row[cartCols.Remark].String(),
			rowId:     row[cartCols.Id].Int64(),
		})
	}

	// 桌位信息
	session, err := getSessionById(ctx, in.SessionId)
	if err != nil {
		return nil, err
	}
	tableId := session[dao.TavernTableSession.Columns().TableId].Int64()
	tableName := ""
	if table, _ := dao.TavernTable.Ctx(ctx).WherePri(tableId).One(); !table.IsEmpty() {
		tableName = table[dao.TavernTable.Columns().Name].String()
	}

	// 本桌第几单
	orderCols := dao.TavernOrder.Columns()
	orderSeq, _ := dao.TavernOrder.Ctx(ctx).
		Where(orderCols.SessionId, in.SessionId).
		Count()
	orderSeq = orderSeq + 1

	// 待支付超时（分钟）来自门店配置
	expireMinutes := 5
	if storeInfo, e := service.TavernStore().Info(ctx); e == nil && storeInfo.PayExpireMinutes > 0 {
		expireMinutes = storeInfo.PayExpireMinutes
	}

	// 待支付超时时间（先付制：超时自动取消）
	expireAt := gtime.New(gtime.Now().Unix() + int64(expireMinutes*60))
	// 优惠券：先试算，再落库（支付成功时才标记已使用）
	var discount float64
	if in.CouponId > 0 {
		discount, err = service.TavernCoupon().CalcDiscount(ctx, memberId, in.CouponId, total)
		if err != nil {
			return nil, err
		}
		if discount > total {
			discount = total
		}
	}
	payAmount := round2(total - discount)
	if payAmount < 0 {
		payAmount = 0
	}

	orderNo := genOrderNo()
	orderId, err := dao.TavernOrder.Ctx(ctx).Data(g.Map{
		orderCols.StoreId:          1,
		orderCols.SessionId:        in.SessionId,
		orderCols.TableId:          tableId,
		orderCols.TableName:        tableName,
		orderCols.OrderNo:          orderNo,
		orderCols.OrderSeq:         orderSeq,
		orderCols.MemberId:         memberId,
		orderCols.PayerMemberId:    memberId,
		orderCols.NicknameSnapshot: profile.Nickname,
		orderCols.TotalAmount:      total,
		orderCols.PayAmount:        payAmount,
		orderCols.DiscountAmount:   discount,
		orderCols.CouponMemberId:   in.CouponId,
		orderCols.PayStatus:        0,
		orderCols.MakingStatus:     0,
		orderCols.RefundStatus:     0,
		orderCols.ExpireAt:         expireAt,
		orderCols.Remark:           in.Remark,
		orderCols.CreatedAt:        gtime.Now(),
		orderCols.UpdatedAt:        gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "创建订单失败")
	}

	itemCols := dao.TavernOrderItem.Columns()
	for _, l := range lines {
		_, err = dao.TavernOrderItem.Ctx(ctx).Data(g.Map{
			itemCols.StoreId:     1,
			itemCols.OrderId:     orderId,
			itemCols.GoodsId:     l.goodsId,
			itemCols.SkuId:       l.skuId,
			itemCols.GoodsName:   l.goodsName,
			itemCols.SkuName:     l.skuName,
			itemCols.UnitPrice:   l.price,
			itemCols.Qty:         l.qty,
			itemCols.Amount:      l.price * float64(l.qty),
			itemCols.IsWine:      boolToInt(l.isWine),
			itemCols.MakerStatus: 0,
			itemCols.Remark:      l.remark,
			itemCols.CreatedAt:   gtime.Now(),
			itemCols.UpdatedAt:   gtime.Now(),
		}).Insert()
		if err != nil {
			return nil, gerror.Wrap(err, "创建订单明细失败")
		}
	}

	// 提交只清掉已下单的行（他人的购物车行不受影响）
	_, _ = dao.TavernCartItem.Ctx(ctx).
		WhereIn(cartCols.Id, in.CartIds).
		Where(cartCols.MemberId, memberId).
		Data(g.Map{
			cartCols.OrderId:   orderId,
			cartCols.DeletedAt: gtime.Now(),
			cartCols.UpdatedAt: gtime.Now(),
		}).Update()

	row, err := dao.TavernOrder.Ctx(ctx).WherePri(orderId).One()
	if err != nil {
		return nil, err
	}
	return toOrderModel(ctx, row)
}

// Pay 支付：先付制下支付成功 = 已接单 = 进入制作
// 说明：当前为模拟支付（渠道 mock）；接入微信支付时，本方法内改为创建预支付单，
//
//	由回调统一调用 markPaid（幂等），并保留原始回调报文用于对账。
func (s *sTavernOrder) Pay(ctx context.Context, in *tavernin.OrderPayInp) (res *tavernin.OrderModel, err error) {
	orderCols := dao.TavernOrder.Columns()
	row, err := dao.TavernOrder.Ctx(ctx).Where(orderCols.OrderNo, in.OrderNo).One()
	if err != nil {
		return nil, err
	}
	if row.IsEmpty() {
		return nil, gerror.New("订单不存在")
	}
	orderId := row[orderCols.Id].Int64()

	// 幂等：已支付直接返回
	if row[orderCols.PayStatus].Int() == 1 {
		return toOrderModel(ctx, row)
	}

	// 超时校验
	if expire := row[orderCols.ExpireAt].GTime(); expire != nil && expire.Before(gtime.Now()) {
		_, _ = dao.TavernOrder.Ctx(ctx).WherePri(orderId).Data(g.Map{
			orderCols.PayStatus: 0,
			orderCols.UpdatedAt: gtime.Now(),
		}).Update()
		return nil, gerror.New("订单已超时，请重新下单")
	}

	channel := in.Channel
	if channel == "" {
		channel = "mock"
	}

	// 支付流水（out_trade_no 唯一索引保证回调幂等）
	payCols := dao.TavernOrderPay.Columns()
	existPay, _ := dao.TavernOrderPay.Ctx(ctx).Where(payCols.OutTradeNo, in.OrderNo).One()
	if existPay.IsEmpty() {
		_, err = dao.TavernOrderPay.Ctx(ctx).Data(g.Map{
			payCols.StoreId:     1,
			payCols.OrderId:     orderId,
			payCols.OutTradeNo:  in.OrderNo,
			payCols.PayChannel:  channel,
			payCols.Amount:      row[orderCols.PayAmount].Float64(),
			payCols.PayStatus:   1,
			payCols.PaidAt:      gtime.Now(),
			payCols.RawCallback: gconv.String(g.Map{"channel": channel, "note": "mock pay"}),
			payCols.CreatedAt:   gtime.Now(),
			payCols.UpdatedAt:   gtime.Now(),
		}).Insert()
		if err != nil {
			return nil, gerror.Wrap(err, "写入支付流水失败")
		}
	}

	// 支付成功 → 立即进入制作（无人工接单环节）
	_, err = dao.TavernOrder.Ctx(ctx).WherePri(orderId).Data(g.Map{
		orderCols.PayStatus:    1,
		orderCols.MakingStatus: 1,
		orderCols.PayChannel:   channel,
		orderCols.PayTime:      gtime.Now(),
		orderCols.UpdatedAt:    gtime.Now(),
	}).Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新订单状态失败")
	}

	itemCols := dao.TavernOrderItem.Columns()
	_, _ = dao.TavernOrderItem.Ctx(ctx).
		Where(itemCols.OrderId, orderId).
		Data(g.Map{
			itemCols.MakerStatus: 1,
			itemCols.UpdatedAt:   gtime.Now(),
		}).Update()

	// 会话累计已付金额
	sessionCols := dao.TavernTableSession.Columns()
	_, _ = dao.TavernTableSession.Ctx(ctx).
		WherePri(row[orderCols.SessionId].Int64()).
		Increment(sessionCols.TotalPaidAmount, row[orderCols.PayAmount].Float64())

	// 核销优惠券（支付成功才真正使用）
	if couponMemberId := row[orderCols.CouponMemberId].Int64(); couponMemberId > 0 {
		if cerr := service.TavernCoupon().UseCoupon(ctx, row[orderCols.MemberId].Int64(), couponMemberId, orderId, row[orderCols.DiscountAmount].Float64()); cerr != nil {
			g.Log().Warningf(ctx, "核销优惠券失败 orderId:%d err:%+v", orderId, cerr)
		}
	}

	// TODO 出票：酒水走吧台、小食走厨房（云打印机 / KDS 推送）
	// TODO 订阅消息：向客人推送「已下单，正在制作」
	g.Log().Infof(ctx, "tavern order paid orderNo:%s amount:%v", in.OrderNo, row[orderCols.PayAmount].Float64())

	updated, err := dao.TavernOrder.Ctx(ctx).WherePri(orderId).One()
	if err != nil {
		return nil, err
	}
	return toOrderModel(ctx, updated)
}

// Detail 订单详情
func (s *sTavernOrder) Detail(ctx context.Context, in *tavernin.OrderDetailInp) (res *tavernin.OrderModel, err error) {
	row, err := dao.TavernOrder.Ctx(ctx).
		Where(dao.TavernOrder.Columns().OrderNo, in.OrderNo).
		One()
	if err != nil {
		return nil, err
	}
	if row.IsEmpty() {
		return nil, gerror.New("订单不存在")
	}
	return toOrderModel(ctx, row)
}

// List 订单列表（本桌 / 我下的单）
func (s *sTavernOrder) List(ctx context.Context, in *tavernin.OrderListInp) (res []*tavernin.OrderModel, err error) {
	memberId := int64(0)
	if in.Mine {
		memberId = service.TavernMember().CurrentId(ctx)
		if memberId == 0 {
			return nil, gerror.New("请先登录")
		}
	}
	return loadOrders(ctx, in.SessionId, memberId, false)
}

func genOrderNo() string {
	return "OD" + gtime.Now().Format("YmdHis") + gconv.String(gtime.TimestampMilli()%1000)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
