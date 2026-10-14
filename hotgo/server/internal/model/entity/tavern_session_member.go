// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernSessionMember is the golang structure for table tavern_session_member.
type TavernSessionMember struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`
	StoreId   int64       `json:"storeId"   orm:"store_id"   description:"门店ID"`
	SessionId int64       `json:"sessionId" orm:"session_id" description:"会话ID"`
	MemberId  int64       `json:"memberId"  orm:"member_id"  description:"会员ID"`
	Nickname  string      `json:"nickname"  orm:"nickname"   description:"昵称快照"`
	Avatar    string      `json:"avatar"    orm:"avatar"     description:"头像快照"`
	IsOwner   int         `json:"isOwner"   orm:"is_owner"   description:"是否开台人：1是 0否"`
	JoinAt    *gtime.Time `json:"joinAt"    orm:"join_at"    description:"加入时间"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
