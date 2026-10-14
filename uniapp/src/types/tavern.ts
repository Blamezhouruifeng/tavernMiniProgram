/** 酒馆小程序领域类型（与后端 hg_tavern_* 表对齐） */

export interface ApiRes<T = unknown> {
  code: number
  message: string
  data: T
}

/** 会员（客人与员工共用同一套会员体系） */
export interface Member {
  id: number
  nickname: string
  avatar: string
  phone: string
  isStaff: boolean
  staffRole: string
  points: number
  balance: number
}

export type StaffRole = 'waiter' | 'bar' | 'kitchen' | 'manager' | ''

/** 门店 */
export interface Store {
  id: number
  name: string
  address: string
  phone: string
  businessOpen: string
  businessClose: string
  lastOrderGap: number
  payFirst: boolean
  payExpireMinutes: number
}

/** 桌台会话成员 */
export interface SessionMember {
  memberId: number
  nickname: string
  avatar: string
  isOwner: boolean
  joinAt: string
}

/** 桌台会话 */
export interface TableSession {
  id: number
  sessionNo: string
  tableId: number
  tableName: string
  status: 1 | 2
  guestCount: number
  ownerMemberId: number
  members: SessionMember[]
  openAt: string
  orderCount: number
  totalPaidAmount: number
}

/** 商品规格 */
export interface GoodsSku {
  id: number
  goodsId: number
  name: string
  price: number
  stock: number
  unitCups: number
  isDepositUnit: boolean
}

/** 商品 */
export interface Goods {
  id: number
  categoryId: number
  name: string
  cover: string
  brief: string
  price: number
  tags: string[]
  isWine: boolean
  onlyDineIn: boolean
  isCleared: boolean
  makeMinutes: number
  status: 1 | 2
  skus: GoodsSku[]
}

/** 商品分类 */
export interface Category {
  id: number
  name: string
  icon: string
  timeSlots: string
  goods: Goods[]
}

/** 购物车行（归属加购人） */
export interface CartItem {
  id: number
  sessionId: number
  memberId: number
  memberName: string
  goodsId: number
  skuId: number
  goodsName: string
  skuName: string
  price: number
  qty: number
  remark: string
}

/** 行级制作状态：0待制作 1制作中 2已上齐 */
export type MakerStatus = 0 | 1 | 2

export interface OrderItem {
  id: number
  orderId: number
  goodsId: number
  skuId: number
  goodsName: string
  skuName: string
  unitPrice: number
  qty: number
  amount: number
  isWine: boolean
  makerStatus: MakerStatus
  servedAt: string
  remark: string
}

/** 支付状态：0未支付 1已支付 2部分退款 3已退款 */
export type PayStatus = 0 | 1 | 2 | 3
/** 制作状态：0未开始 1制作中 2已上齐 3已完成 */
export type MakingStatus = 0 | 1 | 2 | 3
/** 退款状态：0无 1申请中 2已通过 3已拒绝 4已退款 */
export type RefundStatus = 0 | 1 | 2 | 3 | 4

export interface Order {
  id: number
  orderNo: string
  sessionId: number
  tableId: number
  tableName: string
  orderSeq: number
  memberId: number
  memberName: string
  payerMemberId: number
  payStatus: PayStatus
  makingStatus: MakingStatus
  refundStatus: RefundStatus
  payChannel: string
  totalAmount: number
  payAmount: number
  discountAmount: number
  expireAt: string
  payTime: string
  createdAt: string
  /** 退款金额（部分退款按明细行累加） */
  refundAmount?: number
  /** 退款原因 */
  refundReason?: string
  items: OrderItem[]
}

/** 会话快照：菜单页与桌台聚合用 */
export interface SessionSnapshot {
  session: TableSession
  members: SessionMember[]
  cart: CartItem[]
  orders: Order[]
  /** 全桌已点聚合（按菜合并） */
  orderedAgg: OrderedAggItem[]
}

export interface OrderedAggItem {
  goodsId: number
  skuId: number
  goodsName: string
  skuName: string
  qty: number
  /** 已上齐数量 */
  servedQty: number
  /** 谁点的（可能多人） */
  byMembers: string[]
}

/** 存酒 */
export interface BottleDeposit {
  id: number
  depositNo: string
  labelCode: string
  memberId: number
  memberName: string
  phoneTail: string
  sourceType: 1 | 2 | 3
  goodsName: string
  spec: string
  degree: string
  remainRatio: number
  slotCode: string
  photo: string
  status: 1 | 2 | 3 | 4 | 5
  depositAt: string
  expireAt: string
}

/** 优惠券 */
export interface Coupon {
  id: number
  couponId: number
  name: string
  /** 1满减 2折扣 */
  type: number
  amount: number
  minAmount: number
  discount: number
  code: string
  /** 1新人礼 2活动 3后台定向 */
  source: number
  /** 1未使用 2已使用 3已过期 */
  status: number
  expireAt: string
  usable: boolean
  unusableReason: string
}

/** 取酒码 */
export interface TakeCode {
  code: string
  expireSeconds: number
  depositIds: number[]
}
