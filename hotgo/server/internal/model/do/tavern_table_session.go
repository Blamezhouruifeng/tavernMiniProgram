// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernTableSession is the golang structure of table hg_tavern_table_session for DAO operations like Where/Data.
type TavernTableSession struct {
	g.Meta           `orm:"table:hg_tavern_table_session, do:true"`
	Id               any         // 主键
	StoreId          any         // 门店ID
	TableId          any         // 桌位ID
	SessionNo        any         // 会话号
	OwnerMemberId    any         // 开台会员ID
	GuestCount       any         // 就餐人数
	MemberCount      any         // 参与会员数
	OrderCount       any         // 已支付订单数
	TotalPaidAmount  any         // 累计已付金额
	GuestTotalAmount any         // 全桌合计金额
	Status           any         // 会话状态：1进行中 2已归档
	OpenAt           *gtime.Time // 开台时间
	CloseAt          *gtime.Time // 关台时间
	Remark           any         // 备注
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	DeletedAt        *gtime.Time // 删除时间
	CreatedBy        any         // 创建者
	UpdatedBy        any         // 更新者
}
