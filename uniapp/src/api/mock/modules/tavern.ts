/**
 * 酒馆小程序 MVP Mock（前端先行阶段的后端替身）
 *
 * 设计目标：
 *  - 用内存态模拟真实的先付制流程：开台 → 加购 → 提交订单 → 支付 → 制作 → 行级上齐 → 存酒
 *  - 覆盖多人在同一桌台会话下的共享点单（每人一个篮子、全桌可见、按菜聚合）
 *  - 路径与将来真实后端保持一致，后端就绪后只需关闭 mock
 *
 * 注意：支付为「模拟支付」，用于前端联调；生产必须走后端 + 微信支付回调。
 *      制作进度为了演示方便，把 make_minutes 当「秒」推进。
 */
import { defineMock } from '@alova/mock'

/* ---------------- 内存状态 ---------------- */

const now = () => new Date()
const fmt = (d: Date) => {
  const p = (n: number) => String(n).padStart(2, '0')
  return `${d.getFullYear()}-${p(d.getMonth() + 1)}-${p(d.getDate())} ${p(d.getHours())}:${p(d.getMinutes())}:${p(d.getSeconds())}`
}

const store = {
  id: 1,
  name: '晚风民谣小酒馆',
  address: '某某市某某区某某路 88 号',
  phone: '13800000000',
  businessOpen: '18:00',
  businessClose: '02:00',
  lastOrderGap: 30,
  payFirst: true,
  payExpireMinutes: 5,
}

const member = {
  id: 1001,
  nickname: '我',
  avatar: '',
  phone: '13800008888',
  isStaff: false,
  staffRole: '',
  points: 320,
  balance: 0,
}

/** 其他同桌成员（演示共享点单：他们的加购/下单会体现在全桌视图里） */
const mates = [
  { memberId: 1002, nickname: '小美', avatar: '', isOwner: false, joinAt: fmt(now()) },
  { memberId: 1003, nickname: '阿泽', avatar: '', isOwner: false, joinAt: fmt(now()) },
]

/* ---------------- 菜单 ---------------- */

