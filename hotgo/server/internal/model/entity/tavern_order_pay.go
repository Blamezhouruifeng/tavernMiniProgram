// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernOrderPay is the golang structure for table tavern_order_pay.
type TavernOrderPay struct {
	Id            int64       `json:"id"            orm:"id"             description:"主键"`
	StoreId       int64       `json:"storeId"       orm:"store_id"       description:"门店ID"`
	OrderId       int64       `json:"orderId"       orm:"order_id"       description:"订单ID"`
	OutTradeNo    string      `json:"outTradeNo"    orm:"out_trade_no"   description:"商户订单号（唯一，防重复回调）"`
	TransactionId string      `json:"transactionId" orm:"transaction_id" description:"第三方交易号"`
	PayChannel    string      `json:"payChannel"    orm:"pay_channel"    description:"支付渠道"`
	Amount        float64     `json:"amount"        orm:"amount"         description:"支付金额"`
	PayStatus     int         `json:"payStatus"     orm:"pay_status"     description:"状态：0待支付 1已支付 2失败"`
	PaidAt        *gtime.Time `json:"paidAt"        orm:"paid_at"        description:"支付时间"`
	RawCallback   string      `json:"rawCallback"   orm:"raw_callback"   description:"回调原文（对账用）"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
}
