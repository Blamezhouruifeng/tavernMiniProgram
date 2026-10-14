// Package tavernin 酒馆小程序业务入参/出参
// 说明：与 hg_tavern_* 表结构对齐，同时作为前端接口契约（结构体字段名与 JSON 一致）
package tavernin

import (
	"github.com/gogf/gf/v2/os/gtime"
)

/* ============ 门店 ============ */

type StoreInfoModel struct {
	Id               int64  `json:"id"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	Phone            string `json:"phone"`
	BusinessOpen     string `json:"businessOpen"`
	BusinessClose    string `json:"businessClose"`
	LastOrderGap     int    `json:"lastOrderGap"`
	PayFirst         bool   `json:"payFirst"`
	PayExpireMinutes int    `json:"payExpireMinutes"`
}

/* ============ 会员 ============ */

type MemberLoginInp struct {
	Code     string `json:"code" v:"required#登录凭证不能为空"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
}

type MemberInfoModel struct {
	Id        int64   `json:"id"`
	Nickname  string  `json:"nickname"`
	Avatar    string  `json:"avatar"`
	Phone     string  `json:"phone"`
	IsStaff   bool    `json:"isStaff"`
	StaffRole string  `json:"staffRole"`
	Points    int     `json:"points"`
	Balance   float64 `json:"balance"`
}

type MemberLoginModel struct {
	Token     string           `json:"token"`
	ExpiresIn int64            `json:"expiresIn"`
	Member    *MemberInfoModel `json:"member"`
}

type MemberSwitchStaffInp struct {
	IsStaff bool   `json:"isStaff"`
	Role    string `json:"role"`
}

/* ============ 菜单 ============ */