const menu = [
  {
    id: 1,
    name: '精酿啤酒',
    icon: '',
    timeSlots: '',
    goods: [
      {
        id: 101,
        categoryId: 1,
        name: 'IPA 印度淡色艾尔',
        cover: '',
        brief: '柑橘与松脂香，苦度适中',
        price: 38,
        tags: ['招牌'],
        isWine: true,
        onlyDineIn: false,
        isCleared: false,
        makeMinutes: 2,
        status: 1,
        skus: [
          { id: 1011, goodsId: 101, name: '杯 (330ml)', price: 38, stock: -1, unitCups: 1, isDepositUnit: false },
          { id: 1012, goodsId: 101, name: '整瓶 (750ml)', price: 168, stock: 20, unitCups: 2.3, isDepositUnit: true },
        ],
      },
      {
        id: 102,
        categoryId: 1,
        name: '小麦白啤',
        cover: '',
        brief: '香蕉与丁香气息，清爽',
        price: 32,
        tags: ['热销'],
        isWine: true,
        onlyDineIn: false,
        isCleared: false,
        makeMinutes: 2,
        status: 1,
        skus: [
          { id: 1021, goodsId: 102, name: '杯 (500ml)', price: 32, stock: -1, unitCups: 1, isDepositUnit: false },
          { id: 1022, goodsId: 102, name: '整瓶 (1000ml)', price: 128, stock: 12, unitCups: 2, isDepositUnit: true },
        ],
      },
      {
        id: 103,
        categoryId: 1,
        name: '世涛黑啤',
        cover: '',
        brief: '咖啡与巧克力尾韵',
        price: 42,
        tags: [],
        isWine: true,
        onlyDineIn: false,
        isCleared: true,
        makeMinutes: 2,
        status: 1,
        skus: [
          { id: 1031, goodsId: 103, name: '杯 (330ml)', price: 42, stock: -1, unitCups: 1, isDepositUnit: false },
        ],
      },
    ],
  },
  {
    id: 2,
    name: '威士忌 / 洋酒',
    icon: '',
    timeSlots: '',
    goods: [
      {
        id: 201,
        categoryId: 2,
        name: '单一麦芽威士忌',
        cover: '',
        brief: '12 年，雪莉桶',
        price: 68,
        tags: ['招牌'],
        isWine: true,
        onlyDineIn: true,
        isCleared: false,
        makeMinutes: 1,
        status: 1,
        skus: [
          { id: 2011, goodsId: 201, name: '杯 (45ml)', price: 68, stock: -1, unitCups: 1, isDepositUnit: false },
          { id: 2012, goodsId: 201, name: '整瓶 (700ml)', price: 880, stock: 6, unitCups: 15.5, isDepositUnit: true },
        ],
      },
      {
        id: 202,
        categoryId: 2,
        name: '金酒',
        cover: '',
        brief: '杜松子与草本香',
        price: 58,
        tags: [],
        isWine: true,
        onlyDineIn: true,
        isCleared: false,
        makeMinutes: 1,
        status: 1,
        skus: [
          { id: 2021, goodsId: 202, name: '杯 (45ml)', price: 58, stock: -1, unitCups: 1, isDepositUnit: false },
        ],
      },
    ],
  },
  {
    id: 3,
    name: '佐酒小食',
    icon: '',
    timeSlots: '',
    goods: [
      {
        id: 301,
        categoryId: 3,
        name: '烤串拼盘',
        cover: '',
        brief: '8 串，含牛肉、鸡翅、香菇',
        price: 68,
        tags: ['热销'],
        isWine: false,
        onlyDineIn: false,
        isCleared: false,
        makeMinutes: 12,
        status: 1,
        skus: [
          { id: 3011, goodsId: 301, name: '标准份', price: 68, stock: 30, unitCups: 1, isDepositUnit: false },
        ],
      },
      {
        id: 302,
        categoryId: 3,
        name: '芝士薯角',
        cover: '',
        brief: '',
        price: 38,
        tags: [],
        isWine: false,
        onlyDineIn: false,
        isCleared: false,
        makeMinutes: 10,
        status: 1,
        skus: [
          { id: 3021, goodsId: 302, name: '标准份', price: 38, stock: 50, unitCups: 1, isDepositUnit: false },
        ],
      },
      {
        id: 303,
        categoryId: 3,
        name: '凉拌毛豆',
        cover: '',
        brief: '',
        price: 18,
        tags: [],
        isWine: false,
        onlyDineIn: false,
        isCleared: false,
        makeMinutes: 3,
        status: 1,
        skus: [
          { id: 3031, goodsId: 303, name: '标准份', price: 18, stock: 100, unitCups: 1, isDepositUnit: false },
        ],
      },
    ],
  },
]

const allGoods = menu.flatMap(c => c.goods)
const findGoods = (goodsId: number) => allGoods.find(g => g.id === goodsId)
const findSku = (goodsId: number, skuId: number) => findGoods(goodsId)?.skus.find(s => s.id === skuId)

/* ---------------- 会话 / 购物车 / 订单 状态 ---------------- */

let seq = 1
const genNo = (prefix: string) => `${prefix}${Date.now()}${String(seq++).padStart(3, '0')}`

