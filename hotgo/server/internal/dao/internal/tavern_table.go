// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernTableDao is the data access object for the table hg_tavern_table.
type TavernTableDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TavernTableColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TavernTableColumns defines and stores column names for the table hg_tavern_table.
type TavernTableColumns struct {
	Id          string // 主键
	StoreId     string // 门店ID
	AreaId      string // 区域ID
	Name        string // 桌号
	Capacity    string // 可容纳人数
	Qrcode      string // 桌码文件路径
	QrcodeScene string // 桌码scene参数
	SeatStatus  string // 桌位状态：1空闲 2用餐中 3待清台
	Status      string // 状态：1正常 2禁用
	Sort        string // 排序
	Remark      string // 备注
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
	CreatedBy   string // 创建者
	UpdatedBy   string // 更新者
}

// tavernTableColumns holds the columns for the table hg_tavern_table.
var tavernTableColumns = TavernTableColumns{
	Id:          "id",
	StoreId:     "store_id",
	AreaId:      "area_id",
	Name:        "name",
	Capacity:    "capacity",
	Qrcode:      "qrcode",
	QrcodeScene: "qrcode_scene",
	SeatStatus:  "seat_status",
	Status:      "status",
	Sort:        "sort",
	Remark:      "remark",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
}

// NewTavernTableDao creates and returns a new DAO object for table data access.
func NewTavernTableDao(handlers ...gdb.ModelHandler) *TavernTableDao {
	return &TavernTableDao{
		group:    "default",
		table:    "hg_tavern_table",
		columns:  tavernTableColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernTableDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernTableDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernTableDao) Columns() TavernTableColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernTableDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernTableDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernTableDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
