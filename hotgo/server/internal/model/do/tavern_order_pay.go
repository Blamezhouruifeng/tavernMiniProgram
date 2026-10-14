// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernOrderPay is the golang structure of table hg_tavern_order_pay for DAO operations like Where/Data.
type TavernOrderPay struct {
	g.Meta        `orm:"table:hg_tavern_order_pay, do:true"`
	Id            any         // 主键
	StoreId       any         // 门店ID
	OrderId       any         // 订单ID
	OutTradeNo    any         // 商户订单号（唯一，防重复回调）
	TransactionId any         // 第三方交易号
	PayChannel    any         // 支付渠道
	Amount        any         // 支付金额
	PayStatus     any         // 状态：0待支付 1已支付 2失败
	PaidAt        *gtime.Time // 支付时间
	RawCallback   any         // 回调原文（对账用）
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
