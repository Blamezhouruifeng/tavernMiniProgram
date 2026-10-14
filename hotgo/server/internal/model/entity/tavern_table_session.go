// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernTableSession is the golang structure for table tavern_table_session.
type TavernTableSession struct {
	Id               int64       `json:"id"               orm:"id"                 description:"主键"`
	StoreId          int64       `json:"storeId"          orm:"store_id"           description:"门店ID"`
	TableId          int64       `json:"tableId"          orm:"table_id"           description:"桌位ID"`
	SessionNo        string      `json:"sessionNo"        orm:"session_no"         description:"会话号"`
	OwnerMemberId    int64       `json:"ownerMemberId"    orm:"owner_member_id"    description:"开台会员ID"`
	GuestCount       int         `json:"guestCount"       orm:"guest_count"        description:"就餐人数"`
	MemberCount      int         `json:"memberCount"      orm:"member_count"       description:"参与会员数"`
	OrderCount       int         `json:"orderCount"       orm:"order_count"        description:"已支付订单数"`
	TotalPaidAmount  float64     `json:"totalPaidAmount"  orm:"total_paid_amount"  description:"累计已付金额"`
	GuestTotalAmount float64     `json:"guestTotalAmount" orm:"guest_total_amount" description:"全桌合计金额"`
	Status           int         `json:"status"           orm:"status"             description:"会话状态：1进行中 2已归档"`
	OpenAt           *gtime.Time `json:"openAt"           orm:"open_at"            description:"开台时间"`
	CloseAt          *gtime.Time `json:"closeAt"          orm:"close_at"           description:"关台时间"`
	Remark           string      `json:"remark"           orm:"remark"             description:"备注"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:"更新时间"`
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:"删除时间"`
	CreatedBy        int64       `json:"createdBy"        orm:"created_by"         description:"创建者"`
	UpdatedBy        int64       `json:"updatedBy"        orm:"updated_by"         description:"更新者"`
}
