// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernOrder is the golang structure of table hg_tavern_order for DAO operations like Where/Data.
type TavernOrder struct {
	g.Meta           `orm:"table:hg_tavern_order, do:true"`
	Id               any         // 主键
	StoreId          any         // 门店ID
	SessionId        any         // 会话ID
	TableId          any         // 桌位ID
	TableName        any         // 桌号快照
	OrderNo          any         // 订单号
	OrderSeq         any         // 本桌第几单
	MemberId         any         // 下单会员ID
	PayerMemberId    any         // 付款会员ID（可代付）
	NicknameSnapshot any         // 下单人昵称快照
	TotalAmount      any         // 订单总额
	PayAmount        any         // 实付金额
	DiscountAmount   any         // 优惠金额
	PayStatus        any         // 支付状态：0未支付 1已支付 2部分退款 3已退款
	MakingStatus     any         // 制作状态：0未开始 1制作中 2已上齐 3已完成
	RefundStatus     any         // 退款状态：0无 1申请中 2已通过 3已拒绝 4已退款
	RefundAmount     any         // 退款金额
	RefundReason     any         // 退款原因
	RefundAt         *gtime.Time // 退款时间
	RefundAuditBy    any         // 退款审核人
	PayChannel       any         // 支付渠道：wxpay/mock/balance
	OutTradeNo       any         // 商户订单号
	ExpireAt         *gtime.Time // 待支付超时时间
	PayTime          *gtime.Time // 支付时间
	Remark           any         // 备注
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	DeletedAt        *gtime.Time // 删除时间
}
