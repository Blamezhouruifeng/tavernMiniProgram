// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernArea is the golang structure of table hg_tavern_area for DAO operations like Where/Data.
type TavernArea struct {
	g.Meta    `orm:"table:hg_tavern_area, do:true"`
	Id        any         // 主键
	StoreId   any         // 门店ID
	Name      any         // 区域名称
	Status    any         // 状态：1正常 2禁用
	Sort      any         // 排序
	Remark    any         // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	CreatedBy any         // 创建者
	UpdatedBy any         // 更新者
}
