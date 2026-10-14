// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernCouponDao is the data access object for the table hg_tavern_coupon.
type TavernCouponDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  TavernCouponColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// TavernCouponColumns defines and stores column names for the table hg_tavern_coupon.
type TavernCouponColumns struct {
	Id             string // 主键
	StoreId        string // 门店ID
	Name           string // 券名称
	Type           string // 类型：1满减 2折扣
	Amount         string // 满减金额
	MinAmount      string // 使用门槛（订单满多少可用）
	Discount       string // 折扣（如 0.88 表示八八折）
	ValidDays      string // 领取后有效天数
	TotalLimit     string // 发放总量，-1不限
	PerMemberLimit string // 每人限领
	IsNewcomer     string // 是否新人礼（注册自动发放）
	Status         string // 状态：1启用 2停用
	Sort           string // 排序
	Remark         string // 备注
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间
	CreatedBy      string // 创建者
	UpdatedBy      string // 更新者
}

// tavernCouponColumns holds the columns for the table hg_tavern_coupon.
var tavernCouponColumns = TavernCouponColumns{
	Id:             "id",
	StoreId:        "store_id",
	Name:           "name",
	Type:           "type",
	Amount:         "amount",
	MinAmount:      "min_amount",
	Discount:       "discount",
	ValidDays:      "valid_days",
	TotalLimit:     "total_limit",
	PerMemberLimit: "per_member_limit",
	IsNewcomer:     "is_newcomer",
	Status:         "status",
	Sort:           "sort",
	Remark:         "remark",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
	CreatedBy:      "created_by",
	UpdatedBy:      "updated_by",
}

// NewTavernCouponDao creates and returns a new DAO object for table data access.
func NewTavernCouponDao(handlers ...gdb.ModelHandler) *TavernCouponDao {
	return &TavernCouponDao{
		group:    "default",
		table:    "hg_tavern_coupon",
		columns:  tavernCouponColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernCouponDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernCouponDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernCouponDao) Columns() TavernCouponColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernCouponDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernCouponDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernCouponDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
