// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernBottleTakeLogDao is the data access object for the table hg_tavern_bottle_take_log.
type TavernBottleTakeLogDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  TavernBottleTakeLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// TavernBottleTakeLogColumns defines and stores column names for the table hg_tavern_bottle_take_log.
type TavernBottleTakeLogColumns struct {
	Id          string // 主键
	StoreId     string // 门店ID
	DepositId   string // 存酒ID
	MemberId    string // 会员ID
	TakeType    string // 类型：1店内续饮 2整瓶带走 3报损
	TakeRatio   string // 本次取用比例(%)
	RemainRatio string // 取后剩余比例(%)
	SessionId   string // 关联桌台会话ID
	OperatorId  string // 操作人员工会员ID
	Remark      string // 备注
	TakeAt      string // 取用时间
	CreatedAt   string // 创建时间
}

// tavernBottleTakeLogColumns holds the columns for the table hg_tavern_bottle_take_log.
var tavernBottleTakeLogColumns = TavernBottleTakeLogColumns{
	Id:          "id",
	StoreId:     "store_id",
	DepositId:   "deposit_id",
	MemberId:    "member_id",
	TakeType:    "take_type",
	TakeRatio:   "take_ratio",
	RemainRatio: "remain_ratio",
	SessionId:   "session_id",
	OperatorId:  "operator_id",
	Remark:      "remark",
	TakeAt:      "take_at",
	CreatedAt:   "created_at",
}

// NewTavernBottleTakeLogDao creates and returns a new DAO object for table data access.
func NewTavernBottleTakeLogDao(handlers ...gdb.ModelHandler) *TavernBottleTakeLogDao {
	return &TavernBottleTakeLogDao{
		group:    "default",
		table:    "hg_tavern_bottle_take_log",
		columns:  tavernBottleTakeLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernBottleTakeLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernBottleTakeLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernBottleTakeLogDao) Columns() TavernBottleTakeLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernBottleTakeLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernBottleTakeLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernBottleTakeLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
