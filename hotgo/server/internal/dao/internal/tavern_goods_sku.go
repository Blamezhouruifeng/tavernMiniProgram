// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernGoodsSkuDao is the data access object for the table hg_tavern_goods_sku.
type TavernGoodsSkuDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  TavernGoodsSkuColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// TavernGoodsSkuColumns defines and stores column names for the table hg_tavern_goods_sku.
type TavernGoodsSkuColumns struct {
	Id            string // 主键
	StoreId       string // 门店ID
	GoodsId       string // 商品ID
	Name          string // 规格名称：杯/瓶/扎
	Price         string // 售价
	Stock         string // 库存，-1为不限量
	UnitCups      string // 换算杯数（1瓶=N杯）
	IsDepositUnit string // 是否为整瓶可寄存单位：1是 0否
	Status        string // 状态：1正常 2禁用
	Sort          string // 排序
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 删除时间
}

// tavernGoodsSkuColumns holds the columns for the table hg_tavern_goods_sku.
var tavernGoodsSkuColumns = TavernGoodsSkuColumns{
	Id:            "id",
	StoreId:       "store_id",
	GoodsId:       "goods_id",
	Name:          "name",
	Price:         "price",
	Stock:         "stock",
	UnitCups:      "unit_cups",
	IsDepositUnit: "is_deposit_unit",
	Status:        "status",
	Sort:          "sort",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewTavernGoodsSkuDao creates and returns a new DAO object for table data access.
func NewTavernGoodsSkuDao(handlers ...gdb.ModelHandler) *TavernGoodsSkuDao {
	return &TavernGoodsSkuDao{
		group:    "default",
		table:    "hg_tavern_goods_sku",
		columns:  tavernGoodsSkuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernGoodsSkuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernGoodsSkuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernGoodsSkuDao) Columns() TavernGoodsSkuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernGoodsSkuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernGoodsSkuDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernGoodsSkuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
