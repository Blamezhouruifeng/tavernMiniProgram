// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernMember is the golang structure for table tavern_member.
type TavernMember struct {
	Id          int64       `json:"id"          orm:"id"            description:"主键"`
	StoreId     int64       `json:"storeId"     orm:"store_id"      description:"门店ID（0为全店通用）"`
	Openid      string      `json:"openid"      orm:"openid"        description:"微信openid"`
	Unionid     string      `json:"unionid"     orm:"unionid"       description:"微信unionid"`
	Nickname    string      `json:"nickname"    orm:"nickname"      description:"昵称"`
	Avatar      string      `json:"avatar"      orm:"avatar"        description:"头像"`
	Phone       string      `json:"phone"       orm:"phone"         description:"手机号（存酒取酒凭证）"`
	RealName    string      `json:"realName"    orm:"real_name"     description:"真实姓名"`
	IsStaff     int         `json:"isStaff"     orm:"is_staff"      description:"是否员工：1是 0否"`
	StaffRole   string      `json:"staffRole"   orm:"staff_role"    description:"员工角色：waiter服务员/bar吧台/kitchen后厨/manager店长"`
	StaffStatus int         `json:"staffStatus" orm:"staff_status"  description:"员工状态：1在职 2停用"`
	Points      int         `json:"points"      orm:"points"        description:"积分"`
	Balance     float64     `json:"balance"     orm:"balance"       description:"储值余额"`
	TotalAmount float64     `json:"totalAmount" orm:"total_amount"  description:"累计消费"`
	LastLoginAt *gtime.Time `json:"lastLoginAt" orm:"last_login_at" description:"最后登录时间"`
	Status      int         `json:"status"      orm:"status"        description:"状态：1正常 2禁用"`
	Remark      string      `json:"remark"      orm:"remark"        description:"备注"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"    description:"删除时间"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"    description:"创建者"`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"    description:"更新者"`
}
