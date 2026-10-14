// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernCategoryDao is the data access object for the table hg_tavern_category.
type TavernCategoryDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  TavernCategoryColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// TavernCategoryColumns defines and stores column names for the table hg_tavern_category.
type TavernCategoryColumns struct {
	Id        string // 主键
	StoreId   string // 门店ID
	Name      string // 分类名称
	Icon      string // 图标
	TimeSlots string // 可见时段，如 18:00-23:59，空为全天
	Status    string // 状态：1正常 2禁用
	Sort      string // 排序
	Remark    string // 备注
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	CreatedBy string // 创建者
	UpdatedBy string // 更新者
}

// tavernCategoryColumns holds the columns for the table hg_tavern_category.
var tavernCategoryColumns = TavernCategoryColumns{
	Id:        "id",
	StoreId:   "store_id",
	Name:      "name",
	Icon:      "icon",
	TimeSlots: "time_slots",
	Status:    "status",
	Sort:      "sort",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
}

// NewTavernCategoryDao creates and returns a new DAO object for table data access.
func NewTavernCategoryDao(handlers ...gdb.ModelHandler) *TavernCategoryDao {
	return &TavernCategoryDao{
		group:    "default",
		table:    "hg_tavern_category",
		columns:  tavernCategoryColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernCategoryDao) Columns() TavernCategoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernCategoryDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
