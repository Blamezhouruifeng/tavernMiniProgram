/**
 * 酒馆小程序 API（MVP）
 * 说明：当前阶段全部由本地 mock 拦截（uniapp/src/mock/modules/tavern.ts），
 * 后端就绪后关闭 mock 并保持路径不变即可直连。
 */
import type {
  ApiRes,
  BottleDeposit,
  Category,
  Member,
  Order,
  SessionSnapshot,
  Store,
  TableSession,
  TakeCode,
} from '@/types/tavern'
import alovaInstance from '../core/instance'

/* ---------------- 门店与会员 ---------------- */

/** 门店信息（含先付制、打烊截止配置） */
export function getStore() {
  return alovaInstance.Get<ApiRes<Store>>('/tavern/store/info')
}

/** 当前会员信息（含 is_staff，用于决定进入客人视图还是工作台） */
export function getMemberProfile() {
  return alovaInstance.Get<ApiRes<Member>>('/tavern/member/profile')
}

/* ---------------- 桌台会话 ---------------- */

/** 扫码开台 / 进入桌台（二次扫码回到原会话） */
export function openTableSession(data: { tableNo: string, memberId?: number }) {
  return alovaInstance.Post<ApiRes<TableSession>>('/tavern/session/open', data)
}

/** 桌台快照：成员、购物车、订单、全桌已点聚合（轮询兜底） */
export function getSessionSnapshot(params: { sessionId: number }) {
  return alovaInstance.Get<ApiRes<SessionSnapshot>>('/tavern/session/snapshot', { params })
}

/** 清台（店员） */
export function closeTableSession(data: { sessionId: number }) {
  return alovaInstance.Post<ApiRes<null>>('/tavern/session/close', data)
}

/* ---------------- 菜单 ---------------- */

/** 菜单（分类 + 商品 + 规格），分时段由后端过滤 */
export function getMenu() {
  return alovaInstance.Get<ApiRes<{ list: Category[] }>>('/tavern/menu/list')
}

/* ---------------- 购物车（每人一个篮子，全桌可见） ---------------- */

export function addCart(data: { sessionId: number, goodsId: number, skuId: number, qty: number, remark?: string }) {
  return alovaInstance.Post<ApiRes<null>>('/tavern/cart/add', data)
}

export function updateCart(data: { id: number, qty: number, remark?: string }) {
  return alovaInstance.Post<ApiRes<null>>('/tavern/cart/update', data)
}

export function removeCart(data: { id: number }) {
  return alovaInstance.Post<ApiRes<null>>('/tavern/cart/remove', data)
}

/* ---------------- 订单（先付制） ---------------- */

/** 提交订单 → 返回待支付订单（超时时间由后端下发） */
export function submitOrder(data: { sessionId: number, cartIds: number[], remark?: string }) {
  return alovaInstance.Post<ApiRes<Order>>('/tavern/order/submit', data)
}

/** 模拟支付（微信支付商户号就绪后替换为 JSAPI 预支付） */
export function payOrder(data: { orderNo: string, channel?: string }) {
  return alovaInstance.Post<ApiRes<Order>>('/tavern/order/pay', data)
}

export function getOrderDetail(params: { orderNo: string }) {
  return alovaInstance.Get<ApiRes<Order>>('/tavern/order/detail', { params })
}

export function getOrderList(params: { sessionId?: number, mine?: boolean }) {
  return alovaInstance.Get<ApiRes<{ list: Order[] }>>('/tavern/order/list', { params })
}

/* ---------------- 存酒 ---------------- */

export function getBottleList() {
  return alovaInstance.Get<ApiRes<{ list: BottleDeposit[] }>>('/tavern/bottle/list')
}

/** 生成动态取酒码（30 秒有效），到店由店员核销 */
export function createTakeCode(data: { depositIds: number[] }) {
  return alovaInstance.Post<ApiRes<TakeCode>>('/tavern/bottle/takeCode', data)
}

