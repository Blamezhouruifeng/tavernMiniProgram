// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernBottleDeposit is the golang structure of table hg_tavern_bottle_deposit for DAO operations like Where/Data.
type TavernBottleDeposit struct {
	g.Meta        `orm:"table:hg_tavern_bottle_deposit, do:true"`
	Id            any         // 主键
	StoreId       any         // 门店ID
	DepositNo     any         // 存酒单号
	LabelCode     any         // 瓶签码（扫码认领）
	MemberId      any         // 会员ID
	MemberName    any         // 会员姓名快照
	PhoneTail     any         // 手机号后四位（瓶签用）
	SourceType    any         // 来源：1自带代存 2购酒寄存 3退酒入存
	SourceOrderId any         // 来源订单ID（购酒/退酒入存）
	SessionId     any         // 来源桌台会话ID
	GoodsName     any         // 酒品名称
	Spec          any         // 规格，如 750ml
	Degree        any         // 酒精度
	TotalRatio    any         // 存入时剩余比例(%)
	RemainRatio   any         // 当前剩余比例(%)
	SlotId        any         // 储位ID
	SlotCode      any         // 储位号快照
	Photo         any         // 封签照片
	Status        any         // 状态：1在库 2已取完 3已过期 4已转让 5报损
	DepositAt     *gtime.Time // 存入时间
	ExpireAt      *gtime.Time // 到期时间
	Remark        any         // 备注
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
	CreatedBy     any         // 创建者（店员）
	UpdatedBy     any         // 更新者
}
