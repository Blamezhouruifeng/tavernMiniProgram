package tavernin

import (
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

/* ============ 优惠券（C 端） ============ */

type CouponListInp struct {
	// Status 1未使用 2已使用 3已过期，0为全部
	Status int `json:"status"`
	// Amount 订单金额，用于筛选可用券
	Amount float64 `json:"amount"`
}

type CouponModel struct {
	Id             int64       `json:"id"`
	CouponId       int64       `json:"couponId"`
	Name           string      `json:"name"`
	Type           int         `json:"type"`
	Amount         float64     `json:"amount"`
	MinAmount      float64     `json:"minAmount"`
	Discount       float64     `json:"discount"`
	Code           string      `json:"code"`
	Source         int         `json:"source"`
	Status         int         `json:"status"`
	ExpireAt       *gtime.Time `json:"expireAt"`
	Usable         bool        `json:"usable"`
	UnusableReason string      `json:"unusableReason"`
}

/* ============ 优惠券（后台） ============ */

type AdminCouponListInp struct {
	form.PageReq
	Name string `json:"name"`
}

type AdminCouponModel struct {
	Id             int64       `json:"id"`
	Name           string      `json:"name"`
	Type           int         `json:"type"`
	Amount         float64     `json:"amount"`
	MinAmount      float64     `json:"minAmount"`
	Discount       float64     `json:"discount"`
	ValidDays      int         `json:"validDays"`
	TotalLimit     int         `json:"totalLimit"`
	PerMemberLimit int         `json:"perMemberLimit"`
	IsNewcomer     bool        `json:"isNewcomer"`
	Status         int         `json:"status"`
	Sort           int         `json:"sort"`
	Remark         string      `json:"remark"`
	GrantedCount   int         `json:"grantedCount"`
	UsedCount      int         `json:"usedCount"`
	CreatedAt      *gtime.Time `json:"createdAt"`
}

type AdminCouponEditInp struct {
	Id             int64   `json:"id"`
	Name           string  `json:"name" v:"required#券名称不能为空"`
	Type           int     `json:"type"`
	Amount         float64 `json:"amount"`
	MinAmount      float64 `json:"minAmount"`
	Discount       float64 `json:"discount"`
	ValidDays      int     `json:"validDays"`
	TotalLimit     int     `json:"totalLimit"`
	PerMemberLimit int     `json:"perMemberLimit"`
	IsNewcomer     bool    `json:"isNewcomer"`
	Status         int     `json:"status"`
	Sort           int     `json:"sort"`
	Remark         string  `json:"remark"`
}

type AdminCouponGrantInp struct {
	CouponId  int64   `json:"couponId" v:"required#请选择优惠券"`
	MemberIds []int64 `json:"memberIds"`
	Phone     string  `json:"phone"`
	All       bool    `json:"all"`
}
