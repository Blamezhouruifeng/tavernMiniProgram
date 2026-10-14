// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernGoodsSku is the golang structure for table tavern_goods_sku.
type TavernGoodsSku struct {
	Id            int64       `json:"id"            orm:"id"              description:"主键"`
	StoreId       int64       `json:"storeId"       orm:"store_id"        description:"门店ID"`
	GoodsId       int64       `json:"goodsId"       orm:"goods_id"        description:"商品ID"`
	Name          string      `json:"name"          orm:"name"            description:"规格名称：杯/瓶/扎"`
	Price         float64     `json:"price"         orm:"price"           description:"售价"`
	Stock         int         `json:"stock"         orm:"stock"           description:"库存，-1为不限量"`
	UnitCups      float64     `json:"unitCups"      orm:"unit_cups"       description:"换算杯数（1瓶=N杯）"`
	IsDepositUnit int         `json:"isDepositUnit" orm:"is_deposit_unit" description:"是否为整瓶可寄存单位：1是 0否"`
	Status        int         `json:"status"        orm:"status"          description:"状态：1正常 2禁用"`
	Sort          int         `json:"sort"          orm:"sort"            description:"排序"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"      description:"删除时间"`
}