interface MockSession {
  id: number
  sessionNo: string
  tableId: number
  tableName: string
  status: 1 | 2
  guestCount: number
  ownerMemberId: number
  openAt: string
  totalPaidAmount: number
}
let session: MockSession | null = null
let cart: any[] = []
let orders: any[] = []
let deposits: any[] = [
  {
    id: 9001,
    depositNo: 'DY2026001',
    labelCode: 'LK9001',
    memberId: 1001,
    memberName: '我',
    phoneTail: '8888',
    sourceType: 2,
    goodsName: '单一麦芽威士忌 12 年',
    spec: '700ml',
    degree: '40%',
    remainRatio: 68,
    slotCode: 'A-2-3',
    photo: '',
    status: 1,
    depositAt: '2026-09-18 21:20:00',
    expireAt: '2027-03-18 23:59:59',
  },
  {
    id: 9002,
    depositNo: 'DY2026002',
    labelCode: 'LK9002',
    memberId: 1001,
    memberName: '我',
    phoneTail: '8888',
    sourceType: 1,
    goodsName: '自带 山崎 12 年',
    spec: '700ml',
    degree: '43%',
    remainRatio: 25,
    slotCode: 'B-1-2',
    photo: '',
    status: 1,
    depositAt: '2026-08-02 22:05:00',
    expireAt: '2026-12-30 23:59:59',
  },
]
let takeCodes: Record<string, { memberId: number, depositIds: number[], expireAt: number }> = {}

const ok = (data: any = null, message = 'ok') => ({ code: 0, message, data })
const fail = (message: string, status = 400) => ({ status, body: { code: -1, message, data: null } })

/** 演示用：把 make_minutes 当作「秒」推进行级制作状态 */
function advanceMaking() {
  const t = Date.now()
  orders.forEach((o) => {
    if (o.payStatus !== 1) return
    o.items.forEach((it: any) => {
      if (it.makerStatus >= 2) return
      const goods = findGoods(it.goodsId)
      const dur = Math.max(1, (goods?.makeMinutes ?? 5)) * 1000
      const paidAt = new Date(o.payTime.replace(/-/g, '/')).getTime()
      if (t - paidAt >= dur) {
        it.makerStatus = 2
        it.servedAt = fmt(now())
      }
      else {
        it.makerStatus = 1
      }
    })
    const allServed = o.items.every((it: any) => it.makerStatus >= 2)
    o.makingStatus = allServed ? (o.makingStatus >= 2 ? 3 : 2) : 1
  })
  // 待支付超时自动取消（先付制：5 分钟）
  orders.forEach((o) => {
    if (o.payStatus === 0 && new Date(o.expireAt.replace(/-/g, '/')).getTime() < t) {
      o.payStatus = 0
      o.makingStatus = 0
      o.expired = true
    }
  })
}

/** 全桌已点聚合（按菜合并，可追溯是谁点的） */
function buildAgg() {
  const map = new Map<string, any>()
  orders
    .filter(o => o.payStatus === 1 || o.payStatus === 0)
    .forEach((o) => {
      o.items.forEach((it: any) => {
        const key = `${it.goodsId}-${it.skuId}`
        const cur = map.get(key) || {
          goodsId: it.goodsId,
          skuId: it.skuId,
          goodsName: it.goodsName,
          skuName: it.skuName,
          qty: 0,
          servedQty: 0,
          byMembers: [],
        }
        cur.qty += it.qty
        if (it.makerStatus >= 2) cur.servedQty += it.qty
        if (!cur.byMembers.includes(o.memberName)) cur.byMembers.push(o.memberName)
        map.set(key, cur)
      })
    })
  return [...map.values()]
}

function snapshot() {
  if (!session) return null
  advanceMaking()
  const members = [
    { memberId: member.id, nickname: member.nickname, avatar: member.avatar, isOwner: session.ownerMemberId === member.id, joinAt: session.openAt },
    ...mates,
  ]
  return {
    session: { ...session, members, orderCount: orders.filter(o => o.payStatus === 1).length },
    members,
    cart,
    orders,
    orderedAgg: buildAgg(),
  }
}

/* ---------------- Mock 定义 ---------------- */

