// Package tavern 酒馆小程序端接口契约（前台 /api/tavern/**）
// 说明：路径与小程序端 uniapp/src/api/tavern/index.ts 一一对应，便于前后端按契约联调。
package tavern

import (
	"github.com/gogf/gf/v2/frame/g"

	"hotgo/internal/model/input/tavernin"
)

/* ---------------- 门店 ---------------- */

type StoreInfoReq struct {
	g.Meta `path:"/store/info" method:"get" tags:"酒馆-门店" summary:"门店信息"`
}

type StoreInfoRes struct {
	*tavernin.StoreInfoModel
}

/* ---------------- 会员（客人与员工共用同一套会员体系） ---------------- */

type MemberLoginReq struct {
	g.Meta `path:"/member/login" method:"post" tags:"酒馆-会员" summary:"会员登录（微信 code 换 token）"`
	tavernin.MemberLoginInp
}

type MemberLoginRes struct {
	*tavernin.MemberLoginModel
}

type MemberProfileReq struct {
	g.Meta `path:"/member/profile" method:"get" tags:"酒馆-会员" summary:"当前会员信息"`
}

type MemberProfileRes struct {
	*tavernin.MemberInfoModel
}

type MemberSwitchStaffReq struct {
	g.Meta `path:"/member/switchStaff" method:"post" tags:"酒馆-会员" summary:"切换员工身份（演示用）"`
	tavernin.MemberSwitchStaffInp
}

type MemberSwitchStaffRes struct {
	*tavernin.MemberInfoModel
}

/* ---------------- 菜单 ---------------- */

type MenuListReq struct {
	g.Meta `path:"/menu/list" method:"get" tags:"酒馆-菜单" summary:"菜单（分类+商品+规格）"`
}

type MenuListRes struct {
	List []*tavernin.MenuCategoryModel `json:"list"`
}

/* ---------------- 桌台会话 ---------------- */

type SessionOpenReq struct {
	g.Meta `path:"/session/open" method:"post" tags:"酒馆-桌台" summary:"扫码开台/进入桌台"`
	tavernin.SessionOpenInp
}

type SessionOpenRes struct {
	*tavernin.TableSessionModel
}

type SessionSnapshotReq struct {
	g.Meta    `path:"/session/snapshot" method:"get" tags:"酒馆-桌台" summary:"桌台快照（成员/购物车/订单/全桌已点）"`
	SessionId int64 `json:"sessionId" v:"required#会话ID不能为空"`
}

type SessionSnapshotRes struct {
	*tavernin.SessionSnapshotModel
}

type SessionCloseReq struct {
	g.Meta `path:"/session/close" method:"post" tags:"酒馆-桌台" summary:"清台"`
	tavernin.SessionCloseInp
}

type SessionCloseRes struct{}

/* ---------------- 购物车 ---------------- */

type CartAddReq struct {
	g.Meta `path:"/cart/add" method:"post" tags:"酒馆-购物车" summary:"加入购物车"`
	tavernin.CartAddInp
}

type CartAddRes struct{}

type CartUpdateReq struct {
	g.Meta `path:"/cart/update" method:"post" tags:"酒馆-购物车" summary:"修改购物车行"`
	tavernin.CartUpdateInp
}

type CartUpdateRes struct{}

type CartRemoveReq struct {
	g.Meta `path:"/cart/remove" method:"post" tags:"酒馆-购物车" summary:"删除购物车行"`
	tavernin.CartRemoveInp
}

type CartRemoveRes struct{}

/* ---------------- 订单（先付制） ---------------- */

type OrderSubmitReq struct {
	g.Meta `path:"/order/submit" method:"post" tags:"酒馆-订单" summary:"提交订单（待支付）"`
	tavernin.OrderSubmitInp
}

type OrderSubmitRes struct {
	*tavernin.OrderModel
}

type OrderPayReq struct {
	g.Meta `path:"/order/pay" method:"post" tags:"酒馆-订单" summary:"支付订单（当前为模拟支付）"`
	tavernin.OrderPayInp
}

type OrderPayRes struct {
	*tavernin.OrderModel
}

type OrderDetailReq struct {
	g.Meta `path:"/order/detail" method:"get" tags:"酒馆-订单" summary:"订单详情"`
	tavernin.OrderDetailInp
}

type OrderDetailRes struct {
	*tavernin.OrderModel
}

type OrderListReq struct {
	g.Meta `path:"/order/list" method:"get" tags:"酒馆-订单" summary:"订单列表（本桌/我下的单）"`
	tavernin.OrderListInp
}

type OrderListRes struct {
	List []*tavernin.OrderModel `json:"list"`
}

/* ---------------- 存酒 ---------------- */

type BottleListReq struct {
	g.Meta `path:"/bottle/list" method:"get" tags:"酒馆-存酒" summary:"我的存酒"`
	tavernin.BottleListInp
}

type BottleListRes struct {
	List []*tavernin.BottleModel `json:"list"`
}

type BottleTakeCodeReq struct {
	g.Meta `path:"/bottle/takeCode" method:"post" tags:"酒馆-存酒" summary:"生成动态取酒码"`
	tavernin.TakeCodeInp
}

type BottleTakeCodeRes struct {
	*tavernin.TakeCodeModel
}

type BottleDepositFromOrderReq struct {
	g.Meta `path:"/bottle/depositFromOrder" method:"post" tags:"酒馆-存酒" summary:"退酒入存"`
	tavernin.DepositFromOrderInp
}

type BottleDepositFromOrderRes struct {
	*tavernin.BottleModel
}

/* ---------------- 员工工作台 ---------------- */

type StaffTablesReq struct {
	g.Meta `path:"/staff/tables" method:"get" tags:"酒馆-员工" summary:"桌台看板"`
}

type StaffTablesRes struct {
	List []*tavernin.StaffTableModel `json:"list"`
}

type StaffTicketsReq struct {
	g.Meta `path:"/staff/tickets" method:"get" tags:"酒馆-员工" summary:"出票区（已支付订单）"`
}

type StaffTicketsRes struct {
	List []*tavernin.OrderModel `json:"list"`
}

type StaffServeReq struct {
	g.Meta `path:"/staff/serve" method:"post" tags:"酒馆-员工" summary:"行级/整单上齐"`
	tavernin.StaffServeInp
}

type StaffServeRes struct {
	*tavernin.OrderModel
}

type StaffUrgentReq struct {
	g.Meta `path:"/staff/urgent" method:"post" tags:"酒馆-员工" summary:"订单加急"`
	tavernin.StaffUrgentInp
}

type StaffUrgentRes struct{}

type StaffClearGoodsReq struct {
	g.Meta `path:"/staff/goods/clear" method:"post" tags:"酒馆-员工" summary:"沽清/恢复供应"`
	tavernin.StaffClearGoodsInp
}

type StaffClearGoodsRes struct{}

type StaffBottleVerifyReq struct {
	g.Meta `path:"/staff/bottle/verify" method:"post" tags:"酒馆-员工" summary:"存酒核销"`
	tavernin.StaffVerifyBottleInp
}

type StaffBottleVerifyRes struct {
	*tavernin.StaffVerifyBottleModel
}
