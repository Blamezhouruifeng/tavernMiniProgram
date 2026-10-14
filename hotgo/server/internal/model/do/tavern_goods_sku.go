// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernGoodsSku is the golang structure of table hg_tavern_goods_sku for DAO operations like Where/Data.
type TavernGoodsSku struct {
	g.Meta        `orm:"table:hg_tavern_goods_sku, do:true"`
	Id            any         // 主键
	StoreId       any         // 门店ID
	GoodsId       any         // 商品ID
	Name          any         // 规格名称：杯/瓶/扎
	Price         any         // 售价
	Stock         any         // 库存，-1为不限量
	UnitCups      any         // 换算杯数（1瓶=N杯）
	IsDepositUnit any         // 是否为整瓶可寄存单位：1是 0否
	Status        any         // 状态：1正常 2禁用
	Sort          any         // 排序
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
}
