// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernMemberDao is the data access object for the table hg_tavern_member.
type TavernMemberDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  TavernMemberColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// TavernMemberColumns defines and stores column names for the table hg_tavern_member.
type TavernMemberColumns struct {
	Id          string // 主键
	StoreId     string // 门店ID（0为全店通用）
	Openid      string // 微信openid
	Unionid     string // 微信unionid
	Nickname    string // 昵称
	Avatar      string // 头像
	Phone       string // 手机号（存酒取酒凭证）
	RealName    string // 真实姓名
	IsStaff     string // 是否员工：1是 0否
	StaffRole   string // 员工角色：waiter服务员/bar吧台/kitchen后厨/manager店长
	StaffStatus string // 员工状态：1在职 2停用
	Points      string // 积分
	Balance     string // 储值余额
	TotalAmount string // 累计消费
	LastLoginAt string // 最后登录时间
	Status      string // 状态：1正常 2禁用
	Remark      string // 备注
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
	CreatedBy   string // 创建者
	UpdatedBy   string // 更新者
}

// tavernMemberColumns holds the columns for the table hg_tavern_member.
var tavernMemberColumns = TavernMemberColumns{
	Id:          "id",
	StoreId:     "store_id",
	Openid:      "openid",
	Unionid:     "unionid",
	Nickname:    "nickname",
	Avatar:      "avatar",
	Phone:       "phone",
	RealName:    "real_name",
	IsStaff:     "is_staff",
	StaffRole:   "staff_role",
	StaffStatus: "staff_status",
	Points:      "points",
	Balance:     "balance",
	TotalAmount: "total_amount",
	LastLoginAt: "last_login_at",
	Status:      "status",
	Remark:      "remark",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
}

// NewTavernMemberDao creates and returns a new DAO object for table data access.
func NewTavernMemberDao(handlers ...gdb.ModelHandler) *TavernMemberDao {
	return &TavernMemberDao{
		group:    "default",
		table:    "hg_tavern_member",
		columns:  tavernMemberColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernMemberDao) Columns() TavernMemberColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernMemberDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
