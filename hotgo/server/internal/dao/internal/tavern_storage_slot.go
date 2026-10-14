// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernStorageSlotDao is the data access object for the table hg_tavern_storage_slot.
type TavernStorageSlotDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  TavernStorageSlotColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// TavernStorageSlotColumns defines and stores column names for the table hg_tavern_storage_slot.
type TavernStorageSlotColumns struct {
	Id        string // 主键
	StoreId   string // 门店ID
	Area      string // 储区，如 A
	Layer     string // 层，如 2
	Cell      string // 格位，如 3
	Code      string // 储位号，如 A-2-3
	IsUsed    string // 是否占用：1是 0否
	Status    string // 状态：1正常 2禁用
	Remark    string // 备注
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// tavernStorageSlotColumns holds the columns for the table hg_tavern_storage_slot.
var tavernStorageSlotColumns = TavernStorageSlotColumns{
	Id:        "id",
	StoreId:   "store_id",
	Area:      "area",
	Layer:     "layer",
	Cell:      "cell",
	Code:      "code",
	IsUsed:    "is_used",
	Status:    "status",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewTavernStorageSlotDao creates and returns a new DAO object for table data access.
func NewTavernStorageSlotDao(handlers ...gdb.ModelHandler) *TavernStorageSlotDao {
	return &TavernStorageSlotDao{
		group:    "default",
		table:    "hg_tavern_storage_slot",
		columns:  tavernStorageSlotColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernStorageSlotDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernStorageSlotDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernStorageSlotDao) Columns() TavernStorageSlotColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernStorageSlotDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernStorageSlotDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernStorageSlotDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
