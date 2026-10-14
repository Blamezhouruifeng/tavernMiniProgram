// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernTable is the golang structure of table hg_tavern_table for DAO operations like Where/Data.
type TavernTable struct {
	g.Meta      `orm:"table:hg_tavern_table, do:true"`
	Id          any         // 主键
	StoreId     any         // 门店ID
	AreaId      any         // 区域ID
	Name        any         // 桌号
	Capacity    any         // 可容纳人数
	Qrcode      any         // 桌码文件路径
	QrcodeScene any         // 桌码scene参数
	SeatStatus  any         // 桌位状态：1空闲 2用餐中 3待清台
	Status      any         // 状态：1正常 2禁用
	Sort        any         // 排序
	Remark      any         // 备注
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
	CreatedBy   any         // 创建者
	UpdatedBy   any         // 更新者
}
