<script setup lang="ts">
/**
 * 扫码点餐页（先付制）
 * 流程：扫桌码进桌 → 选规格加购 → 与同桌共享可见 → 提交订单 → 支付 → 进入制作
 */
import type { Category, Goods } from '@/types/tavern'
import type { Coupon } from '@/types/tavern'
import { getCouponList, getMenu } from '@/api/tavern'
import { useTavernStore } from '@/store/tavernStore'
import { useGlobalToast } from '@/composables/useGlobalToast'

definePage({
  name: 'menu',
  style: {
    navigationBarTitleText: '点餐',
  },
})

const router = useRouter()
const route = useRoute()
const toast = useGlobalToast()
const tavern = useTavernStore()

const manualTableNo = ref('')
const categories = ref<Category[]>([])
const activeCategory = ref(0)
const loading = ref(true)

/* ---------- 规格选择弹层 ---------- */
const skuVisible = ref(false)
const currentGoods = ref<Goods | null>(null)
const currentSkuId = ref(0)
const currentQty = ref(1)
const currentRemark = ref('')

/* ---------- 购物车弹层 ---------- */
const cartVisible = ref(false)
const cartTab = ref<'mine' | 'table'>('mine')

/* ---------- 优惠券 ---------- */
const coupons = ref<Coupon[]>([])
const selectedCouponId = ref(0)
const couponVisible = ref(false)

const selectedCoupon = computed(() => coupons.value.find(c => c.id === selectedCouponId.value))
/** 预计优惠金额 */
const discount = computed(() => {
  const c = selectedCoupon.value
  if (!c)
    return 0
  if (myCartAmount.value < c.minAmount)
    return 0
  if (c.type === 2)
    return Number((myCartAmount.value * (1 - c.discount)).toFixed(2))
  return Math.min(c.amount, myCartAmount.value)
})
/** 应付金额 */
const payableAmount = computed(() => Math.max(0, Number((myCartAmount.value - discount.value).toFixed(2))))

async function loadCoupons() {
  try {
    const res = await getCouponList({ status: 1, amount: myCartAmount.value })
    coupons.value = res.data.list
  }
  catch {
    coupons.value = []
  }
}

async function openCouponPicker() {
  selectedCouponId.value = 0
  await loadCoupons()
  couponVisible.value = true
}

function pickCoupon(c: Coupon) {
  selectedCouponId.value = c.id
  couponVisible.value = false
}

/* ---------- 本桌已点弹层 ---------- */
const orderedVisible = ref(false)

const currentSku = computed(() => currentGoods.value?.skus.find(s => s.id === currentSkuId.value))
const myCartAmount = computed(() => tavern.myCartAmount)
const myCartCount = computed(() => tavern.myCartCount)
const tableQty = computed(() => tavern.orderedAgg.reduce((s, i) => s + i.qty, 0))

onLoad(async (query: any) => {
  try {
    await tavern.bootstrap()
    // 桌码参数：scene 形如 "t=A01"（真实扫码）或直接 tableNo=A01（开发调试）
    let tableNo = query?.tableNo || ''
    if (!tableNo && query?.scene) {
      const scene = decodeURIComponent(query.scene)
      tableNo = scene.replace(/^t=/, '')
    }
    if (tableNo)
      await tavern.openTable(tableNo)
  }
  catch (e: any) {
    toast.error(e?.message || '初始化失败')
  }
  finally {
    loading.value = false
  }
  await loadMenu()
})

onShow(() => {
  // 轮询兜底：同桌点单后同步（后续替换为 WebSocket 增量推送）
  if (tavern.session)
    tavern.refresh().catch(() => {})
})

async function loadMenu() {
  try {
    const res = await getMenu()
    categories.value = res.data.list
    activeCategory.value = res.data.list[0]?.id || 0
  }
  catch (e: any) {
    toast.error(e?.message || '菜单加载失败')
  }
}

