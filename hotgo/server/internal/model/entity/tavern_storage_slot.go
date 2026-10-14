// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernStorageSlot is the golang structure for table tavern_storage_slot.
type TavernStorageSlot struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`
	StoreId   int64       `json:"storeId"   orm:"store_id"   description:"门店ID"`
	Area      string      `json:"area"      orm:"area"       description:"储区，如 A"`
	Layer     string      `json:"layer"     orm:"layer"      description:"层，如 2"`
	Cell      string      `json:"cell"      orm:"cell"       description:"格位，如 3"`
	Code      string      `json:"code"      orm:"code"       description:"储位号，如 A-2-3"`
	IsUsed    int         `json:"isUsed"    orm:"is_used"    description:"是否占用：1是 0否"`
	Status    int         `json:"status"    orm:"status"     description:"状态：1正常 2禁用"`
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
