// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernOrderPayDao is the data access object for the table hg_tavern_order_pay.
type TavernOrderPayDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  TavernOrderPayColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// TavernOrderPayColumns defines and stores column names for the table hg_tavern_order_pay.
type TavernOrderPayColumns struct {
	Id            string // 主键
	StoreId       string // 门店ID
	OrderId       string // 订单ID
	OutTradeNo    string // 商户订单号（唯一，防重复回调）
	TransactionId string // 第三方交易号
	PayChannel    string // 支付渠道
	Amount        string // 支付金额
	PayStatus     string // 状态：0待支付 1已支付 2失败
	PaidAt        string // 支付时间
	RawCallback   string // 回调原文（对账用）
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// tavernOrderPayColumns holds the columns for the table hg_tavern_order_pay.
var tavernOrderPayColumns = TavernOrderPayColumns{
	Id:            "id",
	StoreId:       "store_id",
	OrderId:       "order_id",
	OutTradeNo:    "out_trade_no",
	TransactionId: "transaction_id",
	PayChannel:    "pay_channel",
	Amount:        "amount",
	PayStatus:     "pay_status",
	PaidAt:        "paid_at",
	RawCallback:   "raw_callback",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewTavernOrderPayDao creates and returns a new DAO object for table data access.
func NewTavernOrderPayDao(handlers ...gdb.ModelHandler) *TavernOrderPayDao {
	return &TavernOrderPayDao{
		group:    "default",
		table:    "hg_tavern_order_pay",
		columns:  tavernOrderPayColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernOrderPayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernOrderPayDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernOrderPayDao) Columns() TavernOrderPayColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernOrderPayDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernOrderPayDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernOrderPayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
