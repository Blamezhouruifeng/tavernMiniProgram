// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernCouponMemberDao is the data access object for the table hg_tavern_coupon_member.
type TavernCouponMemberDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  TavernCouponMemberColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// TavernCouponMemberColumns defines and stores column names for the table hg_tavern_coupon_member.
type TavernCouponMemberColumns struct {
	Id        string // 主键
	StoreId   string // 门店ID
	CouponId  string // 券模板ID
	MemberId  string // 会员ID
	Code      string // 券码
	Source    string // 来源：1新人礼 2活动发放 3后台定向
	Status    string // 状态：1未使用 2已使用 3已过期
	OrderId   string // 使用的订单ID
	Amount    string // 抵扣金额快照
	ExpireAt  string // 过期时间
	UsedAt    string // 使用时间
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// tavernCouponMemberColumns holds the columns for the table hg_tavern_coupon_member.
var tavernCouponMemberColumns = TavernCouponMemberColumns{
	Id:        "id",
	StoreId:   "store_id",
	CouponId:  "coupon_id",
	MemberId:  "member_id",
	Code:      "code",
	Source:    "source",
	Status:    "status",
	OrderId:   "order_id",
	Amount:    "amount",
	ExpireAt:  "expire_at",
	UsedAt:    "used_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTavernCouponMemberDao creates and returns a new DAO object for table data access.
func NewTavernCouponMemberDao(handlers ...gdb.ModelHandler) *TavernCouponMemberDao {
	return &TavernCouponMemberDao{
		group:    "default",
		table:    "hg_tavern_coupon_member",
		columns:  tavernCouponMemberColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernCouponMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernCouponMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernCouponMemberDao) Columns() TavernCouponMemberColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernCouponMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernCouponMemberDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernCouponMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
