// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernCouponMember is the golang structure for table tavern_coupon_member.
type TavernCouponMember struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`
	StoreId   int64       `json:"storeId"   orm:"store_id"   description:"门店ID"`
	CouponId  int64       `json:"couponId"  orm:"coupon_id"  description:"券模板ID"`
	MemberId  int64       `json:"memberId"  orm:"member_id"  description:"会员ID"`
	Code      string      `json:"code"      orm:"code"       description:"券码"`
	Source    int         `json:"source"    orm:"source"     description:"来源：1新人礼 2活动发放 3后台定向"`
	Status    int         `json:"status"    orm:"status"     description:"状态：1未使用 2已使用 3已过期"`
	OrderId   int64       `json:"orderId"   orm:"order_id"   description:"使用的订单ID"`
	Amount    float64     `json:"amount"    orm:"amount"     description:"抵扣金额快照"`
	ExpireAt  *gtime.Time `json:"expireAt"  orm:"expire_at"  description:"过期时间"`
	UsedAt    *gtime.Time `json:"usedAt"    orm:"used_at"    description:"使用时间"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