const activeGoods = computed(() => categories.value.find(c => c.id === activeCategory.value)?.goods || [])

/** 开发/演示：手动输入桌号进桌 */
async function manualOpen(tableNo: string) {
  if (!tableNo) {
    toast.warning('请输入桌号')
    return
  }
  try {
    await tavern.openTable(tableNo)
    toast.success(`已进入 ${tableNo}`)
  }
  catch (e: any) {
    toast.error(e?.message || '开台失败')
  }
}

function openSku(goods: Goods) {
  if (goods.isCleared) {
    toast.warning(`${goods.name} 已沽清`)
    return
  }
  currentGoods.value = goods
  currentSkuId.value = goods.skus[0]?.id || 0
  currentQty.value = 1
  currentRemark.value = ''
  skuVisible.value = true
}

async function confirmAdd() {
  if (!currentGoods.value || !currentSku.value)
    return
  try {
    await tavern.addToCart({
      goodsId: currentGoods.value.id,
      skuId: currentSku.value.id,
      qty: currentQty.value,
      remark: currentRemark.value,
    })
    skuVisible.value = false
    toast.success('已加入购物车')
  }
  catch (e: any) {
    toast.error(e?.message || '加入失败')
  }
}

/** 提交我的购物车 → 生成本桌新订单（先付制） */
async function submitMyCart() {
  if (!tavern.myCart.length) {
    toast.warning('购物车是空的')
    return
  }
  try {
    const order = await tavern.submit(tavern.myCart.map(c => c.id), selectedCouponId.value)
    cartVisible.value = false
    router.push({ name: 'orderDetail', query: { orderNo: order.orderNo } })
  }
  catch (e: any) {
    toast.error(e?.message || '下单失败')
  }
}

function skuLabel(goods: Goods) {
  if (goods.skus.length === 1)
    return `¥${goods.skus[0].price}`
  return `¥${Math.min(...goods.skus.map(s => s.price))} 起`
}
</script>

