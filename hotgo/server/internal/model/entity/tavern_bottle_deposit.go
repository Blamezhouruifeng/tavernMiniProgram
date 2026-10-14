// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernBottleDeposit is the golang structure for table tavern_bottle_deposit.
type TavernBottleDeposit struct {
	Id            int64       `json:"id"            orm:"id"              description:"主键"`
	StoreId       int64       `json:"storeId"       orm:"store_id"        description:"门店ID"`
	DepositNo     string      `json:"depositNo"     orm:"deposit_no"      description:"存酒单号"`
	LabelCode     string      `json:"labelCode"     orm:"label_code"      description:"瓶签码（扫码认领）"`
	MemberId      int64       `json:"memberId"      orm:"member_id"       description:"会员ID"`
	MemberName    string      `json:"memberName"    orm:"member_name"     description:"会员姓名快照"`
	PhoneTail     string      `json:"phoneTail"     orm:"phone_tail"      description:"手机号后四位（瓶签用）"`
	SourceType    int         `json:"sourceType"    orm:"source_type"     description:"来源：1自带代存 2购酒寄存 3退酒入存"`
	SourceOrderId int64       `json:"sourceOrderId" orm:"source_order_id" description:"来源订单ID（购酒/退酒入存）"`
	SessionId     int64       `json:"sessionId"     orm:"session_id"      description:"来源桌台会话ID"`
	GoodsName     string      `json:"goodsName"     orm:"goods_name"      description:"酒品名称"`
	Spec          string      `json:"spec"          orm:"spec"            description:"规格，如 750ml"`
	Degree        string      `json:"degree"        orm:"degree"          description:"酒精度"`
	TotalRatio    float64     `json:"totalRatio"    orm:"total_ratio"     description:"存入时剩余比例(%)"`
	RemainRatio   float64     `json:"remainRatio"   orm:"remain_ratio"    description:"当前剩余比例(%)"`
	SlotId        int64       `json:"slotId"        orm:"slot_id"         description:"储位ID"`
	SlotCode      string      `json:"slotCode"      orm:"slot_code"       description:"储位号快照"`
	Photo         string      `json:"photo"         orm:"photo"           description:"封签照片"`
	Status        int         `json:"status"        orm:"status"          description:"状态：1在库 2已取完 3已过期 4已转让 5报损"`
	DepositAt     *gtime.Time `json:"depositAt"     orm:"deposit_at"      description:"存入时间"`
	ExpireAt      *gtime.Time `json:"expireAt"      orm:"expire_at"       description:"到期时间"`
	Remark        string      `json:"remark"        orm:"remark"          description:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"      description:"删除时间"`
	CreatedBy     int64       `json:"createdBy"     orm:"created_by"      description:"创建者（店员）"`
	UpdatedBy     int64       `json:"updatedBy"     orm:"updated_by"      description:"更新者"`
}
