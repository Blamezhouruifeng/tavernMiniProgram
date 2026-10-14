// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernOrderDao is the data access object for the table hg_tavern_order.
type TavernOrderDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TavernOrderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TavernOrderColumns defines and stores column names for the table hg_tavern_order.
type TavernOrderColumns struct {
	Id               string // 主键
	StoreId          string // 门店ID
	SessionId        string // 会话ID
	TableId          string // 桌位ID
	TableName        string // 桌号快照
	OrderNo          string // 订单号
	OrderSeq         string // 本桌第几单
	MemberId         string // 下单会员ID
	PayerMemberId    string // 付款会员ID（可代付）
	NicknameSnapshot string // 下单人昵称快照
	TotalAmount      string // 订单总额
	PayAmount        string // 实付金额
	DiscountAmount   string // 优惠金额
	CouponMemberId   string // 使用的会员券ID
	PayStatus        string // 支付状态：0未支付 1已支付 2部分退款 3已退款
	MakingStatus     string // 制作状态：0未开始 1制作中 2已上齐 3已完成
	RefundStatus     string // 退款状态：0无 1申请中 2已通过 3已拒绝 4已退款
	RefundAmount     string // 退款金额
	RefundReason     string // 退款原因
	RefundAt         string // 退款时间
	RefundAuditBy    string // 退款审核人
	PayChannel       string // 支付渠道：wxpay/mock/balance
	OutTradeNo       string // 商户订单号
	ExpireAt         string // 待支付超时时间
	PayTime          string // 支付时间
	Remark           string // 备注
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
	DeletedAt        string // 删除时间
}

// tavernOrderColumns holds the columns for the table hg_tavern_order.
var tavernOrderColumns = TavernOrderColumns{
	Id:               "id",
	StoreId:          "store_id",
	SessionId:        "session_id",
	TableId:          "table_id",
	TableName:        "table_name",
	OrderNo:          "order_no",
	OrderSeq:         "order_seq",
	MemberId:         "member_id",
	PayerMemberId:    "payer_member_id",
	NicknameSnapshot: "nickname_snapshot",
	TotalAmount:      "total_amount",
	PayAmount:        "pay_amount",
	DiscountAmount:   "discount_amount",
	CouponMemberId:   "coupon_member_id",
	PayStatus:        "pay_status",
	MakingStatus:     "making_status",
	RefundStatus:     "refund_status",
	RefundAmount:     "refund_amount",
	RefundReason:     "refund_reason",
	RefundAt:         "refund_at",
	RefundAuditBy:    "refund_audit_by",
	PayChannel:       "pay_channel",
	OutTradeNo:       "out_trade_no",
	ExpireAt:         "expire_at",
	PayTime:          "pay_time",
	Remark:           "remark",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// NewTavernOrderDao creates and returns a new DAO object for table data access.
func NewTavernOrderDao(handlers ...gdb.ModelHandler) *TavernOrderDao {
	return &TavernOrderDao{
		group:    "default",
		table:    "hg_tavern_order",
		columns:  tavernOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernOrderDao) Columns() TavernOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernOrderDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TavernOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
