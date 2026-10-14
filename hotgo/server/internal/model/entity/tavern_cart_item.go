// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernCartItem is the golang structure for table tavern_cart_item.
type TavernCartItem struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`
	StoreId   int64       `json:"storeId"   orm:"store_id"   description:"门店ID"`
	SessionId int64       `json:"sessionId" orm:"session_id" description:"会话ID"`
	MemberId  int64       `json:"memberId"  orm:"member_id"  description:"加购会员ID（归属）"`
	GoodsId   int64       `json:"goodsId"   orm:"goods_id"   description:"商品ID"`
	SkuId     int64       `json:"skuId"     orm:"sku_id"     description:"规格ID"`
	Qty       int         `json:"qty"       orm:"qty"        description:"数量"`
	Remark    string      `json:"remark"    orm:"remark"     description:"备注（忌口/做法）"`
	OrderId   int64       `json:"orderId"   orm:"order_id"   description:"下单后回填订单ID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
