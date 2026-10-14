// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernGoodsDao is the data access object for the table hg_tavern_goods.
type TavernGoodsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TavernGoodsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TavernGoodsColumns defines and stores column names for the table hg_tavern_goods.
type TavernGoodsColumns struct {
	Id          string // 主键
	StoreId     string // 门店ID
	CategoryId  string // 分类ID
	Name        string // 商品名称
	Cover       string // 封面图
	Brief       string // 简介
	Price       string // 最低售价（展示用）
	Tags        string // 标签，逗号分隔：招牌,热销,新品
	IsWine      string // 是否酒水：1是 0否
	OnlyDineIn  string // 仅堂食：1是 0否
	IsCleared   string // 是否沽清：1是 0否
	MakeMinutes string // 预计制作分钟数
	Status      string // 状态：1上架 2下架
	Sort        string // 排序
	Remark      string // 备注
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
	CreatedBy   string // 创建者
	UpdatedBy   string // 更新者
}

// tavernGoodsColumns holds the columns for the table hg_tavern_goods.
var tavernGoodsColumns = TavernGoodsColumns{
	Id:          "id",
	StoreId:     "store_id",
	CategoryId:  "category_id",
	Name:        "name",
	Cover:       "cover",
	Brief:       "brief",
	Price:       "price",
	Tags:        "tags",
	IsWine:      "is_wine",
	OnlyDineIn:  "only_dine_in",
	IsCleared:   "is_cleared",
	MakeMinutes: "make_minutes",
	Status:      "status",
	Sort:        "sort",
	Remark:      "remark",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
}

// NewTavernGoodsDao creates and returns a new DAO object for table data access.
func NewTavernGoodsDao(handlers ...gdb.ModelHandler) *TavernGoodsDao {
	return &TavernGoodsDao{
		group:    "default",
		table:    "hg_tavern_goods",
		columns:  tavernGoodsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernGoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernGoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernGoodsDao) Columns() TavernGoodsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernGoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernGoodsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernGoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
