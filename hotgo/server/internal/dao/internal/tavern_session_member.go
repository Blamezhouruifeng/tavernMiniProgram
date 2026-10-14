// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernSessionMemberDao is the data access object for the table hg_tavern_session_member.
type TavernSessionMemberDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  TavernSessionMemberColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// TavernSessionMemberColumns defines and stores column names for the table hg_tavern_session_member.
type TavernSessionMemberColumns struct {
	Id        string // 主键
	StoreId   string // 门店ID
	SessionId string // 会话ID
	MemberId  string // 会员ID
	Nickname  string // 昵称快照
	Avatar    string // 头像快照
	IsOwner   string // 是否开台人：1是 0否
	JoinAt    string // 加入时间
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// tavernSessionMemberColumns holds the columns for the table hg_tavern_session_member.
var tavernSessionMemberColumns = TavernSessionMemberColumns{
	Id:        "id",
	StoreId:   "store_id",
	SessionId: "session_id",
	MemberId:  "member_id",
	Nickname:  "nickname",
	Avatar:    "avatar",
	IsOwner:   "is_owner",
	JoinAt:    "join_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewTavernSessionMemberDao creates and returns a new DAO object for table data access.
func NewTavernSessionMemberDao(handlers ...gdb.ModelHandler) *TavernSessionMemberDao {
	return &TavernSessionMemberDao{
		group:    "default",
		table:    "hg_tavern_session_member",
		columns:  tavernSessionMemberColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernSessionMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernSessionMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernSessionMemberDao) Columns() TavernSessionMemberColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernSessionMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernSessionMemberDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernSessionMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
