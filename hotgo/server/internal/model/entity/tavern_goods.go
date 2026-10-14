// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TavernGoods is the golang structure for table tavern_goods.
type TavernGoods struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键"`
	StoreId     int64       `json:"storeId"     orm:"store_id"     description:"门店ID"`
	CategoryId  int64       `json:"categoryId"  orm:"category_id"  description:"分类ID"`
	Name        string      `json:"name"        orm:"name"         description:"商品名称"`
	Cover       string      `json:"cover"       orm:"cover"        description:"封面图"`
	Brief       string      `json:"brief"       orm:"brief"        description:"简介"`
	Price       float64     `json:"price"       orm:"price"        description:"最低售价（展示用）"`
	Tags        string      `json:"tags"        orm:"tags"         description:"标签，逗号分隔：招牌,热销,新品"`
	IsWine      int         `json:"isWine"      orm:"is_wine"      description:"是否酒水：1是 0否"`
	OnlyDineIn  int         `json:"onlyDineIn"  orm:"only_dine_in" description:"仅堂食：1是 0否"`
	IsCleared   int         `json:"isCleared"   orm:"is_cleared"   description:"是否沽清：1是 0否"`
	MakeMinutes int         `json:"makeMinutes" orm:"make_minutes" description:"预计制作分钟数"`
	Status      int         `json:"status"      orm:"status"       description:"状态：1上架 2下架"`
	Sort        int         `json:"sort"        orm:"sort"         description:"排序"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:"更新者"`
}
