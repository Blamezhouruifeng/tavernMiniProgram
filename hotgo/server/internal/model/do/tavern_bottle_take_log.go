// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernBottleTakeLog is the golang structure of table hg_tavern_bottle_take_log for DAO operations like Where/Data.
type TavernBottleTakeLog struct {
	g.Meta      `orm:"table:hg_tavern_bottle_take_log, do:true"`
	Id          any         // 主键
	StoreId     any         // 门店ID
	DepositId   any         // 存酒ID
	MemberId    any         // 会员ID
	TakeType    any         // 类型：1店内续饮 2整瓶带走 3报损
	TakeRatio   any         // 本次取用比例(%)
	RemainRatio any         // 取后剩余比例(%)
	SessionId   any         // 关联桌台会话ID
	OperatorId  any         // 操作人员工会员ID
	Remark      any         // 备注
	TakeAt      *gtime.Time // 取用时间
	CreatedAt   *gtime.Time // 创建时间
}
