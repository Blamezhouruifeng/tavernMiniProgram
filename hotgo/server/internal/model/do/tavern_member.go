// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernMember is the golang structure of table hg_tavern_member for DAO operations like Where/Data.
type TavernMember struct {
	g.Meta      `orm:"table:hg_tavern_member, do:true"`
	Id          any         // 主键
	StoreId     any         // 门店ID（0为全店通用）
	Openid      any         // 微信openid
	Unionid     any         // 微信unionid
	Nickname    any         // 昵称
	Avatar      any         // 头像
	Phone       any         // 手机号（存酒取酒凭证）
	RealName    any         // 真实姓名
	IsStaff     any         // 是否员工：1是 0否
	StaffRole   any         // 员工角色：waiter服务员/bar吧台/kitchen后厨/manager店长
	StaffStatus any         // 员工状态：1在职 2停用
	Points      any         // 积分
	Balance     any         // 储值余额
	TotalAmount any         // 累计消费
	LastLoginAt *gtime.Time // 最后登录时间
	Status      any         // 状态：1正常 2禁用
	Remark      any         // 备注
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
	CreatedBy   any         // 创建者
	UpdatedBy   any         // 更新者
}
