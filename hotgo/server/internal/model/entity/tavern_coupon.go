// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernCoupon is the golang structure for table tavern_coupon.
type TavernCoupon struct {
	Id             int64       `json:"id"             orm:"id"               description:"主键"`
	StoreId        int64       `json:"storeId"        orm:"store_id"         description:"门店ID"`
	Name           string      `json:"name"           orm:"name"             description:"券名称"`
	Type           int         `json:"type"           orm:"type"             description:"类型：1满减 2折扣"`
	Amount         float64     `json:"amount"         orm:"amount"           description:"满减金额"`
	MinAmount      float64     `json:"minAmount"      orm:"min_amount"       description:"使用门槛（订单满多少可用）"`
	Discount       float64     `json:"discount"       orm:"discount"         description:"折扣（如 0.88 表示八八折）"`
	ValidDays      int         `json:"validDays"      orm:"valid_days"       description:"领取后有效天数"`
	TotalLimit     int         `json:"totalLimit"     orm:"total_limit"      description:"发放总量，-1不限"`
	PerMemberLimit int         `json:"perMemberLimit" orm:"per_member_limit" description:"每人限领"`
	IsNewcomer     int         `json:"isNewcomer"     orm:"is_newcomer"      description:"是否新人礼（注册自动发放）"`
	Status         int         `json:"status"         orm:"status"           description:"状态：1启用 2停用"`
	Sort           int         `json:"sort"           orm:"sort"             description:"排序"`
	Remark         string      `json:"remark"         orm:"remark"           description:"备注"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"       description:"删除时间"`
	CreatedBy      int64       `json:"createdBy"      orm:"created_by"       description:"创建者"`
	UpdatedBy      int64       `json:"updatedBy"      orm:"updated_by"       description:"更新者"`
}
