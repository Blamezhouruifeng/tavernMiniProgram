// Package service 酒馆小程序业务服务接口声明
// 说明：与 HotGo 主模块同目录拆分（一个业务域一个文件），logic 实现位于 internal/logic/tavern
package service

import (
	"context"

	"hotgo/internal/model/input/tavernin"
)

type (
	// ITavernStore 门店
	ITavernStore interface {
		// Info 获取门店信息（含先付制、打烊截止配置）
		Info(ctx context.Context) (res *tavernin.StoreInfoModel, err error)
	}

	// ITavernMember 会员（客人与员工共用同一套会员体系）
	ITavernMember interface {
		// Login 微信登录（建档 / 绑定）
		Login(ctx context.Context, in *tavernin.MemberLoginInp) (res *tavernin.MemberLoginModel, err error)
		// Profile 当前会员信息
		Profile(ctx context.Context) (res *tavernin.MemberInfoModel, err error)
		// SwitchStaff 演示用：切换员工身份（真实环境由后台标记）
		SwitchStaff(ctx context.Context, in *tavernin.MemberSwitchStaffInp) (res *tavernin.MemberInfoModel, err error)
		// BindPhone 绑定手机号（存酒取酒凭证的前置条件）
		BindPhone(ctx context.Context, in *tavernin.BindPhoneInp) (res *tavernin.MemberInfoModel, err error)
		// CurrentId 当前登录会员ID
		CurrentId(ctx context.Context) int64
	}

	// ITavernMenu 菜单
	ITavernMenu interface {
		// List 菜单（分类 + 商品 + 规格，含分时段过滤）
		List(ctx context.Context) (res []*tavernin.MenuCategoryModel, err error)
	}

	// ITavernSession 桌台会话
	ITavernSession interface {
		// Open 扫码开台（二次扫码回到原会话）
		Open(ctx context.Context, in *tavernin.SessionOpenInp) (res *tavernin.TableSessionModel, err error)
		// Snapshot 会话快照（成员 / 购物车 / 订单 / 全桌已点聚合）
		Snapshot(ctx context.Context, sessionId int64) (res *tavernin.SessionSnapshotModel, err error)
		// Close 清台
		Close(ctx context.Context, in *tavernin.SessionCloseInp) (err error)
		// CurrentSessionId 当前桌台会话ID
		CurrentSessionId(ctx context.Context, sessionId int64) (id int64, err error)
		// Call 呼叫服务员 / 催单
		Call(ctx context.Context, in *tavernin.SessionCallInp) (err error)
	}

	// ITavernCart 购物车（每人一个篮子，全桌可见）
	ITavernCart interface {
		// Add 加入购物车
		Add(ctx context.Context, in *tavernin.CartAddInp) (err error)
		// Update 修改自己篮子里的行
		Update(ctx context.Context, in *tavernin.CartUpdateInp) (err error)
		// Remove 删除自己篮子里的行
		Remove(ctx context.Context, in *tavernin.CartRemoveInp) (err error)
	}

	// ITavernOrder 订单（先付制：支付成功 = 进入制作）
	ITavernOrder interface {
		// Submit 提交订单（返回待支付订单）
		Submit(ctx context.Context, in *tavernin.OrderSubmitInp) (res *tavernin.OrderModel, err error)
		// Pay 支付（当前为模拟支付，预留微信支付回调与幂等）
		Pay(ctx context.Context, in *tavernin.OrderPayInp) (res *tavernin.OrderModel, err error)
		// Detail 订单详情
		Detail(ctx context.Context, in *tavernin.OrderDetailInp) (res *tavernin.OrderModel, err error)
		// List 订单列表（本桌 / 我下的单）
		List(ctx context.Context, in *tavernin.OrderListInp) (res []*tavernin.OrderModel, err error)
		// ApplyRefund 申请退款（制作前可自助通过，制作后需员工审核）
		ApplyRefund(ctx context.Context, in *tavernin.RefundApplyInp) (res *tavernin.OrderModel, err error)
		// AuditRefund 退款审核（员工）
		AuditRefund(ctx context.Context, in *tavernin.RefundAuditInp) (res *tavernin.OrderModel, err error)
	}

	// ITavernDeposit 存酒
	ITavernDeposit interface {
		// List 我的存酒
		List(ctx context.Context, in *tavernin.BottleListInp) (res []*tavernin.BottleModel, err error)
		// TakeCode 生成动态取酒码（30 秒有效）
		TakeCode(ctx context.Context, in *tavernin.TakeCodeInp) (res *tavernin.TakeCodeModel, err error)
		// DepositFromOrder 退酒入存
		DepositFromOrder(ctx context.Context, in *tavernin.DepositFromOrderInp) (res *tavernin.BottleModel, err error)
	}

	// ITavernCoupon 优惠券（新人礼 / 发放 / 试算 / 核销）
	ITavernCoupon interface {
		// MyList 我的优惠券
		MyList(ctx context.Context, in *tavernin.CouponListInp) (res []*tavernin.CouponModel, err error)
		// GrantNewcomer 新人礼自动发放
		GrantNewcomer(ctx context.Context, memberId int64) (err error)
		// Grant 发放优惠券
		Grant(ctx context.Context, couponId, memberId int64, source int) (res *tavernin.CouponModel, err error)
		// CalcDiscount 试算优惠金额
		CalcDiscount(ctx context.Context, memberId, couponMemberId int64, amount float64) (discount float64, err error)
		// UseCoupon 标记已使用（支付成功）
		UseCoupon(ctx context.Context, memberId, couponMemberId, orderId int64, discount float64) (err error)
		// ReturnCoupon 退款时退回
		ReturnCoupon(ctx context.Context, orderId int64) (err error)
	}

	// ITavernAdmin 后台管理（HotGo 后台）
	ITavernAdmin interface {
		// StoreView 门店配置详情
		StoreView(ctx context.Context) (res *tavernin.AdminStoreViewModel, err error)
		// StoreEdit 保存门店配置
		StoreEdit(ctx context.Context, in *tavernin.AdminStoreEditInp) (err error)

		// AreaList 区域列表
		AreaList(ctx context.Context, in *tavernin.AdminAreaListInp) (list []*tavernin.AdminAreaModel, totalCount int, err error)
		// AreaEdit 保存区域
		AreaEdit(ctx context.Context, in *tavernin.AdminAreaEditInp) (err error)
		// AreaDelete 删除区域
		AreaDelete(ctx context.Context, in *tavernin.AdminDeleteInp) (err error)

		// TableList 桌位列表
		TableList(ctx context.Context, in *tavernin.AdminTableListInp) (list []*tavernin.AdminTableModel, totalCount int, err error)
		// TableEdit 保存桌位（自动生成桌码 scene）
		TableEdit(ctx context.Context, in *tavernin.AdminTableEditInp) (err error)
		// TableDelete 删除桌位
		TableDelete(ctx context.Context, in *tavernin.AdminDeleteInp) (err error)

		// CategoryList 分类列表
		CategoryList(ctx context.Context, in *tavernin.AdminCategoryListInp) (list []*tavernin.AdminCategoryModel, totalCount int, err error)
		// CategoryEdit 保存分类
		CategoryEdit(ctx context.Context, in *tavernin.AdminCategoryEditInp) (err error)
		// CategoryDelete 删除分类
		CategoryDelete(ctx context.Context, in *tavernin.AdminDeleteInp) (err error)

		// GoodsList 商品列表
		GoodsList(ctx context.Context, in *tavernin.AdminGoodsListInp) (list []*tavernin.AdminGoodsModel, totalCount int, err error)
		// GoodsView 商品详情（含规格）
		GoodsView(ctx context.Context, in *tavernin.AdminGoodsViewInp) (res *tavernin.AdminGoodsModel, err error)
		// GoodsEdit 保存商品（含规格子表）
		GoodsEdit(ctx context.Context, in *tavernin.AdminGoodsEditInp) (err error)
		// GoodsDelete 删除商品
		GoodsDelete(ctx context.Context, in *tavernin.AdminDeleteInp) (err error)
		// GoodsStatus 上下架
		GoodsStatus(ctx context.Context, in *tavernin.AdminGoodsStatusInp) (err error)

		// OrderList 订单查询
		OrderList(ctx context.Context, in *tavernin.AdminOrderListInp) (list []*tavernin.AdminOrderModel, totalCount int, err error)

		// BottleList 存酒台账
		BottleList(ctx context.Context, in *tavernin.AdminBottleListInp) (list []*tavernin.BottleModel, totalCount int, err error)
		// BottleUpdate 存酒维护（延期/报损/改储位）
		BottleUpdate(ctx context.Context, in *tavernin.AdminBottleUpdateInp) (err error)

		// AbnormalOrders 异常订单（支付成功未出票 / 退款待审 / 超时未取消）
		AbnormalOrders(ctx context.Context) (list []*tavernin.AdminAbnormalOrderModel, err error)
		// Reissue 手动补出票（先付制：客人已付款但未进入制作时使用）
		Reissue(ctx context.Context, in *tavernin.AdminReissueInp) (err error)
		// Reconcile 支付对账（订单 ↔ 支付流水）
		Reconcile(ctx context.Context, in *tavernin.AdminReconcileInp) (res *tavernin.AdminReconcileModel, err error)

		// CouponList 优惠券模板列表
		CouponList(ctx context.Context, in *tavernin.AdminCouponListInp) (list []*tavernin.AdminCouponModel, totalCount int, err error)
		// CouponEdit 保存优惠券模板
		CouponEdit(ctx context.Context, in *tavernin.AdminCouponEditInp) (err error)
		// CouponGrant 发放优惠券（定向/全员）
		CouponGrant(ctx context.Context, in *tavernin.AdminCouponGrantInp) (count int, err error)
	}

	// ITavernStaff 员工工作台
	ITavernStaff interface {
		// Tables 桌台看板
		Tables(ctx context.Context) (res []*tavernin.StaffTableModel, err error)
		// Tickets 出票区（已支付订单）
		Tickets(ctx context.Context) (res []*tavernin.OrderModel, err error)
		// Serve 行级 / 整单上齐
		Serve(ctx context.Context, in *tavernin.StaffServeInp) (res *tavernin.OrderModel, err error)
		// Urgent 加急
		Urgent(ctx context.Context, in *tavernin.StaffUrgentInp) (err error)
		// ClearGoods 沽清
		ClearGoods(ctx context.Context, in *tavernin.StaffClearGoodsInp) (err error)
		// VerifyBottle 存酒核销
		VerifyBottle(ctx context.Context, in *tavernin.StaffVerifyBottleInp) (res *tavernin.StaffVerifyBottleModel, err error)
		// DepositRegister 代存登记（自带酒水，实物入库 + 打印瓶签）
		DepositRegister(ctx context.Context, in *tavernin.DepositRegisterInp) (res *tavernin.BottleModel, err error)
		// MoveTable 转台/并台（订单与存酒归属随会话迁移）
		MoveTable(ctx context.Context, in *tavernin.MoveTableInp) (err error)
		// HandleCall 响应呼叫（员工确认）
		HandleCall(ctx context.Context, in *tavernin.SessionCallInp) (err error)
	}
)