<template>
  <view class="min-h-100vh pb-30 bg-[var(--wot-filled-content)]">
    <!-- 未开台：演示用手动进桌 -->
    <view v-if="!tavern.session" class="p-6">
      <view class="mb-4 rounded-3 p-5 wot-bg-filled-oppo">
        <text class="block text-4 font-bold wot-text-text-main">
          扫码进桌
        </text>
        <text class="mt-2 block text-3 wot-text-text-secondary">
          请扫描桌面上的桌码；开发调试可直接输入桌号
        </text>
        <view class="mt-4 flex items-center gap-2">
          <wd-input v-model="manualTableNo" placeholder="如 A01" custom-class="flex-1" />
          <wd-button type="primary" size="small" @click="manualOpen(manualTableNo)">
            进桌
          </wd-button>
        </view>
      </view>
    </view>

    <template v-else>
      <!-- 桌台信息 + 本桌已点吸顶栏 -->
      <view class="sticky top-0 z-10 flex items-center justify-between px-4 py-3 wot-bg-filled-oppo">
        <view class="flex items-center gap-2">
          <wd-icon name="home" size="16px" />
          <text class="text-4 font-bold wot-text-text-main">
            {{ tavern.tableName }}
          </text>
          <text class="text-2.5 wot-text-text-secondary">
            先付制 · 支付后立即制作
          </text>
        </view>
        <view class="flex items-center gap-1" @click="orderedVisible = true">
          <text class="text-3 wot-text-text-main">
            本桌已点 {{ tableQty }} 件
          </text>
          <wd-icon name="arrow-right" size="14px" />
        </view>
      </view>

      <view class="flex" style="height: calc(100vh - 96px)">
        <!-- 左侧分类 -->
        <scroll-view scroll-y class="w-24 shrink-0 h-full">
          <view
            v-for="c in categories" :key="c.id"
            class="px-3 py-4 text-3"
            :class="c.id === activeCategory ? 'bg-white font-bold wot-text-theme' : 'wot-text-text-secondary'"
            @click="activeCategory = c.id"
          >
            {{ c.name }}
          </view>
        </scroll-view>

        <!-- 右侧商品 -->
        <scroll-view scroll-y class="flex-1 h-full px-3">
          <view v-for="g in activeGoods" :key="g.id" class="mb-3 rounded-3 bg-white p-3">
            <view class="flex items-start justify-between">
              <view class="flex-1">
                <view class="flex items-center gap-1">
                  <text class="text-3.5 font-bold wot-text-text-main">
                    {{ g.name }}
                  </text>
                  <text
                    v-for="t in g.tags" :key="t"
                    class="rounded-1 bg-[var(--wot-color-warning-light)] px-1 text-2 text-[var(--wot-color-warning)]"
                  >
                    {{ t }}
                  </text>
                  <text v-if="g.isCleared" class="text-2 text-[var(--wot-color-danger)]">已沽清</text>
                  <text v-if="g.onlyDineIn" class="text-2 wot-text-text-auxiliary">仅堂食</text>
                </view>
                <text class="mt-1 block text-2.5 wot-text-text-secondary">
                  {{ g.brief }}
                </text>
                <text class="mt-1 block text-2.5 wot-text-text-auxiliary">
                  约 {{ g.makeMinutes }} 分钟出餐
                </text>
              </view>
              <view class="ml-3 flex flex-col items-end">
                <text class="text-4 font-bold text-[var(--wot-color-danger)]">
                  {{ skuLabel(g) }}
                </text>
                <wd-button
                  size="small" type="primary" custom-class="mt-2"
                  :disabled="g.isCleared"
                  @click="openSku(g)"
                >
                  选规格
                </wd-button>
              </view>
            </view>
          </view>
          <view class="h-24" />
        </scroll-view>
      </view>

      <!-- 底部购物车栏 -->
      <view class="fixed bottom-0 left-0 right-0 z-20 flex items-center justify-between px-4 py-3 bg-white shadow">
        <view class="flex items-center gap-3" @click="cartVisible = true">
          <view class="relative">
            <wd-icon name="cart" size="26px" />
            <text
              v-if="myCartCount"
              class="absolute -top-1 -right-2 rounded-full bg-[var(--wot-color-danger)] px-1 text-2 text-white"
            >
              {{ myCartCount }}
            </text>
          </view>
          <view>
            <text class="block text-4 font-bold text-[var(--wot-color-danger)]">
              ¥{{ myCartAmount.toFixed(2) }}
            </text>
            <text class="block text-2.5 wot-text-text-auxiliary">
              我的购物车（全桌可见）
            </text>
          </view>
        </view>
        <wd-button type="primary" :disabled="!myCartCount" @click="submitMyCart">
          去支付
        </wd-button>
      </view>

      <!-- 规格弹层 -->
      <wd-popup v-model="skuVisible" position="bottom" round custom-style="padding: 16px">
        <view v-if="currentGoods">
          <text class="block text-4 font-bold wot-text-text-main">
            {{ currentGoods.name }}
          </text>
          <text class="mt-1 block text-2.5 wot-text-text-secondary">
            {{ currentGoods.brief }}
          </text>

          <text class="mt-4 block text-3 font-bold">规格</text>
          <view class="mt-2 flex flex-wrap gap-2">
            <view
              v-for="s in currentGoods.skus" :key="s.id"
              class="rounded-2 border border-solid px-3 py-2 text-3"
              :class="s.id === currentSkuId ? 'border-[var(--wot-color-theme)] wot-text-theme' : 'border-[var(--wot-color-border)] wot-text-text-secondary'"
              @click="currentSkuId = s.id"
            >
              {{ s.name }} ¥{{ s.price }}
            </view>
          </view>

          <view class="mt-4 flex items-center justify-between">
            <text class="text-3 font-bold">数量</text>
            <view class="flex items-center gap-3">
              <wd-button size="small" @click="currentQty = Math.max(1, currentQty - 1)">
                -
              </wd-button>
              <text class="text-4">{{ currentQty }}</text>
              <wd-button size="small" @click="currentQty++">
                +
              </wd-button>
            </view>
          </view>

          <view class="mt-4">
            <text class="block text-3 font-bold">备注</text>
            <wd-input v-model="currentRemark" placeholder="如：加冰、不要香菜" custom-class="mt-2" />
          </view>

          <wd-button type="primary" block custom-class="mt-5" @click="confirmAdd">
            加入购物车 ¥{{ ((currentSku?.price || 0) * currentQty).toFixed(2) }}
          </wd-button>
        </view>
      </wd-popup>

      <!-- 购物车弹层：我的 / 全桌 -->
      <wd-popup v-model="cartVisible" position="bottom" round custom-style="padding: 16px">
        <view class="flex items-center gap-4">
          <text
            class="text-3.5" :class="cartTab === 'mine' ? 'font-bold wot-text-theme' : 'wot-text-text-secondary'"
            @click="cartTab = 'mine'"
          >
            我的购物车
          </text>
          <text
            class="text-3.5" :class="cartTab === 'table' ? 'font-bold wot-text-theme' : 'wot-text-text-secondary'"
            @click="cartTab = 'table'"
          >
            同桌其他人的
          </text>
        </view>

        <scroll-view scroll-y class="mt-3" style="max-height: 45vh">
          <template v-if="cartTab === 'mine'">
            <view v-if="!tavern.myCart.length" class="py-6 text-center text-3 wot-text-text-auxiliary">
              购物车是空的
            </view>
            <view v-for="c in tavern.myCart" :key="c.id" class="flex items-center justify-between border-b border-solid border-[var(--wot-color-border)] py-3">
              <view class="flex-1">
                <text class="block text-3 wot-text-text-main">
                  {{ c.goodsName }} · {{ c.skuName }}
                </text>
                <text class="block text-2.5 wot-text-text-auxiliary">
                  ¥{{ c.price }} {{ c.remark ? '· ' + c.remark : '' }}
                </text>
              </view>
              <view class="flex items-center gap-2">
                <wd-button size="small" @click="tavern.updateCartQty(c.id, Math.max(1, c.qty - 1))">
                  -
                </wd-button>
                <text>{{ c.qty }}</text>
                <wd-button size="small" @click="tavern.updateCartQty(c.id, c.qty + 1)">
                  +
                </wd-button>
                <wd-button size="small" type="danger" @click="tavern.removeCartItem(c.id)">
                  删
                </wd-button>
              </view>
            </view>
          </template>

          <template v-else>
            <view v-if="!tavern.othersCart.length" class="py-6 text-center text-3 wot-text-text-auxiliary">
              同桌还没有人加购
            </view>
            <view v-for="c in tavern.othersCart" :key="c.id" class="flex items-center justify-between border-b border-solid border-[var(--wot-color-border)] py-3">
              <view class="flex-1">
                <text class="block text-3 wot-text-text-main">
                  {{ c.goodsName }} · {{ c.skuName }}
                </text>
                <text class="block text-2.5 wot-text-text-auxiliary">
                  {{ c.memberName }} 的篮子 · ¥{{ c.price }} × {{ c.qty }}
                </text>
              </view>
              <text class="text-2.5 wot-text-text-auxiliary">仅可查看</text>
            </view>
          </template>
        </scroll-view>

        <view class="mt-3 flex items-center justify-between border-t border-solid border-[var(--wot-color-border)] pt-3" @click="openCouponPicker">
          <text class="text-3">
            优惠券
          </text>
          <text class="text-3" :class="selectedCoupon ? 'text-[var(--wot-color-danger)]' : 'wot-text-text-auxiliary'">
            {{ selectedCoupon ? `${selectedCoupon.name} -¥${discount.toFixed(2)}` : `${coupons.length ? coupons.length + ' 张可用' : '暂无可用'} >` }}
          </text>
        </view>
        <view class="mt-3 flex items-center justify-between">
          <text class="text-4 font-bold text-[var(--wot-color-danger)]">
            应付 ¥{{ payableAmount.toFixed(2) }}
          </text>
          <wd-button type="primary" :disabled="!myCartCount" @click="submitMyCart">
            提交并支付
          </wd-button>
        </view>
      </wd-popup>

      <!-- 优惠券选择 -->
      <wd-popup v-model="couponVisible" position="bottom" round custom-style="padding: 16px">
        <text class="block text-4 font-bold">
          选择优惠券
        </text>
        <text class="mt-1 block text-2.5 wot-text-text-auxiliary">
          当前购物车 ¥{{ myCartAmount.toFixed(2) }}
        </text>
        <scroll-view scroll-y class="mt-3" style="max-height: 45vh">
          <view v-if="!coupons.length" class="py-6 text-center text-3 wot-text-text-auxiliary">
            暂无可用优惠券
          </view>
          <view
            v-for="c in coupons" :key="c.id"
            class="mt-2 flex items-center justify-between rounded-3 p-3"
            :class="c.usable ? 'bg-white' : 'bg-[var(--wot-color-border-light)]'"
            @click="c.usable && pickCoupon(c)"
          >
            <view>
              <text class="block text-3 font-bold" :class="c.usable ? 'wot-text-text-main' : 'wot-text-text-auxiliary'">
                {{ c.name }}
              </text>
              <text class="block text-2.5 wot-text-text-auxiliary">
                满 {{ c.minAmount.toFixed(0) }} 可用 · {{ (c.expireAt || '').slice(0, 10) }} 到期
              </text>
            </view>
            <text class="text-3" :class="c.usable ? 'text-[var(--wot-color-danger)]' : 'wot-text-text-auxiliary'">
              {{ c.usable ? (c.type === 2 ? (c.discount * 10).toFixed(1) + ' 折' : '-¥' + c.amount.toFixed(0)) : c.unusableReason }}
            </text>
          </view>
        </scroll-view>
        <wd-button block custom-class="mt-3" @click="selectedCouponId = 0; couponVisible = false">
          不使用优惠券
        </wd-button>
      </wd-popup>

      <!-- 本桌已点（按菜聚合，可看谁点的） -->
      <wd-popup v-model="orderedVisible" position="bottom" round custom-style="padding: 16px">
        <text class="block text-4 font-bold">
          本桌已点 {{ tableQty }} 件
        </text>
        <text class="mt-1 block text-2.5 wot-text-text-auxiliary">
          全桌可见 · 已付合计 ¥{{ tavern.tablePaidAmount.toFixed(2) }}
        </text>
        <scroll-view scroll-y class="mt-3" style="max-height: 55vh">
          <view v-if="!tavern.orderedAgg.length" class="py-6 text-center text-3 wot-text-text-auxiliary">
            还没有人点东西
          </view>
          <view v-for="i in tavern.orderedAgg" :key="`${i.goodsId}-${i.skuId}`" class="flex items-center justify-between border-b border-solid border-[var(--wot-color-border)] py-3">
            <view class="flex-1">
              <text class="block text-3 wot-text-text-main">
                {{ i.goodsName }} · {{ i.skuName }} × {{ i.qty }}
              </text>
              <text class="block text-2.5 wot-text-text-auxiliary">
                {{ i.byMembers.join('、') }} 点的 · 已上 {{ i.servedQty }}/{{ i.qty }}
              </text>
            </view>
            <text
              class="text-2.5"
              :class="i.servedQty >= i.qty ? 'text-[var(--wot-color-success)]' : 'text-[var(--wot-color-warning)]'"
            >
              {{ i.servedQty >= i.qty ? '已上齐' : '制作中' }}
            </text>
          </view>
        </scroll-view>
      </wd-popup>
    </template>
  </view>
</template>
