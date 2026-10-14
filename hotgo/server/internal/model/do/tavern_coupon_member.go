// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernCouponMember is the golang structure of table hg_tavern_coupon_member for DAO operations like Where/Data.
type TavernCouponMember struct {
	g.Meta    `orm:"table:hg_tavern_coupon_member, do:true"`
	Id        any         // 主键
	StoreId   any         // 门店ID
	CouponId  any         // 券模板ID
	MemberId  any         // 会员ID
	Code      any         // 券码
	Source    any         // 来源：1新人礼 2活动发放 3后台定向
	Status    any         // 状态：1未使用 2已使用 3已过期
	OrderId   any         // 使用的订单ID
	Amount    any         // 抵扣金额快照
	ExpireAt  *gtime.Time // 过期时间
	UsedAt    *gtime.Time // 使用时间
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
