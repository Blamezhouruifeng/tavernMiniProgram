/**
 * 酒馆点餐状态：身份、桌台会话、购物车（每人一个篮子、全桌可见）、订单
 *
 * 约定：
 *  - 购物车归属：每人一个篮子，只能改自己的行；全桌可看
 *  - 提交只清掉已下单的行，别人的购物车不受影响
 *  - 支付采用先付制：支付成功 = 进入制作
 */
import { defineStore } from 'pinia'
import { ensureLogin } from '@/api/core/auth'
import type { CartItem, Member, Order, SessionSnapshot, Store, TableSession } from '@/types/tavern'
import {
  addCart,
  getMemberProfile,
  getSessionSnapshot,
  getStore,
  openTableSession,
  payOrder,
  removeCart,
  submitOrder,
  updateCart,
} from '@/api/tavern'

export const useTavernStore = defineStore('tavern', {
  state: () => ({
    /** 门店配置（含先付制、超时分钟数） */
    store: null as Store | null,
    /** 当前会员（客人与员工共用同一套会员体系，is_staff 决定视图） */
    member: null as Member | null,
    /** 当前桌台会话 */
    session: null as TableSession | null,
    /** 会话快照（成员 / 购物车 / 订单 / 全桌已点聚合） */
    snapshot: null as SessionSnapshot | null,
    loading: false,
    lastError: '',
    /** 待支付订单（提交后留在前端，用于超时倒计时） */
    pendingOrder: null as Order | null,
  }),

  getters: {
    isStaff: state => !!state.member?.isStaff,
    tableName: state => state.session?.tableName || '',
    /** 我的购物车（只有我能改） */
    myCart: (state): CartItem[] => state.snapshot?.cart.filter(c => c.memberId === state.member?.id) || [],
    /** 同桌其他人的购物车 */
    othersCart: (state): CartItem[] => state.snapshot?.cart.filter(c => c.memberId !== state.member?.id) || [],
    myCartAmount(): number {
      return this.myCart.reduce((s, c) => s + c.price * c.qty, 0)
    },
    myCartCount(): number {
      return this.myCart.reduce((s, c) => s + c.qty, 0)
    },
    /** 全桌已点（按菜聚合，含"谁点的"） */
    orderedAgg: state => state.snapshot?.orderedAgg || [],
    /** 全桌已付订单 */
    paidOrders: (state): Order[] => state.snapshot?.orders.filter(o => o.payStatus === 1) || [],
    /** 我下的订单 */
    myOrders: (state): Order[] => state.snapshot?.orders.filter(o => o.memberId === state.member?.id) || [],
    /** 全桌已付累计 */
    tablePaidAmount: state => state.snapshot?.session.totalPaidAmount || 0,
  },

  actions: {
    /** 启动时先确保登录，再拉取门店与会员身份，决定视图（客人 / 工作台） */
    async bootstrap() {
      await ensureLogin()
      const [s, m] = await Promise.all([getStore(), getMemberProfile()])
      this.store = s.data
      this.member = m.data
    },

    /** 扫码进桌 / 二次扫码回到原会话 */
    async openTable(tableNo: string) {
      const res = await openTableSession({ tableNo })
      this.session = res.data
      await this.refresh()
      return res.data
    },

    /** 拉取会话快照（轮询兜底；后续接入 WebSocket 增量） */
    async refresh() {
      if (!this.session)
        return
      const res = await getSessionSnapshot({ sessionId: this.session.id })
      this.snapshot = res.data
      this.session = res.data.session
    },

    async addToCart(payload: { goodsId: number, skuId: number, qty: number, remark?: string }) {
      if (!this.session)
        throw new Error('请先扫码开台')
      await addCart({ sessionId: this.session.id, ...payload })
      await this.refresh()
    },

    async updateCartQty(id: number, qty: number) {
      await updateCart({ id, qty })
      await this.refresh()
    },

    async removeCartItem(id: number) {
      await removeCart({ id })
      await this.refresh()
    },

    /** 提交订单（先付制：下单 → 待支付） */
    async submit(cartIds: number[], remark?: string) {
      if (!this.session)
        throw new Error('请先扫码开台')
      const res = await submitOrder({ sessionId: this.session.id, cartIds, remark })
      this.pendingOrder = res.data
      await this.refresh()
      return res.data
    },

    /** 模拟支付：支付成功即进入制作 */
    async pay(orderNo: string) {
      const res = await payOrder({ orderNo })
      this.pendingOrder = null
      await this.refresh()
      return res.data
    },
  },
})
