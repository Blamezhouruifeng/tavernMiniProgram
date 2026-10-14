// Package tavern 酒馆后台管理接口契约（/admin/tavern/**）
package tavern

import (
	"github.com/gogf/gf/v2/frame/g"

	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/tavernin"
)

/* ---------------- 门店配置 ---------------- */

type StoreViewReq struct {
	g.Meta `path:"/tavern/store/view" method:"get" tags:"酒馆-门店" summary:"获取门店配置"`
}

type StoreViewRes struct {
	*tavernin.AdminStoreViewModel
}

type StoreEditReq struct {
	g.Meta `path:"/tavern/store/edit" method:"post" tags:"酒馆-门店" summary:"保存门店配置"`
	tavernin.AdminStoreEditInp
}

type StoreEditRes struct{}

/* ---------------- 区域 ---------------- */

type AreaListReq struct {
	g.Meta `path:"/tavern/area/list" method:"get" tags:"酒馆-区域" summary:"区域列表"`
	tavernin.AdminAreaListInp
}

type AreaListRes struct {
	form.PageRes
	List []*tavernin.AdminAreaModel `json:"list"`
}

type AreaEditReq struct {
	g.Meta `path:"/tavern/area/edit" method:"post" tags:"酒馆-区域" summary:"保存区域"`
	tavernin.AdminAreaEditInp
}

type AreaEditRes struct{}

type AreaDeleteReq struct {
	g.Meta `path:"/tavern/area/delete" method:"post" tags:"酒馆-区域" summary:"删除区域"`
	tavernin.AdminDeleteInp
}

type AreaDeleteRes struct{}

/* ---------------- 桌位 ---------------- */

type TableListReq struct {
	g.Meta `path:"/tavern/table/list" method:"get" tags:"酒馆-桌位" summary:"桌位列表"`
	tavernin.AdminTableListInp
}

type TableListRes struct {
	form.PageRes
	List []*tavernin.AdminTableModel `json:"list"`
}

type TableEditReq struct {
	g.Meta `path:"/tavern/table/edit" method:"post" tags:"酒馆-桌位" summary:"保存桌位"`
	tavernin.AdminTableEditInp
}

type TableEditRes struct{}

type TableDeleteReq struct {
	g.Meta `path:"/tavern/table/delete" method:"post" tags:"酒馆-桌位" summary:"删除桌位"`
	tavernin.AdminDeleteInp
}

type TableDeleteRes struct{}

/* ---------------- 分类 ---------------- */

type CategoryListReq struct {
	g.Meta `path:"/tavern/category/list" method:"get" tags:"酒馆-分类" summary:"分类列表"`
	tavernin.AdminCategoryListInp
}

type CategoryListRes struct {
	form.PageRes
	List []*tavernin.AdminCategoryModel `json:"list"`
}

type CategoryEditReq struct {
	g.Meta `path:"/tavern/category/edit" method:"post" tags:"酒馆-分类" summary:"保存分类"`
	tavernin.AdminCategoryEditInp
}

type CategoryEditRes struct{}

type CategoryDeleteReq struct {
	g.Meta `path:"/tavern/category/delete" method:"post" tags:"酒馆-分类" summary:"删除分类"`
	tavernin.AdminDeleteInp
}

type CategoryDeleteRes struct{}

/* ---------------- 商品 ---------------- */

type GoodsListReq struct {
	g.Meta `path:"/tavern/goods/list" method:"get" tags:"酒馆-商品" summary:"商品列表"`
	tavernin.AdminGoodsListInp
}

type GoodsListRes struct {
	form.PageRes
	List []*tavernin.AdminGoodsModel `json:"list"`
}

type GoodsViewReq struct {
	g.Meta `path:"/tavern/goods/view" method:"get" tags:"酒馆-商品" summary:"商品详情（含规格）"`
	tavernin.AdminGoodsViewInp
}

type GoodsViewRes struct {
	*tavernin.AdminGoodsModel
}

type GoodsEditReq struct {
	g.Meta `path:"/tavern/goods/edit" method:"post" tags:"酒馆-商品" summary:"保存商品（含规格）"`
	tavernin.AdminGoodsEditInp
}

type GoodsEditRes struct{}

type GoodsDeleteReq struct {
	g.Meta `path:"/tavern/goods/delete" method:"post" tags:"酒馆-商品" summary:"删除商品"`
	tavernin.AdminDeleteInp
}

type GoodsDeleteRes struct{}

type GoodsStatusReq struct {
	g.Meta `path:"/tavern/goods/status" method:"post" tags:"酒馆-商品" summary:"商品上下架"`
	tavernin.AdminGoodsStatusInp
}

type GoodsStatusRes struct{}

/* ---------------- 订单查询 ---------------- */

type OrderListReq struct {
	g.Meta `path:"/tavern/order/list" method:"get" tags:"酒馆-订单" summary:"订单查询"`
	tavernin.AdminOrderListInp
}

type OrderListRes struct {
	form.PageRes
	List []*tavernin.AdminOrderModel `json:"list"`
}

/* ---------------- 存酒台账 ---------------- */

type BottleListReq struct {
	g.Meta `path:"/tavern/bottle/list" method:"get" tags:"酒馆-存酒" summary:"存酒台账"`
	tavernin.AdminBottleListInp
}

type BottleListRes struct {
	form.PageRes
	List []*tavernin.BottleModel `json:"list"`
}

type BottleUpdateReq struct {
	g.Meta `path:"/tavern/bottle/update" method:"post" tags:"酒馆-存酒" summary:"存酒维护（延期/报损/改储位）"`
	tavernin.AdminBottleUpdateInp
}

type BottleUpdateRes struct{}
