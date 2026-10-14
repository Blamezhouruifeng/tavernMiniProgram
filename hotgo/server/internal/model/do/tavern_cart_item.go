// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernCartItem is the golang structure of table hg_tavern_cart_item for DAO operations like Where/Data.
type TavernCartItem struct {
	g.Meta    `orm:"table:hg_tavern_cart_item, do:true"`
	Id        any         // 主键
	StoreId   any         // 门店ID
	SessionId any         // 会话ID
	MemberId  any         // 加购会员ID（归属）
	GoodsId   any         // 商品ID
	SkuId     any         // 规格ID
	Qty       any         // 数量
	Remark    any         // 备注（忌口/做法）
	OrderId   any         // 下单后回填订单ID
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
