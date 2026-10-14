// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernOrderItemDao is the data access object for the table hg_tavern_order_item.
type TavernOrderItemDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  TavernOrderItemColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// TavernOrderItemColumns defines and stores column names for the table hg_tavern_order_item.
type TavernOrderItemColumns struct {
	Id          string // 主键
	StoreId     string // 门店ID
	OrderId     string // 订单ID
	GoodsId     string // 商品ID
	SkuId       string // 规格ID
	GoodsName   string // 商品名称快照
	SkuName     string // 规格名称快照
	UnitPrice   string // 单价
	Qty         string // 数量
	Amount      string // 小计
	IsWine      string // 是否酒水（出票分单用）
	MakerStatus string // 行状态：0待制作 1制作中 2已上齐
	ServedAt    string // 上齐时间
	Remark      string // 备注
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// tavernOrderItemColumns holds the columns for the table hg_tavern_order_item.
var tavernOrderItemColumns = TavernOrderItemColumns{
	Id:          "id",
	StoreId:     "store_id",
	OrderId:     "order_id",
	GoodsId:     "goods_id",
	SkuId:       "sku_id",
	GoodsName:   "goods_name",
	SkuName:     "sku_name",
	UnitPrice:   "unit_price",
	Qty:         "qty",
	Amount:      "amount",
	IsWine:      "is_wine",
	MakerStatus: "maker_status",
	ServedAt:    "served_at",
	Remark:      "remark",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewTavernOrderItemDao creates and returns a new DAO object for table data access.
func NewTavernOrderItemDao(handlers ...gdb.ModelHandler) *TavernOrderItemDao {
	return &TavernOrderItemDao{
		group:    "default",
		table:    "hg_tavern_order_item",
		columns:  tavernOrderItemColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernOrderItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernOrderItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernOrderItemDao) Columns() TavernOrderItemColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernOrderItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernOrderItemDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernOrderItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
