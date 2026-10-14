// Package tavern 酒馆小程序端控制器（薄控制器：只做参数透传与结果组装）
package tavern

import (
	"context"

	apiTavern "hotgo/api/api/tavern"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

/* ---------------- 门店 ---------------- */

type cStore struct{}

var Store = cStore{}

func (c *cStore) StoreInfo(ctx context.Context, req *apiTavern.StoreInfoReq) (res *apiTavern.StoreInfoRes, err error) {
	info, err := service.TavernStore().Info(ctx)
	if err != nil {
		return
	}
	res = &apiTavern.StoreInfoRes{StoreInfoModel: info}
	return
}

/* ---------------- 会员 ---------------- */

type cMember struct{}

var Member = cMember{}

func (c *cMember) MemberLogin(ctx context.Context, req *apiTavern.MemberLoginReq) (res *apiTavern.MemberLoginRes, err error) {
	login, err := service.TavernMember().Login(ctx, &req.MemberLoginInp)
	if err != nil {
		return
	}
	res = &apiTavern.MemberLoginRes{MemberLoginModel: login}
	return
}

func (c *cMember) MemberProfile(ctx context.Context, req *apiTavern.MemberProfileReq) (res *apiTavern.MemberProfileRes, err error) {
	info, err := service.TavernMember().Profile(ctx)
	if err != nil {
		return
	}
	res = &apiTavern.MemberProfileRes{MemberInfoModel: info}
	return
}

func (c *cMember) MemberSwitchStaff(ctx context.Context, req *apiTavern.MemberSwitchStaffReq) (res *apiTavern.MemberSwitchStaffRes, err error) {
	info, err := service.TavernMember().SwitchStaff(ctx, &req.MemberSwitchStaffInp)
	if err != nil {
		return
	}
	res = &apiTavern.MemberSwitchStaffRes{MemberInfoModel: info}
	return
}

/* ---------------- 菜单 ---------------- */

type cMenu struct{}

var Menu = cMenu{}

func (c *cMenu) MenuList(ctx context.Context, req *apiTavern.MenuListReq) (res *apiTavern.MenuListRes, err error) {
	list, err := service.TavernMenu().List(ctx)
	if err != nil {
		return
	}
	if list == nil {
		list = make([]*tavernin.MenuCategoryModel, 0)
	}
	res = &apiTavern.MenuListRes{List: list}
	return
}

/* ---------------- 桌台会话 ---------------- */

type cSession struct{}

var Session = cSession{}

func (c *cSession) SessionOpen(ctx context.Context, req *apiTavern.SessionOpenReq) (res *apiTavern.SessionOpenRes, err error) {
	session, err := service.TavernSession().Open(ctx, &req.SessionOpenInp)
	if err != nil {
		return
	}
	res = &apiTavern.SessionOpenRes{TableSessionModel: session}
	return
}

func (c *cSession) SessionSnapshot(ctx context.Context, req *apiTavern.SessionSnapshotReq) (res *apiTavern.SessionSnapshotRes, err error) {
	snapshot, err := service.TavernSession().Snapshot(ctx, req.SessionId)
	if err != nil {
		return
	}
	res = &apiTavern.SessionSnapshotRes{SessionSnapshotModel: snapshot}
	return
}

func (c *cSession) SessionClose(ctx context.Context, req *apiTavern.SessionCloseReq) (res *apiTavern.SessionCloseRes, err error) {
	err = service.TavernSession().Close(ctx, &req.SessionCloseInp)
	return
}

/* ---------------- 购物车 ---------------- */

type cCart struct{}

var Cart = cCart{}

func (c *cCart) CartAdd(ctx context.Context, req *apiTavern.CartAddReq) (res *apiTavern.CartAddRes, err error) {
	err = service.TavernCart().Add(ctx, &req.CartAddInp)
	return
}

func (c *cCart) CartUpdate(ctx context.Context, req *apiTavern.CartUpdateReq) (res *apiTavern.CartUpdateRes, err error) {
	err = service.TavernCart().Update(ctx, &req.CartUpdateInp)
	return
}

func (c *cCart) CartRemove(ctx context.Context, req *apiTavern.CartRemoveReq) (res *apiTavern.CartRemoveRes, err error) {
	err = service.TavernCart().Remove(ctx, &req.CartRemoveInp)
	return
}

/* ---------------- 订单 ---------------- */

type cOrder struct{}

var Order = cOrder{}

func (c *cOrder) OrderSubmit(ctx context.Context, req *apiTavern.OrderSubmitReq) (res *apiTavern.OrderSubmitRes, err error) {
	order, err := service.TavernOrder().Submit(ctx, &req.OrderSubmitInp)
	if err != nil {
		return
	}
	res = &apiTavern.OrderSubmitRes{OrderModel: order}
	return
}

func (c *cOrder) OrderPay(ctx context.Context, req *apiTavern.OrderPayReq) (res *apiTavern.OrderPayRes, err error) {
	order, err := service.TavernOrder().Pay(ctx, &req.OrderPayInp)
	if err != nil {
		return
	}
	res = &apiTavern.OrderPayRes{OrderModel: order}
	return
}

func (c *cOrder) OrderDetail(ctx context.Context, req *apiTavern.OrderDetailReq) (res *apiTavern.OrderDetailRes, err error) {
	order, err := service.TavernOrder().Detail(ctx, &req.OrderDetailInp)
	if err != nil {
		return
	}
	res = &apiTavern.OrderDetailRes{OrderModel: order}
	return
}

func (c *cOrder) OrderList(ctx context.Context, req *apiTavern.OrderListReq) (res *apiTavern.OrderListRes, err error) {
	list, err := service.TavernOrder().List(ctx, &req.OrderListInp)
	if err != nil {
		return
	}
	res = &apiTavern.OrderListRes{List: list}
	return
}

/* ---------------- 存酒 ---------------- */

type cBottle struct{}

var Bottle = cBottle{}

func (c *cBottle) BottleList(ctx context.Context, req *apiTavern.BottleListReq) (res *apiTavern.BottleListRes, err error) {
	list, err := service.TavernDeposit().List(ctx, &req.BottleListInp)
	if err != nil {
		return
	}
	res = &apiTavern.BottleListRes{List: list}
	return
}

func (c *cBottle) BottleTakeCode(ctx context.Context, req *apiTavern.BottleTakeCodeReq) (res *apiTavern.BottleTakeCodeRes, err error) {
	code, err := service.TavernDeposit().TakeCode(ctx, &req.TakeCodeInp)
	if err != nil {
		return
	}
	res = &apiTavern.BottleTakeCodeRes{TakeCodeModel: code}
	return
}

func (c *cBottle) BottleDepositFromOrder(ctx context.Context, req *apiTavern.BottleDepositFromOrderReq) (res *apiTavern.BottleDepositFromOrderRes, err error) {
	bottle, err := service.TavernDeposit().DepositFromOrder(ctx, &req.DepositFromOrderInp)
	if err != nil {
		return
	}
	res = &apiTavern.BottleDepositFromOrderRes{BottleModel: bottle}
	return
}

/* ---------------- 员工工作台 ---------------- */

type cStaff struct{}

var Staff = cStaff{}

func (c *cStaff) StaffTables(ctx context.Context, req *apiTavern.StaffTablesReq) (res *apiTavern.StaffTablesRes, err error) {
	list, err := service.TavernStaff().Tables(ctx)
	if err != nil {
		return
	}
	res = &apiTavern.StaffTablesRes{List: list}
	return
}

func (c *cStaff) StaffTickets(ctx context.Context, req *apiTavern.StaffTicketsReq) (res *apiTavern.StaffTicketsRes, err error) {
	list, err := service.TavernStaff().Tickets(ctx)
	if err != nil {
		return
	}
	res = &apiTavern.StaffTicketsRes{List: list}
	return
}

func (c *cStaff) StaffServe(ctx context.Context, req *apiTavern.StaffServeReq) (res *apiTavern.StaffServeRes, err error) {
	order, err := service.TavernStaff().Serve(ctx, &req.StaffServeInp)
	if err != nil {
		return
	}
	res = &apiTavern.StaffServeRes{OrderModel: order}
	return
}

func (c *cStaff) StaffUrgent(ctx context.Context, req *apiTavern.StaffUrgentReq) (res *apiTavern.StaffUrgentRes, err error) {
	err = service.TavernStaff().Urgent(ctx, &req.StaffUrgentInp)
	return
}

func (c *cStaff) StaffClearGoods(ctx context.Context, req *apiTavern.StaffClearGoodsReq) (res *apiTavern.StaffClearGoodsRes, err error) {
	err = service.TavernStaff().ClearGoods(ctx, &req.StaffClearGoodsInp)
	return
}

func (c *cStaff) StaffBottleVerify(ctx context.Context, req *apiTavern.StaffBottleVerifyReq) (res *apiTavern.StaffBottleVerifyRes, err error) {
	verify, err := service.TavernStaff().VerifyBottle(ctx, &req.StaffVerifyBottleInp)
	if err != nil {
		return
	}
	res = &apiTavern.StaffBottleVerifyRes{StaffVerifyBottleModel: verify}
	return
}

func (c *cBottle) CouponList(ctx context.Context, req *apiTavern.CouponListReq) (res *apiTavern.CouponListRes, err error) {
	list, err := service.TavernCoupon().MyList(ctx, &req.CouponListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*tavernin.CouponModel{}
	}
	res = &apiTavern.CouponListRes{List: list}
	return
}
