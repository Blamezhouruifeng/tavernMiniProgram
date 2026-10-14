package sys

import (
	"context"

	adminTavern "hotgo/api/admin/tavern"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

var (
	TavernAdmin = cTavernAdmin{}
)

type cTavernAdmin struct{}

/* ---------------- 门店 ---------------- */

func (c *cTavernAdmin) StoreView(ctx context.Context, req *adminTavern.StoreViewReq) (res *adminTavern.StoreViewRes, err error) {
	view, err := service.TavernAdmin().StoreView(ctx)
	if err != nil {
		return
	}
	res = &adminTavern.StoreViewRes{AdminStoreViewModel: view}
	return
}

func (c *cTavernAdmin) StoreEdit(ctx context.Context, req *adminTavern.StoreEditReq) (res *adminTavern.StoreEditRes, err error) {
	err = service.TavernAdmin().StoreEdit(ctx, &req.AdminStoreEditInp)
	return
}

/* ---------------- 区域 ---------------- */

func (c *cTavernAdmin) AreaList(ctx context.Context, req *adminTavern.AreaListReq) (res *adminTavern.AreaListRes, err error) {
	list, totalCount, err := service.TavernAdmin().AreaList(ctx, &req.AdminAreaListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*tavernin.AdminAreaModel{}
	}
	res = new(adminTavern.AreaListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *cTavernAdmin) AreaEdit(ctx context.Context, req *adminTavern.AreaEditReq) (res *adminTavern.AreaEditRes, err error) {
	err = service.TavernAdmin().AreaEdit(ctx, &req.AdminAreaEditInp)
	return
}

func (c *cTavernAdmin) AreaDelete(ctx context.Context, req *adminTavern.AreaDeleteReq) (res *adminTavern.AreaDeleteRes, err error) {
	err = service.TavernAdmin().AreaDelete(ctx, &req.AdminDeleteInp)
	return
}

/* ---------------- 桌位 ---------------- */

func (c *cTavernAdmin) TableList(ctx context.Context, req *adminTavern.TableListReq) (res *adminTavern.TableListRes, err error) {
	list, totalCount, err := service.TavernAdmin().TableList(ctx, &req.AdminTableListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*tavernin.AdminTableModel{}
	}
	res = new(adminTavern.TableListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *cTavernAdmin) TableEdit(ctx context.Context, req *adminTavern.TableEditReq) (res *adminTavern.TableEditRes, err error) {
	err = service.TavernAdmin().TableEdit(ctx, &req.AdminTableEditInp)
	return
}

func (c *cTavernAdmin) TableDelete(ctx context.Context, req *adminTavern.TableDeleteReq) (res *adminTavern.TableDeleteRes, err error) {
	err = service.TavernAdmin().TableDelete(ctx, &req.AdminDeleteInp)
	return
}

/* ---------------- 分类 ---------------- */

func (c *cTavernAdmin) CategoryList(ctx context.Context, req *adminTavern.CategoryListReq) (res *adminTavern.CategoryListRes, err error) {
	list, totalCount, err := service.TavernAdmin().CategoryList(ctx, &req.AdminCategoryListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*tavernin.AdminCategoryModel{}
	}
	res = new(adminTavern.CategoryListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *cTavernAdmin) CategoryEdit(ctx context.Context, req *adminTavern.CategoryEditReq) (res *adminTavern.CategoryEditRes, err error) {
	err = service.TavernAdmin().CategoryEdit(ctx, &req.AdminCategoryEditInp)
	return
}

func (c *cTavernAdmin) CategoryDelete(ctx context.Context, req *adminTavern.CategoryDeleteReq) (res *adminTavern.CategoryDeleteRes, err error) {
	err = service.TavernAdmin().CategoryDelete(ctx, &req.AdminDeleteInp)
	return
}

/* ---------------- 商品 ---------------- */

func (c *cTavernAdmin) GoodsList(ctx context.Context, req *adminTavern.GoodsListReq) (res *adminTavern.GoodsListRes, err error) {
	list, totalCount, err := service.TavernAdmin().GoodsList(ctx, &req.AdminGoodsListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*tavernin.AdminGoodsModel{}
	}
	res = new(adminTavern.GoodsListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *cTavernAdmin) GoodsView(ctx context.Context, req *adminTavern.GoodsViewReq) (res *adminTavern.GoodsViewRes, err error) {
	view, err := service.TavernAdmin().GoodsView(ctx, &req.AdminGoodsViewInp)
	if err != nil {
		return
	}
	res = &adminTavern.GoodsViewRes{AdminGoodsModel: view}
	return
}

func (c *cTavernAdmin) GoodsEdit(ctx context.Context, req *adminTavern.GoodsEditReq) (res *adminTavern.GoodsEditRes, err error) {
	err = service.TavernAdmin().GoodsEdit(ctx, &req.AdminGoodsEditInp)
	return
}

func (c *cTavernAdmin) GoodsDelete(ctx context.Context, req *adminTavern.GoodsDeleteReq) (res *adminTavern.GoodsDeleteRes, err error) {
	err = service.TavernAdmin().GoodsDelete(ctx, &req.AdminDeleteInp)
	return
}

func (c *cTavernAdmin) GoodsStatus(ctx context.Context, req *adminTavern.GoodsStatusReq) (res *adminTavern.GoodsStatusRes, err error) {
	err = service.TavernAdmin().GoodsStatus(ctx, &req.AdminGoodsStatusInp)
	return
}

/* ---------------- 订单 ---------------- */

func (c *cTavernAdmin) OrderList(ctx context.Context, req *adminTavern.OrderListReq) (res *adminTavern.OrderListRes, err error) {
	list, totalCount, err := service.TavernAdmin().OrderList(ctx, &req.AdminOrderListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*tavernin.AdminOrderModel{}
	}
	res = new(adminTavern.OrderListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

/* ---------------- 存酒 ---------------- */

func (c *cTavernAdmin) BottleList(ctx context.Context, req *adminTavern.BottleListReq) (res *adminTavern.BottleListRes, err error) {
	list, totalCount, err := service.TavernAdmin().BottleList(ctx, &req.AdminBottleListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*tavernin.BottleModel{}
	}
	res = new(adminTavern.BottleListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *cTavernAdmin) BottleUpdate(ctx context.Context, req *adminTavern.BottleUpdateReq) (res *adminTavern.BottleUpdateRes, err error) {
	err = service.TavernAdmin().BottleUpdate(ctx, &req.AdminBottleUpdateInp)
	return
}

/* ---------------- 异常订单与对账 ---------------- */

func (c *cTavernAdmin) OrderAbnormal(ctx context.Context, req *adminTavern.OrderAbnormalReq) (res *adminTavern.OrderAbnormalRes, err error) {
	list, err := service.TavernAdmin().AbnormalOrders(ctx)
	if err != nil {
		return
	}
	if list == nil {
		list = []*tavernin.AdminAbnormalOrderModel{}
	}
	res = &adminTavern.OrderAbnormalRes{List: list}
	return
}

func (c *cTavernAdmin) OrderReissue(ctx context.Context, req *adminTavern.OrderReissueReq) (res *adminTavern.OrderReissueRes, err error) {
	err = service.TavernAdmin().Reissue(ctx, &req.AdminReissueInp)
	return
}

func (c *cTavernAdmin) Reconcile(ctx context.Context, req *adminTavern.ReconcileReq) (res *adminTavern.ReconcileRes, err error) {
	view, err := service.TavernAdmin().Reconcile(ctx, &req.AdminReconcileInp)
	if err != nil {
		return
	}
	res = &adminTavern.ReconcileRes{AdminReconcileModel: view}
	return
}
