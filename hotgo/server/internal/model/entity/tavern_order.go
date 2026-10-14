// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernOrder is the golang structure for table tavern_order.
type TavernOrder struct {
	Id               int64       `json:"id"               orm:"id"                description:"主键"`
	StoreId          int64       `json:"storeId"          orm:"store_id"          description:"门店ID"`
	SessionId        int64       `json:"sessionId"        orm:"session_id"        description:"会话ID"`
	TableId          int64       `json:"tableId"          orm:"table_id"          description:"桌位ID"`
	TableName        string      `json:"tableName"        orm:"table_name"        description:"桌号快照"`
	OrderNo          string      `json:"orderNo"          orm:"order_no"          description:"订单号"`
	OrderSeq         int         `json:"orderSeq"         orm:"order_seq"         description:"本桌第几单"`
	MemberId         int64       `json:"memberId"         orm:"member_id"         description:"下单会员ID"`
	PayerMemberId    int64       `json:"payerMemberId"    orm:"payer_member_id"   description:"付款会员ID（可代付）"`
	NicknameSnapshot string      `json:"nicknameSnapshot" orm:"nickname_snapshot" description:"下单人昵称快照"`
	TotalAmount      float64     `json:"totalAmount"      orm:"total_amount"      description:"订单总额"`
	PayAmount        float64     `json:"payAmount"        orm:"pay_amount"        description:"实付金额"`
	DiscountAmount   float64     `json:"discountAmount"   orm:"discount_amount"   description:"优惠金额"`
	CouponMemberId   int64       `json:"couponMemberId"   orm:"coupon_member_id"  description:"使用的会员券ID"`
	PayStatus        int         `json:"payStatus"        orm:"pay_status"        description:"支付状态：0未支付 1已支付 2部分退款 3已退款"`
	MakingStatus     int         `json:"makingStatus"     orm:"making_status"     description:"制作状态：0未开始 1制作中 2已上齐 3已完成"`
	RefundStatus     int         `json:"refundStatus"     orm:"refund_status"     description:"退款状态：0无 1申请中 2已通过 3已拒绝 4已退款"`
	RefundAmount     float64     `json:"refundAmount"     orm:"refund_amount"     description:"退款金额"`
	RefundReason     string      `json:"refundReason"     orm:"refund_reason"     description:"退款原因"`
	RefundAt         *gtime.Time `json:"refundAt"         orm:"refund_at"         description:"退款时间"`
	RefundAuditBy    int64       `json:"refundAuditBy"    orm:"refund_audit_by"   description:"退款审核人"`
	PayChannel       string      `json:"payChannel"       orm:"pay_channel"       description:"支付渠道：wxpay/mock/balance"`
	OutTradeNo       string      `json:"outTradeNo"       orm:"out_trade_no"      description:"商户订单号"`
	ExpireAt         *gtime.Time `json:"expireAt"         orm:"expire_at"         description:"待支付超时时间"`
	PayTime          *gtime.Time `json:"payTime"          orm:"pay_time"          description:"支付时间"`
	Remark           string      `json:"remark"           orm:"remark"            description:"备注"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"        description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"        description:"更新时间"`
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"        description:"删除时间"`
}
