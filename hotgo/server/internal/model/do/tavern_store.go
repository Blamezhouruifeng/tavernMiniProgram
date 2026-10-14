// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernStore is the golang structure of table hg_tavern_store for DAO operations like Where/Data.
type TavernStore struct {
	g.Meta           `orm:"table:hg_tavern_store, do:true"`
	Id               any         // 主键
	Name             any         // 门店名称
	Brief            any         // 简介
	Address          any         // 地址
	Phone            any         // 联系电话
	BusinessOpen     any         // 开市时间
	BusinessClose    any         // 打烊时间（可跨天）
	LastOrderGap     any         // 打烊前多少分钟停止下单
	PayFirst         any         // 是否先付制：1是 0否
	PayExpireMinutes any         // 待支付超时分钟数
	MemberQrcode     any         // 员工码（兜底入口）路径
	Status           any         // 状态：1正常 2停用
	Sort             any         // 排序
	Remark           any         // 备注
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	DeletedAt        *gtime.Time // 删除时间
	CreatedBy        any         // 创建者
	UpdatedBy        any         // 更新者
}
