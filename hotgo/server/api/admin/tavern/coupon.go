package tavern

import (
	"github.com/gogf/gf/v2/frame/g"

	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/tavernin"
)

/* ---------------- 优惠券（后台） ---------------- */

type CouponListReq struct {
	g.Meta `path:"/tavern/coupon/list" method:"get" tags:"酒馆-优惠券" summary:"优惠券模板列表"`
	tavernin.AdminCouponListInp
}

type CouponListRes struct {
	form.PageRes
	List []*tavernin.AdminCouponModel `json:"list"`
}

type CouponEditReq struct {
	g.Meta `path:"/tavern/coupon/edit" method:"post" tags:"酒馆-优惠券" summary:"保存优惠券模板"`
	tavernin.AdminCouponEditInp
}

type CouponEditRes struct{}

type CouponGrantReq struct {
	g.Meta `path:"/tavern/coupon/grant" method:"post" tags:"酒馆-优惠券" summary:"发放优惠券（定向/全体）"`
	tavernin.AdminCouponGrantInp
}

type CouponGrantRes struct {
	Count int `json:"count"`
}
