// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernStorageSlot is the golang structure of table hg_tavern_storage_slot for DAO operations like Where/Data.
type TavernStorageSlot struct {
	g.Meta    `orm:"table:hg_tavern_storage_slot, do:true"`
	Id        any         // 主键
	StoreId   any         // 门店ID
	Area      any         // 储区，如 A
	Layer     any         // 层，如 2
	Cell      any         // 格位，如 3
	Code      any         // 储位号，如 A-2-3
	IsUsed    any         // 是否占用：1是 0否
	Status    any         // 状态：1正常 2禁用
	Remark    any         // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
