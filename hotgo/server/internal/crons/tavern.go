package crons

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"hotgo/internal/dao"
	"hotgo/internal/library/cron"
)

func init() {
	cron.Register(TavernOrderExpire)
	cron.Register(TavernBottleExpire)
}

// TavernOrderExpire 待支付超时自动取消（先付制：未支付订单不占库存、不出票）
// 注意：需要在后台「系统设置 → 定时任务」中按 name = tavern_order_expire 添加任务策略
var TavernOrderExpire = &cTavernOrderExpire{name: "tavern_order_expire"}

type cTavernOrderExpire struct{ name string }

func (c *cTavernOrderExpire) GetName() string { return c.name }

func (c *cTavernOrderExpire) Execute(ctx context.Context, parser *cron.Parser) (err error) {
	orderCols := dao.TavernOrder.Columns()
	rows, err := dao.TavernOrder.Ctx(ctx).
		Fields(orderCols.Id, orderCols.OrderNo).
		Where(orderCols.PayStatus, 0).
		WhereLT(orderCols.ExpireAt, gtime.Now()).
		Limit(200).
		All()
	if err != nil {
		parser.Logger.Warning(ctx, "TavernOrderExpire query err:%+v", err)
		return
	}
	if len(rows) == 0 {
		return
	}

	for _, row := range rows {
		orderId := row[orderCols.Id].Int64()
		if _, err = dao.TavernOrder.Ctx(ctx).WherePri(orderId).Data(g.Map{
			orderCols.MakingStatus: 0,
			orderCols.Remark:       "【已超时取消】",
			orderCols.UpdatedAt:    gtime.Now(),
		}).Update(); err != nil {
			parser.Logger.Warning(ctx, "TavernOrderExpire update err:%+v", err)
			return
		}
		parser.Logger.Infof(ctx, "待支付订单已超时取消 orderNo:%s", row[orderCols.OrderNo].String())
	}
	return
}

// TavernBottleExpire 存酒到期处理：标记过期并记录（到期前提醒在业务侧按 expire_at 计算展示）
var TavernBottleExpire = &cTavernBottleExpire{name: "tavern_bottle_expire"}

type cTavernBottleExpire struct{ name string }

func (c *cTavernBottleExpire) GetName() string { return c.name }

func (c *cTavernBottleExpire) Execute(ctx context.Context, parser *cron.Parser) (err error) {
	depCols := dao.TavernBottleDeposit.Columns()
	affected, err := dao.TavernBottleDeposit.Ctx(ctx).
		Where(depCols.Status, 1).
		WhereLT(depCols.ExpireAt, gtime.Now()).
		Data(g.Map{
			depCols.Status:    3,
			depCols.UpdatedAt: gtime.Now(),
		}).UpdateAndGetAffected()
	if err != nil {
		parser.Logger.Warning(ctx, "TavernBottleExpire err:%+v", err)
		return
	}
	if affected > 0 {
		parser.Logger.Infof(ctx, "存酒到期处理完成，共 %d 条转为已过期", affected)
	}
	return
}
