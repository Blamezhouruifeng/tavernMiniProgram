// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernOrderItem is the golang structure for table tavern_order_item.
type TavernOrderItem struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键"`
	StoreId     int64       `json:"storeId"     orm:"store_id"     description:"门店ID"`
	OrderId     int64       `json:"orderId"     orm:"order_id"     description:"订单ID"`
	GoodsId     int64       `json:"goodsId"     orm:"goods_id"     description:"商品ID"`
	SkuId       int64       `json:"skuId"       orm:"sku_id"       description:"规格ID"`
	GoodsName   string      `json:"goodsName"   orm:"goods_name"   description:"商品名称快照"`
	SkuName     string      `json:"skuName"     orm:"sku_name"     description:"规格名称快照"`
	UnitPrice   float64     `json:"unitPrice"   orm:"unit_price"   description:"单价"`
	Qty         int         `json:"qty"         orm:"qty"          description:"数量"`
	Amount      float64     `json:"amount"      orm:"amount"       description:"小计"`
	IsWine      int         `json:"isWine"      orm:"is_wine"      description:"是否酒水（出票分单用）"`
	MakerStatus int         `json:"makerStatus" orm:"maker_status" description:"行状态：0待制作 1制作中 2已上齐"`
	ServedAt    *gtime.Time `json:"servedAt"    orm:"served_at"    description:"上齐时间"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
}
