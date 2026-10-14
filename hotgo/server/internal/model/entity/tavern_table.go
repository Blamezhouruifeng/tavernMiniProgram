// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernTable is the golang structure for table tavern_table.
type TavernTable struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键"`
	StoreId     int64       `json:"storeId"     orm:"store_id"     description:"门店ID"`
	AreaId      int64       `json:"areaId"      orm:"area_id"      description:"区域ID"`
	Name        string      `json:"name"        orm:"name"         description:"桌号"`
	Capacity    int         `json:"capacity"    orm:"capacity"     description:"可容纳人数"`
	Qrcode      string      `json:"qrcode"      orm:"qrcode"       description:"桌码文件路径"`
	QrcodeScene string      `json:"qrcodeScene" orm:"qrcode_scene" description:"桌码scene参数"`
	SeatStatus  int         `json:"seatStatus"  orm:"seat_status"  description:"桌位状态：1空闲 2用餐中 3待清台"`
	Status      int         `json:"status"      orm:"status"       description:"状态：1正常 2禁用"`
	Sort        int         `json:"sort"        orm:"sort"         description:"排序"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:"更新者"`
}