type MenuSkuModel struct {
	Id            int64   `json:"id"`
	GoodsId       int64   `json:"goodsId"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Stock         int     `json:"stock"`
	UnitCups      float64 `json:"unitCups"`
	IsDepositUnit bool    `json:"isDepositUnit"`
}

type MenuGoodsModel struct {
	Id          int64           `json:"id"`
	CategoryId  int64           `json:"categoryId"`
	Name        string          `json:"name"`
	Cover       string          `json:"cover"`
	Brief       string          `json:"brief"`
	Price       float64         `json:"price"`
	Tags        []string        `json:"tags"`
	IsWine      bool            `json:"isWine"`
	OnlyDineIn  bool            `json:"onlyDineIn"`
	IsCleared   bool            `json:"isCleared"`
	MakeMinutes int             `json:"makeMinutes"`
	Status      int             `json:"status"`
	Skus        []*MenuSkuModel `json:"skus"`
}

type MenuCategoryModel struct {
	Id        int64             `json:"id"`
	Name      string            `json:"name"`
	Icon      string            `json:"icon"`
	TimeSlots string            `json:"timeSlots"`
	Goods     []*MenuGoodsModel `json:"goods"`
}

/* ============ 桌台会话 ============ */

type SessionOpenInp struct {
	TableNo string `json:"tableNo" v:"required#请提供桌号"`
}

type SessionMemberModel struct {
	MemberId int64       `json:"memberId"`
	Nickname string      `json:"nickname"`
	Avatar   string      `json:"avatar"`
	IsOwner  bool        `json:"isOwner"`
	JoinAt   *gtime.Time `json:"joinAt"`
}

type TableSessionModel struct {
	Id              int64                 `json:"id"`
	SessionNo       string                `json:"sessionNo"`
	TableId         int64                 `json:"tableId"`
	TableName       string                `json:"tableName"`
	Status          int                   `json:"status"`
	GuestCount      int                   `json:"guestCount"`
	OwnerMemberId   int64                 `json:"ownerMemberId"`
	OpenAt          *gtime.Time           `json:"openAt"`
	OrderCount      int                   `json:"orderCount"`
	TotalPaidAmount float64               `json:"totalPaidAmount"`
	Members         []*SessionMemberModel `json:"members"`
}

type SessionCloseInp struct {
	SessionId int64 `json:"sessionId" v:"required#会话ID不能为空"`
}

/* ============ 购物车 ============ */

type CartAddInp struct {
	SessionId int64  `json:"sessionId" v:"required#会话ID不能为空"`
	GoodsId   int64  `json:"goodsId" v:"required#商品ID不能为空"`
	SkuId     int64  `json:"skuId" v:"required#规格ID不能为空"`
	Qty       int    `json:"qty"`
	Remark    string `json:"remark"`
}

type CartUpdateInp struct {
	Id     int64  `json:"id" v:"required#购物车行ID不能为空"`
	Qty    int    `json:"qty"`
	Remark string `json:"remark"`
}

type CartRemoveInp struct {
	Id int64 `json:"id" v:"required#购物车行ID不能为空"`
}

type CartItemModel struct {
	Id         int64   `json:"id"`
	SessionId  int64   `json:"sessionId"`
	MemberId   int64   `json:"memberId"`
	MemberName string  `json:"memberName"`
	GoodsId    int64   `json:"goodsId"`
	SkuId      int64   `json:"skuId"`
	GoodsName  string  `json:"goodsName"`
	SkuName    string  `json:"skuName"`
	Price      float64 `json:"price"`
	Qty        int     `json:"qty"`
	Remark     string  `json:"remark"`
}

/* ============ 订单（先付制） ============ */

type OrderSubmitInp struct {
	SessionId int64   `json:"sessionId" v:"required#会话ID不能为空"`
	CartIds   []int64 `json:"cartIds" v:"required#请选择要下单的商品"`
	Remark    string  `json:"remark"`
}

type OrderPayInp struct {
	OrderNo string `json:"orderNo" v:"required#订单号不能为空"`
	Channel string `json:"channel"`
}

type OrderDetailInp struct {
	OrderNo string `json:"orderNo" v:"required#订单号不能为空"`
}

type OrderListInp struct {
	SessionId int64 `json:"sessionId"`
	Mine      bool  `json:"mine"`
}

type OrderItemModel struct {
	Id          int64       `json:"id"`
	OrderId     int64       `json:"orderId"`
	GoodsId     int64       `json:"goodsId"`
	SkuId       int64       `json:"skuId"`
	GoodsName   string      `json:"goodsName"`
	SkuName     string      `json:"skuName"`
	UnitPrice   float64     `json:"unitPrice"`
	Qty         int         `json:"qty"`
	Amount      float64     `json:"amount"`
	IsWine      bool        `json:"isWine"`
	MakerStatus int         `json:"makerStatus"`
	ServedAt    *gtime.Time `json:"servedAt"`
	Remark      string      `json:"remark"`
}

type OrderModel struct {
	Id             int64             `json:"id"`
	OrderNo        string            `json:"orderNo"`
	SessionId      int64             `json:"sessionId"`
	TableId        int64             `json:"tableId"`
	TableName      string            `json:"tableName"`
	OrderSeq       int               `json:"orderSeq"`
	MemberId       int64             `json:"memberId"`
	MemberName     string            `json:"memberName"`
	PayerMemberId  int64             `json:"payerMemberId"`
	PayStatus      int               `json:"payStatus"`
	MakingStatus   int               `json:"makingStatus"`
	RefundStatus   int               `json:"refundStatus"`
	PayChannel     string            `json:"payChannel"`
	TotalAmount    float64           `json:"totalAmount"`
	PayAmount      float64           `json:"payAmount"`
	DiscountAmount float64           `json:"discountAmount"`
	ExpireAt       *gtime.Time       `json:"expireAt"`
	PayTime        *gtime.Time       `json:"payTime"`
	CreatedAt      *gtime.Time       `json:"createdAt"`
	Items          []*OrderItemModel `json:"items"`
}

/* ============ 会话快照 ============ */

type OrderedAggModel struct {
	GoodsId   int64    `json:"goodsId"`
	SkuId     int64    `json:"skuId"`
	GoodsName string   `json:"goodsName"`
	SkuName   string   `json:"skuName"`
	Qty       int      `json:"qty"`
	ServedQty int      `json:"servedQty"`
	ByMembers []string `json:"byMembers"`
}

type SessionSnapshotModel struct {
	Session    *TableSessionModel    `json:"session"`
	Members    []*SessionMemberModel `json:"members"`
	Cart       []*CartItemModel      `json:"cart"`
	Orders     []*OrderModel         `json:"orders"`
	OrderedAgg []*OrderedAggModel    `json:"orderedAgg"`
}

/* ============ 存酒 ============ */

type BottleListInp struct {
	All bool `json:"all"`
}

type BottleModel struct {
	Id          int64       `json:"id"`
	DepositNo   string      `json:"depositNo"`
	LabelCode   string      `json:"labelCode"`
	MemberId    int64       `json:"memberId"`
	MemberName  string      `json:"memberName"`
	PhoneTail   string      `json:"phoneTail"`
	SourceType  int         `json:"sourceType"`
	GoodsName   string      `json:"goodsName"`
	Spec        string      `json:"spec"`
	Degree      string      `json:"degree"`
	RemainRatio float64     `json:"remainRatio"`
	SlotCode    string      `json:"slotCode"`
	Photo       string      `json:"photo"`
	Status      int         `json:"status"`
	DepositAt   *gtime.Time `json:"depositAt"`
	ExpireAt    *gtime.Time `json:"expireAt"`
}

type TakeCodeInp struct {
	DepositIds []int64 `json:"depositIds"`
}

type TakeCodeModel struct {
	Code          string  `json:"code"`
	ExpireSeconds int     `json:"expireSeconds"`
	DepositIds    []int64 `json:"depositIds"`
}

type DepositFromOrderInp struct {
	OrderId int64   `json:"orderId" v:"required#订单ID不能为空"`
	ItemIds []int64 `json:"itemIds" v:"required#请选择要寄存的酒品"`
}

/* ============ 员工工作台 ============ */

type StaffTableModel struct {
	TableId         int64   `json:"tableId"`
	TableName       string  `json:"tableName"`
	SeatStatus      int     `json:"seatStatus"`
	GuestCount      int     `json:"guestCount"`
	OpenAt          string  `json:"openAt"`
	SessionId       int64   `json:"sessionId"`
	PaidAmount      float64 `json:"paidAmount"`
	OrderCount      int     `json:"orderCount"`
	PendingPayCount int     `json:"pendingPayCount"`
	MakingCount     int     `json:"makingCount"`
	Called          bool    `json:"called"`
}

type StaffServeInp struct {
	OrderNo string  `json:"orderNo" v:"required#订单号不能为空"`
	ItemIds []int64 `json:"itemIds"`
}

type StaffUrgentInp struct {
	OrderNo string `json:"orderNo" v:"required#订单号不能为空"`
}

type StaffClearGoodsInp struct {
	GoodsId   int64 `json:"goodsId" v:"required#商品ID不能为空"`
	IsCleared bool  `json:"isCleared"`
}

type StaffVerifyBottleInp struct {
	Code      string  `json:"code" v:"required#取酒码不能为空"`
	DepositId int64   `json:"depositId"`
	TakeRatio float64 `json:"takeRatio"`
}

type StaffVerifyBottleModel struct {
	Deposit   *BottleModel `json:"deposit"`
	MemberId  int64        `json:"memberId"`
	TakeRatio float64      `json:"takeRatio"`
}
