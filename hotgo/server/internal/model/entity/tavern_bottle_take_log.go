// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernBottleTakeLog is the golang structure for table tavern_bottle_take_log.
type TavernBottleTakeLog struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键"`
	StoreId     int64       `json:"storeId"     orm:"store_id"     description:"门店ID"`
	DepositId   int64       `json:"depositId"   orm:"deposit_id"   description:"存酒ID"`
	MemberId    int64       `json:"memberId"    orm:"member_id"    description:"会员ID"`
	TakeType    int         `json:"takeType"    orm:"take_type"    description:"类型：1店内续饮 2整瓶带走 3报损"`
	TakeRatio   float64     `json:"takeRatio"   orm:"take_ratio"   description:"本次取用比例(%)"`
	RemainRatio float64     `json:"remainRatio" orm:"remain_ratio" description:"取后剩余比例(%)"`
	SessionId   int64       `json:"sessionId"   orm:"session_id"   description:"关联桌台会话ID"`
	OperatorId  int64       `json:"operatorId"  orm:"operator_id"  description:"操作人员工会员ID"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
	TakeAt      *gtime.Time `json:"takeAt"      orm:"take_at"      description:"取用时间"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
}