export default defineMock({
  '[GET]/tavern/store/info': () => ok(store),

  '[GET]/tavern/member/profile': () => ok(member),

  '[GET]/tavern/menu/list': () => ok({ list: menu }),

  // 扫码开台 / 二次扫码回到原会话
  '[POST]/tavern/session/open': ({ data }: any) => {
    const tableNo = data?.tableNo || 'A01'
    if (session && session.status === 1) {
      return ok({ ...session, members: snapshot()!.members, orderCount: orders.filter(o => o.payStatus === 1).length, totalPaidAmount: session.totalPaidAmount })
    }
    session = {
      id: 5001,
      sessionNo: genNo('SS'),
      tableId: Number(tableNo.replace(/\D/g, '')) || 1,
      tableName: tableNo,
      status: 1,
      guestCount: 3,
      ownerMemberId: member.id,
      openAt: fmt(now()),
      totalPaidAmount: 0,
    }
    cart = []
    orders = []
    // 演示共享点单：同桌成员已经点了两样
    const mateOrder = {
      id: 7000,
      orderNo: genNo('OD'),
      sessionId: session.id,
      tableId: session.tableId,
      tableName: session.tableName,
      orderSeq: 1,
      memberId: 1002,
      memberName: '小美',
      payerMemberId: 1002,
      payStatus: 1,
      makingStatus: 2,
      refundStatus: 0,
      payChannel: 'mock',
      totalAmount: 64,
      payAmount: 64,
      discountAmount: 0,
      expireAt: fmt(now()),
      payTime: fmt(now()),
      createdAt: fmt(now()),
      items: [
        { id: 8001, orderId: 7000, goodsId: 102, skuId: 1021, goodsName: '小麦白啤', skuName: '杯 (500ml)', unitPrice: 32, qty: 1, amount: 32, isWine: true, makerStatus: 2, servedAt: fmt(now()), remark: '' },
        { id: 8002, orderId: 7000, goodsId: 3031, skuId: 3031, goodsName: '凉拌毛豆', skuName: '标准份', unitPrice: 18, qty: 1, amount: 18, isWine: false, makerStatus: 2, servedAt: fmt(now()), remark: '不要香菜' },
      ],
    }
    orders.push(mateOrder)
    session.totalPaidAmount = 64
    return ok({ ...session, members: snapshot()!.members, orderCount: 1, totalPaidAmount: 64 })
  },

  '[GET]/tavern/session/snapshot': () => {
    const s = snapshot()
    return s ? ok(s) : fail('会话不存在')
  },

  '[POST]/tavern/session/close': () => {
    if (session) session.status = 2
    const s = session ? { ...session } : null
    session = null
    cart = []
    orders = []
    return ok(s)
  },

  /* -------- 购物车：每人一个篮子，全桌可见 -------- */
  '[POST]/tavern/cart/add': ({ data }: any) => {
    if (!session) return fail('请先扫码开台')
    const goods = findGoods(data.goodsId)
    const sku = findSku(data.goodsId, data.skuId)
    if (!goods || !sku) return fail('商品不存在')
    if (goods.isCleared) return fail(`${goods.name} 已沽清`)
    const exist = cart.find(c => c.memberId === member.id && c.goodsId === data.goodsId && c.skuId === data.skuId && c.remark === (data.remark || ''))
    if (exist) {
      exist.qty += data.qty || 1
    }
    else {
      cart.push({
        id: Date.now() + Math.floor(Math.random() * 1000),
        sessionId: session.id,
        memberId: member.id,
        memberName: member.nickname,
        goodsId: goods.id,
        skuId: sku.id,
        goodsName: goods.name,
        skuName: sku.name,
        price: sku.price,
        qty: data.qty || 1,
        remark: data.remark || '',
      })
    }
    return ok(null)
  },

  '[POST]/tavern/cart/update': ({ data }: any) => {
    const item = cart.find(c => c.id === data.id)
    if (!item) return fail('购物车行不存在')
    // 归属不可越权：只能改自己的行
    if (item.memberId !== member.id) return fail('只能修改自己添加的商品')
    item.qty = data.qty
    if (data.remark !== undefined) item.remark = data.remark
    return ok(null)
  },

  '[POST]/tavern/cart/remove': ({ data }: any) => {
    const item = cart.find(c => c.id === data.id)
    if (!item) return fail('购物车行不存在')
    if (item.memberId !== member.id) return fail('只能删除自己添加的商品')
    cart = cart.filter(c => c.id !== data.id)
    return ok(null)
  },

  /* -------- 订单：先付制 -------- */
  '[POST]/tavern/order/submit': ({ data }: any) => {
    if (!session) return fail('请先扫码开台')
    const ids: number[] = data.cartIds || []
    const rows = cart.filter(c => ids.includes(c.id))
    if (!rows.length) return fail('请选择要下单的商品')
    // 行级校验：沽清 / 下架只让该行失败
    const bad = rows.find(r => findGoods(r.goodsId)?.isCleared)
    if (bad) return fail(`${bad.goodsName} 已沽清，请从购物车中移除后再下单`)

    const orderId = Date.now()
    const orderNo = genNo('OD')
    const expire = new Date(Date.now() + store.payExpireMinutes * 60 * 1000)
    const order = {
      id: orderId,
      orderNo,
      sessionId: session.id,
      tableId: session.tableId,
      tableName: session.tableName,
      orderSeq: orders.length + 1,
      memberId: member.id,
      memberName: member.nickname,
      payerMemberId: member.id,
      payStatus: 0,
      makingStatus: 0,
      refundStatus: 0,
      payChannel: '',
      totalAmount: rows.reduce((s: number, r: any) => s + r.price * r.qty, 0),
      payAmount: rows.reduce((s: number, r: any) => s + r.price * r.qty, 0),
      discountAmount: 0,
      expireAt: fmt(expire),
      payTime: '',
      createdAt: fmt(now()),
      items: rows.map((r: any, i: number) => ({
        id: orderId * 10 + i,
        orderId,
        goodsId: r.goodsId,
        skuId: r.skuId,
        goodsName: r.goodsName,
        skuName: r.skuName,
        unitPrice: r.price,
        qty: r.qty,
        amount: r.price * r.qty,
        isWine: !!findGoods(r.goodsId)?.isWine,
        makerStatus: 0,
        servedAt: '',
        remark: r.remark,
      })),
    }
    orders.push(order)
    // 提交只清掉已下单的行
    cart = cart.filter(c => !ids.includes(c.id))
    return ok(order)
  },

  // 模拟支付：真实环境替换为微信 JSAPI 预支付 + 回调
  '[POST]/tavern/order/pay': ({ data }: any) => {
    const order = orders.find(o => o.orderNo === data.orderNo)
    if (!order) return fail('订单不存在')
    if (order.payStatus === 1) return order // 幂等
    order.payStatus = 1
    order.makingStatus = 1
    order.payChannel = data.channel || 'mock'
    order.payTime = fmt(now())
    order.items.forEach((it: any) => { it.makerStatus = 1 })
    if (session) session.totalPaidAmount += order.payAmount
    return ok(order)
  },

  '[GET]/tavern/order/detail': ({ query }: any) => {
    advanceMaking()
    const order = orders.find(o => o.orderNo === query.orderNo)
    return order ? ok(order) : fail('订单不存在')
  },

  '[GET]/tavern/order/list': ({ query }: any) => {
    advanceMaking()
    let list = orders.slice().reverse()
    if (query.sessionId) list = list.filter(o => String(o.sessionId) === String(query.sessionId))
    if (query.mine === true || query.mine === 'true') list = list.filter(o => o.memberId === member.id)
    return ok({ list })
  },

  /* -------- 存酒 -------- */
  '[GET]/tavern/bottle/list': () => ok({ list: deposits.filter(d => d.memberId === member.id) }),

  '[POST]/tavern/bottle/takeCode': ({ data }: any) => {
    const code = String(Math.floor(100000 + Math.random() * 900000))
    takeCodes[code] = { memberId: member.id, depositIds: data.depositIds || [], expireAt: Date.now() + 30 * 1000 }
    return ok({ code, expireSeconds: 30, depositIds: data.depositIds || [] })
  },

  '[POST]/tavern/bottle/depositFromOrder': ({ data }: any) => {
    const order = orders.find(o => o.id === data.orderId)
    if (!order || order.payStatus !== 1) return fail('仅支付成功的订单可转存')
    const item = order.items.find((it: any) => (data.itemIds || []).includes(it.id))
    if (!item) return fail('请选择要寄存的酒品')
    const dep = {
      id: Date.now(),
      depositNo: genNo('DY'),
      labelCode: genNo('LK'),
      memberId: member.id,
      memberName: member.nickname,
      phoneTail: member.phone.slice(-4),
      sourceType: 3,
      goodsName: item.goodsName,
      spec: item.skuName,
      degree: '',
      remainRatio: 100,
      slotCode: '待分配',
      photo: '',
      status: 1,
      depositAt: fmt(now()),
      expireAt: fmt(new Date(Date.now() + 180 * 24 * 3600 * 1000)),
    }
    deposits.unshift(dep)
    return ok(dep)
  },

  /* -------- 员工工作台（与客人同一小程序，凭会员 is_staff + 角色） -------- */

  // 演示用：切换员工身份，便于在前端验证工作台（真实环境由后台标记决定）
  '[POST]/tavern/member/switchStaff': ({ data }: any) => {
    member.isStaff = !!data.isStaff
    member.staffRole = member.isStaff ? (data.role || 'waiter') : ''
    return ok(member)
  },

  // 桌台看板
  '[GET]/tavern/staff/tables': () => {
    advanceMaking()
    const making = orders.filter(o => o.payStatus === 1).flatMap((o: any) => o.items.filter((i: any) => i.makerStatus < 2)).length
    return ok({ list: [
      {
        tableId: session?.tableId || 1,
        tableName: session?.tableName || 'A01',
        seatStatus: session ? 2 : 1,
        guestCount: session?.guestCount || 3,
        openAt: session?.openAt || '',
        sessionId: session?.id || 0,
        paidAmount: session?.totalPaidAmount || 0,
        orderCount: orders.filter(o => o.payStatus === 1).length,
        pendingPayCount: orders.filter(o => o.payStatus === 0).length,
        makingCount: making,
        called: false,
      },
      { tableId: 2, tableName: 'A02', seatStatus: 1, guestCount: 0, openAt: '', sessionId: 0, paidAmount: 0, orderCount: 0, pendingPayCount: 0, makingCount: 0, called: false },
      { tableId: 3, tableName: 'A03', seatStatus: 3, guestCount: 0, openAt: '', sessionId: 0, paidAmount: 0, orderCount: 0, pendingPayCount: 0, makingCount: 0, called: false },
      { tableId: 4, tableName: 'B01', seatStatus: 1, guestCount: 0, openAt: '', sessionId: 0, paidAmount: 0, orderCount: 0, pendingPayCount: 0, makingCount: 0, called: false },
    ] })
  },

  // 出票区：已支付订单（先付制下支付成功即出票，可上齐/加急，不可拒接）
  '[GET]/tavern/staff/tickets': () => {
    advanceMaking()
    return ok({ list: orders.filter(o => o.payStatus === 1 && o.makingStatus < 3).map(o => ({ ...o, urgent: !!o.urgent })) })
  },

  // 行级上齐 / 整单上齐
  '[POST]/tavern/staff/serve': ({ data }: any) => {
    const order = orders.find(o => o.orderNo === data.orderNo)
    if (!order) return fail('订单不存在')
    const ids: number[] = data.itemIds || []
    order.items.forEach((it: any) => {
      if (!ids.length || ids.includes(it.id)) {
        it.makerStatus = 2
        it.servedAt = fmt(now())
      }
    })
    order.makingStatus = order.items.every((it: any) => it.makerStatus >= 2) ? 2 : 1
    return ok(order)
  },

  // 加急
  '[POST]/tavern/staff/urgent': ({ data }: any) => {
    const order = orders.find(o => o.orderNo === data.orderNo)
    if (!order) return fail('订单不存在')
    order.urgent = true
    return ok(null)
  },

  // 沽清（即时对客人端生效）
  '[POST]/tavern/staff/goods/clear': ({ data }: any) => {
    const goods = findGoods(data.goodsId)
    if (!goods) return fail('商品不存在')
    goods.isCleared = !!data.isCleared
    return ok(goods)
  },

  '[POST]/tavern/member/bindPhone': ({ data }: any) => {
    member.phone = data.phone
    return ok(member)
  },

  '[POST]/tavern/session/call': () => ok(null),

  '[POST]/tavern/order/refund/apply': ({ data }: any) => {
    const order = orders.find(o => o.orderNo === data.orderNo)
    if (!order) return fail('订单不存在')
    if (order.payStatus !== 1) return fail('仅已支付订单可申请退款')
    const making = order.items.some((i: any) => i.makerStatus > 0)
    order.refundReason = data.reason
    order.refundAmount = order.payAmount
    order.refundStatus = making ? 1 : 4
    if (!making) order.payStatus = 3
    return ok(order)
  },

  '[POST]/tavern/order/refund/audit': ({ data }: any) => {
    const order = orders.find(o => o.orderNo === data.orderNo)
    if (!order) return fail('订单不存在')
    if (order.refundStatus !== 1) return fail('该订单当前不在待审核状态')
    order.refundStatus = data.approve ? 4 : 3
    if (data.approve) order.payStatus = 2
    return ok(order)
  },

  '[POST]/tavern/staff/bottle/deposit': ({ data }: any) => {
    const dep = {
      id: Date.now(),
      depositNo: genNo('DY'),
      labelCode: genNo('LK'),
      memberId: member.id,
      memberName: member.nickname,
      phoneTail: (data.phone || member.phone || '').slice(-4),
      sourceType: 1,
      goodsName: data.goodsName,
      spec: data.spec || '',
      degree: data.degree || '',
      remainRatio: data.remainRatio || 100,
      slotCode: data.slotCode || '待分配',
      photo: '',
      status: 1,
      depositAt: fmt(now()),
      expireAt: fmt(new Date(Date.now() + 180 * 24 * 3600 * 1000)),
    }
    deposits.unshift(dep)
    return ok(dep)
  },

  '[POST]/tavern/staff/table/move': ({ data }: any) => {
    if (session && data.toTableNo) session.tableName = data.toTableNo
    return ok(null)
  },

  '[POST]/tavern/staff/call/handle': () => ok(null),

  // 存酒核销：取酒码 + 店员确认（手机号后四位由店员现场核对）
  '[POST]/tavern/staff/bottle/verify': ({ data }: any) => {
    const rec = takeCodes[data.code]
    if (!rec || rec.expireAt < Date.now()) return fail('取酒码无效或已过期')
    const dep = deposits.find(d => d.id === (data.depositId || rec.depositIds[0]))
    if (!dep) return fail('存酒记录不存在')
    if (dep.status !== 1) return fail('该存酒已不在库')
    const takeRatio = Number(data.takeRatio || dep.remainRatio)
    dep.remainRatio = Math.max(0, Number((dep.remainRatio - takeRatio).toFixed(2)))
    if (dep.remainRatio <= 0) dep.status = 2
    deposits = deposits.map(d => (d.id === dep.id ? dep : d))
    return ok({ deposit: dep, memberId: rec.memberId, takeRatio })
  },
})
