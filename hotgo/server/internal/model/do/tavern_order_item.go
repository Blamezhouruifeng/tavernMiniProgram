// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernOrderItem is the golang structure of table hg_tavern_order_item for DAO operations like Where/Data.
type TavernOrderItem struct {
	g.Meta      `orm:"table:hg_tavern_order_item, do:true"`
	Id          any         // 主键
	StoreId     any         // 门店ID
	OrderId     any         // 订单ID
	GoodsId     any         // 商品ID
	SkuId       any         // 规格ID
	GoodsName   any         // 商品名称快照
	SkuName     any         // 规格名称快照
	UnitPrice   any         // 单价
	Qty         any         // 数量
	Amount      any         // 小计
	IsWine      any         // 是否酒水（出票分单用）
	MakerStatus any         // 行状态：0待制作 1制作中 2已上齐
	ServedAt    *gtime.Time // 上齐时间
	Remark      any         // 备注
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
