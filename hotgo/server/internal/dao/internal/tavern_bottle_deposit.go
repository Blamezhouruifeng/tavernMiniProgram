// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TavernBottleDepositDao is the data access object for the table hg_tavern_bottle_deposit.
type TavernBottleDepositDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  TavernBottleDepositColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// TavernBottleDepositColumns defines and stores column names for the table hg_tavern_bottle_deposit.
type TavernBottleDepositColumns struct {
	Id            string // 主键
	StoreId       string // 门店ID
	DepositNo     string // 存酒单号
	LabelCode     string // 瓶签码（扫码认领）
	MemberId      string // 会员ID
	MemberName    string // 会员姓名快照
	PhoneTail     string // 手机号后四位（瓶签用）
	SourceType    string // 来源：1自带代存 2购酒寄存 3退酒入存
	SourceOrderId string // 来源订单ID（购酒/退酒入存）
	SessionId     string // 来源桌台会话ID
	GoodsName     string // 酒品名称
	Spec          string // 规格，如 750ml
	Degree        string // 酒精度
	TotalRatio    string // 存入时剩余比例(%)
	RemainRatio   string // 当前剩余比例(%)
	SlotId        string // 储位ID
	SlotCode      string // 储位号快照
	Photo         string // 封签照片
	Status        string // 状态：1在库 2已取完 3已过期 4已转让 5报损
	DepositAt     string // 存入时间
	ExpireAt      string // 到期时间
	Remark        string // 备注
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 删除时间
	CreatedBy     string // 创建者（店员）
	UpdatedBy     string // 更新者
}

// tavernBottleDepositColumns holds the columns for the table hg_tavern_bottle_deposit.
var tavernBottleDepositColumns = TavernBottleDepositColumns{
	Id:            "id",
	StoreId:       "store_id",
	DepositNo:     "deposit_no",
	LabelCode:     "label_code",
	MemberId:      "member_id",
	MemberName:    "member_name",
	PhoneTail:     "phone_tail",
	SourceType:    "source_type",
	SourceOrderId: "source_order_id",
	SessionId:     "session_id",
	GoodsName:     "goods_name",
	Spec:          "spec",
	Degree:        "degree",
	TotalRatio:    "total_ratio",
	RemainRatio:   "remain_ratio",
	SlotId:        "slot_id",
	SlotCode:      "slot_code",
	Photo:         "photo",
	Status:        "status",
	DepositAt:     "deposit_at",
	ExpireAt:      "expire_at",
	Remark:        "remark",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	CreatedBy:     "created_by",
	UpdatedBy:     "updated_by",
}

// NewTavernBottleDepositDao creates and returns a new DAO object for table data access.
func NewTavernBottleDepositDao(handlers ...gdb.ModelHandler) *TavernBottleDepositDao {
	return &TavernBottleDepositDao{
		group:    "default",
		table:    "hg_tavern_bottle_deposit",
		columns:  tavernBottleDepositColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TavernBottleDepositDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TavernBottleDepositDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TavernBottleDepositDao) Columns() TavernBottleDepositColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TavernBottleDepositDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TavernBottleDepositDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TavernBottleDepositDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
