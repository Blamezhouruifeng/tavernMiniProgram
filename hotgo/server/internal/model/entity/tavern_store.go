// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernStore is the golang structure for table tavern_store.
type TavernStore struct {
	Id               int64       `json:"id"               orm:"id"                 description:"主键"`
	Name             string      `json:"name"             orm:"name"               description:"门店名称"`
	Brief            string      `json:"brief"            orm:"brief"              description:"简介"`
	Address          string      `json:"address"          orm:"address"            description:"地址"`
	Phone            string      `json:"phone"            orm:"phone"              description:"联系电话"`
	BusinessOpen     string      `json:"businessOpen"     orm:"business_open"      description:"开市时间"`
	BusinessClose    string      `json:"businessClose"    orm:"business_close"     description:"打烊时间（可跨天）"`
	LastOrderGap     int         `json:"lastOrderGap"     orm:"last_order_gap"     description:"打烊前多少分钟停止下单"`
	PayFirst         int         `json:"payFirst"         orm:"pay_first"          description:"是否先付制：1是 0否"`
	PayExpireMinutes int         `json:"payExpireMinutes" orm:"pay_expire_minutes" description:"待支付超时分钟数"`
	MemberQrcode     string      `json:"memberQrcode"     orm:"member_qrcode"      description:"员工码（兜底入口）路径"`
	Status           int         `json:"status"           orm:"status"             description:"状态：1正常 2停用"`
	Sort             int         `json:"sort"             orm:"sort"               description:"排序"`
	Remark           string      `json:"remark"           orm:"remark"             description:"备注"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:"更新时间"`
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:"删除时间"`
	CreatedBy        int64       `json:"createdBy"        orm:"created_by"         description:"创建者"`
	UpdatedBy        int64       `json:"updatedBy"        orm:"updated_by"         description:"更新者"`
}
