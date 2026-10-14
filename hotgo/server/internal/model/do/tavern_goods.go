// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernGoods is the golang structure of table hg_tavern_goods for DAO operations like Where/Data.
type TavernGoods struct {
	g.Meta      `orm:"table:hg_tavern_goods, do:true"`
	Id          any         // 主键
	StoreId     any         // 门店ID
	CategoryId  any         // 分类ID
	Name        any         // 商品名称
	Cover       any         // 封面图
	Brief       any         // 简介
	Price       any         // 最低售价（展示用）
	Tags        any         // 标签，逗号分隔：招牌,热销,新品
	IsWine      any         // 是否酒水：1是 0否
	OnlyDineIn  any         // 仅堂食：1是 0否
	IsCleared   any         // 是否沽清：1是 0否
	MakeMinutes any         // 预计制作分钟数
	Status      any         // 状态：1上架 2下架
	Sort        any         // 排序
	Remark      any         // 备注
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
	CreatedBy   any         // 创建者
	UpdatedBy   any         // 更新者
}
