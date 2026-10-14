package tavern

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"hotgo/internal/dao"
)

// 库存处理约定：
//   - stock < 0 表示不限量，不做扣减；
//   - 提交订单时按「行级原子扣减」预占库存，扣不到即返回该行缺货，避免超卖；
//   - 订单超时取消 / 制作前退款时回滚库存，并通过备注标记保证只回滚一次。

// deductStock 原子扣减库存（仅限限量规格）
func deductStock(ctx context.Context, skuId int64, qty int) (err error) {
	cols := dao.TavernGoodsSku.Columns()
	sku, err := dao.TavernGoodsSku.Ctx(ctx).WherePri(skuId).One()
	if err != nil {
		return err
	}
	if sku.IsEmpty() {
		return gerror.New("规格不存在")
	}
	stock := sku[cols.Stock].Int()
	if stock < 0 {
		return nil // 不限量
	}

	result, err := dao.TavernGoodsSku.Ctx(ctx).
		WherePri(skuId).
		Where("stock >= ?", qty).
		Decrement(cols.Stock, qty)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		goodsName := ""
		if goods, _ := dao.TavernGoods.Ctx(ctx).WherePri(sku[cols.GoodsId].Int64()).One(); !goods.IsEmpty() {
			goodsName = goods[dao.TavernGoods.Columns().Name].String()
		}
		if goodsName == "" {
			goodsName = "所选商品"
		}
		return gerror.New(goodsName + " 库存不足，请调整数量或移除后重试")
	}
	return nil
}

// restoreStockByOrder 回滚订单占用的库存（幂等：以订单备注标记为准）
func restoreStockByOrder(ctx context.Context, orderId int64) (err error) {
	orderCols := dao.TavernOrder.Columns()
	order, err := dao.TavernOrder.Ctx(ctx).WherePri(orderId).One()
	if err != nil {
		return err
	}
	if order.IsEmpty() {
		return nil
	}
	if gstr.Contains(order[orderCols.Remark].String(), stockReturnedMark) {
		return nil
	}

	itemCols := dao.TavernOrderItem.Columns()
	items, err := dao.TavernOrderItem.Ctx(ctx).Where(itemCols.OrderId, orderId).All()
	if err != nil {
		return err
	}

	skuCols := dao.TavernGoodsSku.Columns()
	for _, it := range items {
		skuId := it[itemCols.SkuId].Int64()
		sku, e := dao.TavernGoodsSku.Ctx(ctx).WherePri(skuId).One()
		if e != nil || sku.IsEmpty() {
			continue
		}
		if sku[skuCols.Stock].Int() < 0 {
			continue // 不限量规格无需回滚
		}
		if _, e = dao.TavernGoodsSku.Ctx(ctx).
			WherePri(skuId).
			Increment(skuCols.Stock, it[itemCols.Qty].Int()); e != nil {
			return e
		}
	}

	_, err = dao.TavernOrder.Ctx(ctx).WherePri(orderId).Data(g.Map{
		orderCols.Remark:    gstr.Trim(order[orderCols.Remark].String() + " " + stockReturnedMark),
		orderCols.UpdatedAt: gtime.Now(),
	}).Update()
	return err
}

// stockReturnedMark 库存已回滚标记（用于幂等）
const stockReturnedMark = "【库存已回滚】"

// RestoreStockByOrder 供其它包（如定时任务）调用的库存回滚入口
type stockRestorer struct{}

// RestoreStockByOrder 回滚订单占用的库存（幂等）
func RestoreStockByOrder(ctx context.Context, orderId int64) error {
	return restoreStockByOrder(ctx, orderId)
}

var _ = stockRestorer{}