/* ---------------- 员工工作台（与客人同一小程序，凭会员 is_staff + 角色） ---------------- */

export interface StaffTable {
  tableId: number
  tableName: string
  /** 1空闲 2用餐中 3待清台 */
  seatStatus: number
  guestCount: number
  openAt: string
  sessionId: number
  paidAmount: number
  orderCount: number
  pendingPayCount: number
  makingCount: number
  called: boolean
}

/** 桌台看板 */
export function getStaffTables() {
  return alovaInstance.Get<ApiRes<{ list: StaffTable[] }>>('/tavern/staff/tables')
}

/** 出票区（已支付订单，支付即出票） */
export function getStaffTickets() {
  return alovaInstance.Get<ApiRes<{ list: Order[] }>>('/tavern/staff/tickets')
}

/** 行级 / 整单上齐 */
export function staffServe(data: { orderNo: string, itemIds?: number[] }) {
  return alovaInstance.Post<ApiRes<Order>>('/tavern/staff/serve', data)
}

/** 加急 */
export function staffUrgent(data: { orderNo: string }) {
  return alovaInstance.Post<ApiRes<null>>('/tavern/staff/urgent', data)
}

/** 沽清切换（即时对客人端生效） */
export function staffClearGoods(data: { goodsId: number, isCleared: boolean }) {
  return alovaInstance.Post<ApiRes<null>>('/tavern/staff/goods/clear', data)
}

/** 存酒核销（手机号 + 取酒码 + 店员确认） */
export function staffVerifyBottle(data: { code: string, depositId?: number, takeRatio?: number }) {
  return alovaInstance.Post<ApiRes<{ deposit: BottleDeposit, memberId: number, takeRatio: number }>>('/tavern/staff/bottle/verify', data)
}

/** 绑定手机号（存酒取酒凭证的前置条件） */
export function bindPhone(data: { phone: string }) {
  return alovaInstance.Post<ApiRes<Member>>('/tavern/member/bindPhone', data)
}

/** 呼叫服务员 / 催单 */
export function callWaiter(data: { sessionId: number, type: 'waiter' | 'urge', orderNo?: string, remark?: string }) {
  return alovaInstance.Post<ApiRes<null>>('/tavern/session/call', data)
}

/** 申请退款（按明细行；制作前可自助通过） */
export function refundApply(data: { orderNo: string, itemIds?: number[], reason: string }) {
  return alovaInstance.Post<ApiRes<Order>>('/tavern/order/refund/apply', data)
}

/** 退款审核（员工） */
export function refundAudit(data: { orderNo: string, approve: boolean, reason?: string }) {
  return alovaInstance.Post<ApiRes<Order>>('/tavern/order/refund/audit', data)
}

/** 员工：代存登记（自带酒水） */
export function staffDepositRegister(data: { phone?: string, memberId?: number, goodsName: string, spec?: string, degree?: string, remainRatio?: number, slotCode?: string, photo?: string, remark?: string }) {
  return alovaInstance.Post<ApiRes<BottleDeposit>>('/tavern/staff/bottle/deposit', data)
}

/** 员工：转台/并台 */
export function staffMoveTable(data: { sessionId: number, toTableNo: string }) {
  return alovaInstance.Post<ApiRes<null>>('/tavern/staff/table/move', data)
}

/** 员工：响应呼叫 */
export function staffHandleCall(data: { sessionId: number }) {
  return alovaInstance.Post<ApiRes<null>>('/tavern/staff/call/handle', data)
}

/** 演示用：切换员工身份（真实环境由后台「会员标记为员工」决定） */
export function switchStaffRole(data: { isStaff: boolean, role?: string }) {
  return alovaInstance.Post<ApiRes<Member>>('/tavern/member/switchStaff', data)
}

/** 退酒入存（支付后的整瓶酒直接转为存酒） */
export function depositFromOrder(data: { orderId: number, itemIds: number[] }) {
  return alovaInstance.Post<ApiRes<BottleDeposit>>('/tavern/bottle/depositFromOrder', data)
}
