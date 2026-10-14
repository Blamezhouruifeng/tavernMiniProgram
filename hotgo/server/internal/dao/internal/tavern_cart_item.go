// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernCartItemDao is the data access object for the table hg_tavern_cart_item.
type TavernCartItemDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  TavernCartItemColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// TavernCartItemColumns defines and stores column names for the table hg_tavern_cart_item.
type TavernCartItemColumns struct {
	Id        string // 主键
	StoreId   string // 门店ID
	SessionId string // 会话ID
	MemberId  string // 加购会员ID（归属）
	GoodsId   string // 商品ID
	SkuId     string // 规格ID
	Qty       string // 数量
	Remark    string // 备注（忌口/做法）
	OrderId   string // 下单后回填订单ID
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// tavernCartItemColumns holds the columns for the table hg_tavern_cart_item.
var tavernCartItemColumns = TavernCartItemColumns{
	Id:        "id",
	StoreId:   "store_id",
	SessionId: "session_id",
	MemberId:  "member_id",
	GoodsId:   "goods_id",
	SkuId:     "sku_id",
	Qty:       "qty",
	Remark:    "remark",
	OrderId:   "order_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewTavernCartItemDao creates and returns a new DAO object for table data access.
func NewTavernCartItemDao(handlers ...gdb.ModelHandler) *TavernCartItemDao {
	return &TavernCartItemDao{
		group:    "default",
		table:    "hg_tavern_cart_item",
		columns:  tavernCartItemColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernCartItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernCartItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernCartItemDao) Columns() TavernCartItemColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernCartItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernCartItemDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernCartItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