var (
	localTavernStore   ITavernStore
	localTavernMember  ITavernMember
	localTavernMenu    ITavernMenu
	localTavernSession ITavernSession
	localTavernCart    ITavernCart
	localTavernOrder   ITavernOrder
	localTavernDeposit ITavernDeposit
	localTavernStaff   ITavernStaff
	localTavernAdmin   ITavernAdmin
	localTavernCoupon  ITavernCoupon
)

func TavernCoupon() ITavernCoupon {
	if localTavernCoupon == nil {
		panic("implement not found for interface ITavernCoupon, forgot register?")
	}
	return localTavernCoupon
}

func RegisterTavernCoupon(i ITavernCoupon) { localTavernCoupon = i }

func TavernAdmin() ITavernAdmin {
	if localTavernAdmin == nil {
		panic("implement not found for interface ITavernAdmin, forgot register?")
	}
	return localTavernAdmin
}

func RegisterTavernAdmin(i ITavernAdmin) { localTavernAdmin = i }

func TavernStore() ITavernStore {
	if localTavernStore == nil {
		panic("implement not found for interface ITavernStore, forgot register?")
	}
	return localTavernStore
}

func RegisterTavernStore(i ITavernStore) { localTavernStore = i }

func TavernMember() ITavernMember {
	if localTavernMember == nil {
		panic("implement not found for interface ITavernMember, forgot register?")
	}
	return localTavernMember
}

