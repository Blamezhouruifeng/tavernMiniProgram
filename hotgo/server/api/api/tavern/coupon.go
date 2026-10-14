package tavern

import (
	"github.com/gogf/gf/v2/frame/g"

	"hotgo/internal/model/input/tavernin"
)

/* ---------------- 优惠券（C 端） ---------------- */

type CouponListReq struct {
	g.Meta `path:"/coupon/list" method:"get" tags:"酒馆-优惠券" summary:"我的优惠券（可按金额筛选可用）"`
	tavernin.CouponListInp
}

type CouponListRes struct {
	List []*tavernin.CouponModel `json:"list"`
}
