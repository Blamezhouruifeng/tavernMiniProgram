package tavernin

import (
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

/* ============ 后台管理（HotGo 后台）入参/出参 ============ */

type AdminStoreViewModel struct {
	Id               int64  `json:"id"`
	Name             string `json:"name"`
	Brief            string `json:"brief"`
	Address          string `json:"address"`
	Phone            string `json:"phone"`
	BusinessOpen     string `json:"businessOpen"`
	BusinessClose    string `json:"businessClose"`
	LastOrderGap     int    `json:"lastOrderGap"`
	PayFirst         bool   `json:"payFirst"`
	PayExpireMinutes int    `json:"payExpireMinutes"`
}

type AdminStoreEditInp struct {
	Id               int64  `json:"id"`
	Name             string `json:"name" v:"required#门店名称不能为空"`
	Brief            string `json:"brief"`
	Address          string `json:"address"`
	Phone            string `json:"phone"`
	BusinessOpen     string `json:"businessOpen"`
	BusinessClose    string `json:"businessClose"`
	LastOrderGap     int    `json:"lastOrderGap"`
	PayFirst         bool   `json:"payFirst"`
	PayExpireMinutes int    `json:"payExpireMinutes"`
}

/* -------- 区域 -------- */

type AdminAreaListInp struct {
	form.PageReq
	Name string `json:"name"`
}

type AdminAreaModel struct {
	Id        int64       `json:"id"`
	StoreId   int64       `json:"storeId"`
	Name      string      `json:"name"`
	Status    int         `json:"status"`
	Sort      int         `json:"sort"`
	Remark    string      `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

type AdminAreaEditInp struct {
	Id     int64  `json:"id"`
	Name   string `json:"name" v:"required#区域名称不能为空"`
	Status int    `json:"status"`
	Sort   int    `json:"sort"`
	Remark string `json:"remark"`
}

type AdminDeleteInp struct {
	Ids []int64 `json:"ids"`
}

/* -------- 桌位 -------- */

type AdminTableListInp struct {
	form.PageReq
	Name   string `json:"name"`
	AreaId int64  `json:"areaId"`
}

type AdminTableModel struct {
	Id         int64       `json:"id"`
	AreaId     int64       `json:"areaId"`
	AreaName   string      `json:"areaName"`
	Name       string      `json:"name"`
	Capacity   int         `json:"capacity"`
	Scene      string      `json:"qrcodeScene"`
	SeatStatus int         `json:"seatStatus"`
	Status     int         `json:"status"`
	Sort       int         `json:"sort"`
	CreatedAt  *gtime.Time `json:"createdAt"`
}

type AdminTableEditInp struct {
	Id       int64  `json:"id"`
	AreaId   int64  `json:"areaId"`
	Name     string `json:"name" v:"required#桌号不能为空"`
	Capacity int    `json:"capacity"`
	Status   int    `json:"status"`
	Sort     int    `json:"sort"`
}

/* -------- 分类 -------- */

type AdminCategoryListInp struct {
	form.PageReq
	Name string `json:"name"`
}

type AdminCategoryModel struct {
	Id        int64       `json:"id"`
	Name      string      `json:"name"`
	Icon      string      `json:"icon"`
	TimeSlots string      `json:"timeSlots"`
	Status    int         `json:"status"`
	Sort      int         `json:"sort"`
	CreatedAt *gtime.Time `json:"createdAt"`
}

type AdminCategoryEditInp struct {
	Id        int64  `json:"id"`
	Name      string `json:"name" v:"required#分类名称不能为空"`
	Icon      string `json:"icon"`
	TimeSlots string `json:"timeSlots"`
	Status    int    `json:"status"`
	Sort      int    `json:"sort"`
}

/* -------- 商品（含规格） -------- */

type AdminGoodsListInp struct {
	form.PageReq
	Name       string `json:"name"`
	CategoryId int64  `json:"categoryId"`
	Status     int    `json:"status"`
}

type AdminGoodsSkuModel struct {
	Id            int64   `json:"id"`
	GoodsId       int64   `json:"goodsId"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Stock         int     `json:"stock"`
	UnitCups      float64 `json:"unitCups"`
	IsDepositUnit bool    `json:"isDepositUnit"`
	Status        int     `json:"status"`
	Sort          int     `json:"sort"`
}

type AdminGoodsModel struct {
	Id           int64                 `json:"id"`
	CategoryId   int64                 `json:"categoryId"`
	CategoryName string                `json:"categoryName"`
	Name         string                `json:"name"`
	Cover        string                `json:"cover"`
	Brief        string                `json:"brief"`
	Price        float64               `json:"price"`
	Tags         string                `json:"tags"`
	IsWine       bool                  `json:"isWine"`
	OnlyDineIn   bool                  `json:"onlyDineIn"`
	IsCleared    bool                  `json:"isCleared"`
	MakeMinutes  int                   `json:"makeMinutes"`
	Status       int                   `json:"status"`
	Sort         int                   `json:"sort"`
	CreatedAt    *gtime.Time           `json:"createdAt"`
	Skus         []*AdminGoodsSkuModel `json:"skus"`
}

type AdminGoodsViewInp struct {
	Id int64 `json:"id" v:"required#商品ID不能为空"`
}

type AdminGoodsEditInp struct {
	Id          int64                 `json:"id"`
	CategoryId  int64                 `json:"categoryId" v:"required#请选择分类"`
	Name        string                `json:"name" v:"required#商品名称不能为空"`
	Cover       string                `json:"cover"`
	Brief       string                `json:"brief"`
	Tags        string                `json:"tags"`
	IsWine      bool                  `json:"isWine"`
	OnlyDineIn  bool                  `json:"onlyDineIn"`
	IsCleared   bool                  `json:"isCleared"`
	MakeMinutes int                   `json:"makeMinutes"`
	Status      int                   `json:"status"`
	Sort        int                   `json:"sort"`
	Skus        []*AdminGoodsSkuModel `json:"skus"`
}

type AdminGoodsStatusInp struct {
	Id     int64 `json:"id" v:"required#商品ID不能为空"`
	Status int   `json:"status"`
}

/* -------- 订单查询 -------- */

type AdminOrderListInp struct {
	form.PageReq
	OrderNo   string        `json:"orderNo"`
	TableName string        `json:"tableName"`
	PayStatus int           `json:"payStatus"`
	DateRange []*gtime.Time `json:"dateRange"`
}

type AdminOrderModel struct {
	Id           int64       `json:"id"`
	OrderNo      string      `json:"orderNo"`
	TableName    string      `json:"tableName"`
	OrderSeq     int         `json:"orderSeq"`
	MemberName   string      `json:"memberName"`
	PayStatus    int         `json:"payStatus"`
	MakingStatus int         `json:"makingStatus"`
	RefundStatus int         `json:"refundStatus"`
	PayAmount    float64     `json:"payAmount"`
	RefundAmount float64     `json:"refundAmount"`
	PayChannel   string      `json:"payChannel"`
	PayTime      *gtime.Time `json:"payTime"`
	CreatedAt    *gtime.Time `json:"createdAt"`
}

/* -------- 存酒台账 -------- */

type AdminBottleListInp struct {
	form.PageReq
	Keyword    string `json:"keyword"`
	Status     int    `json:"status"`
	SourceType int    `json:"sourceType"`
}

type AdminAbnormalOrderModel struct {
	Id           int64       `json:"id"`
	OrderNo      string      `json:"orderNo"`
	TableName    string      `json:"tableName"`
	MemberName   string      `json:"memberName"`
	PayStatus    int         `json:"payStatus"`
	MakingStatus int         `json:"makingStatus"`
	RefundStatus int         `json:"refundStatus"`
	PayAmount    float64     `json:"payAmount"`
	PayTime      *gtime.Time `json:"payTime"`
	ExpireAt     *gtime.Time `json:"expireAt"`
	CreatedAt    *gtime.Time `json:"createdAt"`
	Issue        string      `json:"issue"`
}

type AdminReissueInp struct {
	OrderNo string `json:"orderNo" v:"required#订单号不能为空"`
}

type AdminReconcileInp struct {
	Date string `json:"date"`
}

type AdminReconcileItem struct {
	OrderNo    string      `json:"orderNo"`
	PayAmount  float64     `json:"payAmount"`
	PayChannel string      `json:"payChannel"`
	PayTime    *gtime.Time `json:"payTime"`
	HasRecord  bool        `json:"hasRecord"`
	RecordAmt  float64     `json:"recordAmt"`
}

type AdminReconcileModel struct {
	Date            string                `json:"date"`
	OrderCount      int                   `json:"orderCount"`
	OrderAmount     float64               `json:"orderAmount"`
	PayRecordAmount float64               `json:"payRecordAmount"`
	RefundAmount    float64               `json:"refundAmount"`
	Diff            float64               `json:"diff"`
	Items           []*AdminReconcileItem `json:"items"`
}

type AdminBottleUpdateInp struct {
	Id          int64   `json:"id" v:"required#存酒ID不能为空"`
	Status      int     `json:"status"`
	RemainRatio float64 `json:"remainRatio"`
	SlotCode    string  `json:"slotCode"`
	ExpireAt    string  `json:"expireAt"`
	Remark      string  `json:"remark"`
}