func RegisterTavernMember(i ITavernMember) { localTavernMember = i }

func TavernMenu() ITavernMenu {
	if localTavernMenu == nil {
		panic("implement not found for interface ITavernMenu, forgot register?")
	}
	return localTavernMenu
}

func RegisterTavernMenu(i ITavernMenu) { localTavernMenu = i }

func TavernSession() ITavernSession {
	if localTavernSession == nil {
		panic("implement not found for interface ITavernSession, forgot register?")
	}
	return localTavernSession
}

func RegisterTavernSession(i ITavernSession) { localTavernSession = i }

func TavernCart() ITavernCart {
	if localTavernCart == nil {
		panic("implement not found for interface ITavernCart, forgot register?")
	}
	return localTavernCart
}

func RegisterTavernCart(i ITavernCart) { localTavernCart = i }

func TavernOrder() ITavernOrder {
	if localTavernOrder == nil {
		panic("implement not found for interface ITavernOrder, forgot register?")
	}
	return localTavernOrder
}

func RegisterTavernOrder(i ITavernOrder) { localTavernOrder = i }

func TavernDeposit() ITavernDeposit {
	if localTavernDeposit == nil {
		panic("implement not found for interface ITavernDeposit, forgot register?")
	}
	return localTavernDeposit
}

func RegisterTavernDeposit(i ITavernDeposit) { localTavernDeposit = i }

func TavernStaff() ITavernStaff {
	if localTavernStaff == nil {
		panic("implement not found for interface ITavernStaff, forgot register?")
	}
	return localTavernStaff
}

func RegisterTavernStaff(i ITavernStaff) { localTavernStaff = i }
