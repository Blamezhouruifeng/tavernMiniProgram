// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernTableSessionDao is the data access object for the table hg_tavern_table_session.
type TavernTableSessionDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  TavernTableSessionColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// TavernTableSessionColumns defines and stores column names for the table hg_tavern_table_session.
type TavernTableSessionColumns struct {
	Id               string // 主键
	StoreId          string // 门店ID
	TableId          string // 桌位ID
	SessionNo        string // 会话号
	OwnerMemberId    string // 开台会员ID
	GuestCount       string // 就餐人数
	MemberCount      string // 参与会员数
	OrderCount       string // 已支付订单数
	TotalPaidAmount  string // 累计已付金额
	GuestTotalAmount string // 全桌合计金额
	Status           string // 会话状态：1进行中 2已归档
	OpenAt           string // 开台时间
	CloseAt          string // 关台时间
	Remark           string // 备注
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
	DeletedAt        string // 删除时间
	CreatedBy        string // 创建者
	UpdatedBy        string // 更新者
}

// tavernTableSessionColumns holds the columns for the table hg_tavern_table_session.
var tavernTableSessionColumns = TavernTableSessionColumns{
	Id:               "id",
	StoreId:          "store_id",
	TableId:          "table_id",
	SessionNo:        "session_no",
	OwnerMemberId:    "owner_member_id",
	GuestCount:       "guest_count",
	MemberCount:      "member_count",
	OrderCount:       "order_count",
	TotalPaidAmount:  "total_paid_amount",
	GuestTotalAmount: "guest_total_amount",
	Status:           "status",
	OpenAt:           "open_at",
	CloseAt:          "close_at",
	Remark:           "remark",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
	CreatedBy:        "created_by",
	UpdatedBy:        "updated_by",
}

// NewTavernTableSessionDao creates and returns a new DAO object for table data access.
func NewTavernTableSessionDao(handlers ...gdb.ModelHandler) *TavernTableSessionDao {
	return &TavernTableSessionDao{
		group:    "default",
		table:    "hg_tavern_table_session",
		columns:  tavernTableSessionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernTableSessionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernTableSessionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernTableSessionDao) Columns() TavernTableSessionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernTableSessionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernTableSessionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernTableSessionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
