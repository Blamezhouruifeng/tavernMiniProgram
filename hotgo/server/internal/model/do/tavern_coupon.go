// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernCoupon is the golang structure of table hg_tavern_coupon for DAO operations like Where/Data.
type TavernCoupon struct {
	g.Meta         `orm:"table:hg_tavern_coupon, do:true"`
	Id             any         // 主键
	StoreId        any         // 门店ID
	Name           any         // 券名称
	Type           any         // 类型：1满减 2折扣
	Amount         any         // 满减金额
	MinAmount      any         // 使用门槛（订单满多少可用）
	Discount       any         // 折扣（如 0.88 表示八八折）
	ValidDays      any         // 领取后有效天数
	TotalLimit     any         // 发放总量，-1不限
	PerMemberLimit any         // 每人限领
	IsNewcomer     any         // 是否新人礼（注册自动发放）
	Status         any         // 状态：1启用 2停用
	Sort           any         // 排序
	Remark         any         // 备注
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
	CreatedBy      any         // 创建者
	UpdatedBy      any         // 更新者
}
