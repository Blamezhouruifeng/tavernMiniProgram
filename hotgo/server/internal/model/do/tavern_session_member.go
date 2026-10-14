// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernSessionMember is the golang structure of table hg_tavern_session_member for DAO operations like Where/Data.
type TavernSessionMember struct {
	g.Meta    `orm:"table:hg_tavern_session_member, do:true"`
	Id        any         // 主键
	StoreId   any         // 门店ID
	SessionId any         // 会话ID
	MemberId  any         // 会员ID
	Nickname  any         // 昵称快照
	Avatar    any         // 头像快照
	IsOwner   any         // 是否开台人：1是 0否
	JoinAt    *gtime.Time // 加入时间
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
