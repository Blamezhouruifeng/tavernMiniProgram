<script setup lang="ts">
/** 订单进度页：先付制 —— 支付成功即为「制作中」，行级展示每道菜的状态 */
import type { Order, OrderItem } from '@/types/tavern'
import { callWaiter, depositFromOrder, getOrderDetail, payOrder, refundApply } from '@/api/tavern'
import { useTavernStore } from '@/store/tavernStore'
import { useGlobalToast } from '@/composables/useGlobalToast'

definePage({
  name: 'orderDetail',
  style: {
    navigationBarTitleText: '订单详情',
  },
})

const router = useRouter()
const route = useRoute()
const toast = useGlobalToast()
const tavern = useTavernStore()

const order = ref<Order | null>(null)
const leftSeconds = ref(0)
let timer: any = null

const orderNo = computed(() => (route.query?.orderNo as string) || tavern.pendingOrder?.orderNo || '')

onLoad(async () => {
  await load()
  startCountdown()
})

onUnload(() => {
  if (timer)
    clearInterval(timer)
})

async function load() {
  if (!orderNo.value)
    return
  try {
    const res = await getOrderDetail({ orderNo: orderNo.value })
    order.value = res.data
  }
  catch (e: any) {
    toast.error(e?.message || '订单加载失败')
  }
}

function startCountdown() {
  timer = setInterval(async () => {
    // 未支付：倒计时并轮询状态（支付后进入制作）
    if (order.value?.payStatus === 0) {
      const expire = new Date((order.value.expireAt || '').replace(/-/g, '/')).getTime()
      leftSeconds.value = Math.max(0, Math.floor((expire - Date.now()) / 1000))
      if (leftSeconds.value <= 0) {
        toast.warning('订单已超时取消')
        await load()
        return
      }
    }
    await load()
  }, 3000)
}

async function pay() {
  if (!order.value)
    return
  try {
    await payOrder({ orderNo: order.value.orderNo })
    toast.success('支付成功，吧台已开始制作')
    await load()
    await tavern.refresh().catch(() => {})
  }
  catch (e: any) {
    toast.error(e?.message || '支付失败')
  }
}

/** 退酒入存：整瓶酒不喝完直接转存 */
async function toDeposit(item: OrderItem) {
  try {
    await depositFromOrder({ orderId: order.value!.id, itemIds: [item.id] })
    toast.success('已转入存酒，可在「我的存酒」查看')
  }
  catch (e: any) {
    toast.error(e?.message || '转存失败')
  }
}

/* ---------- 退款申请 ---------- */
const refundVisible = ref(false)
const refundReason = ref('')
const refundItemIds = ref<number[]>([])

function openRefund(item?: OrderItem) {
  refundItemIds.value = item ? [item.id] : []
  refundReason.value = ''
  refundVisible.value = true
}

async function submitRefund() {
  if (!refundReason.value) {
    toast.warning('请填写退款原因')
    return
  }
  try {
    const res = await refundApply({ orderNo: orderNo.value, itemIds: refundItemIds.value, reason: refundReason.value })
    toast.success(res.data.refundStatus === 4 ? '已退款（原路退回）' : '已提交，等待店员审核')
    refundVisible.value = false
    await load()
  }
  catch (e: any) {
    toast.error(e?.message || '申请失败')
  }
}

/* ---------- 呼叫服务员 / 催单 ---------- */
async function call(type: 'waiter' | 'urge') {
  try {
    await callWaiter({ sessionId: order.value!.sessionId, type, orderNo: type === 'urge' ? orderNo.value : undefined })
    toast.success(type === 'urge' ? '已催单，请稍候' : '已呼叫服务员')
  }
  catch (e: any) {
    toast.error(e?.message || '操作失败')
  }
}

const refundText = (s: number) => ({ 0: '', 1: '退款审核中', 2: '已通过', 3: '已驳回', 4: '已退款' } as Record<number, string>)[s] || ''

const makerText = (s: number) => ({ 0: '待制作', 1: '制作中', 2: '已上齐' } as Record<number, string>)[s] || '待制作'
const payText = (s: number) => ({ 0: '待支付', 1: '已支付', 2: '部分退款', 3: '已退款' } as Record<number, string>)[s] || ''
</script>

<template>
  <view class="min-h-100vh p-4 bg-[var(--wot-filled-content)]">
    <view v-if="!order" class="py-10 text-center text-3 wot-text-text-auxiliary">
      加载中…
    </view>

    <template v-else>
      <!-- 顶部状态 -->
      <view class="rounded-3 p-4 wot-bg-filled-oppo">
        <view class="flex items-center justify-between">
          <text class="text-4 font-bold wot-text-text-main">
            {{ order.tableName }} · 第 {{ order.orderSeq }} 单
          </text>
          <text
            class="text-3"
            :class="order.payStatus === 1 ? 'text-[var(--wot-color-success)]' : 'text-[var(--wot-color-warning)]'"
          >
            {{ payText(order.payStatus) }}
          </text>
        </view>
        <text class="mt-2 block text-2.5 wot-text-text-auxiliary">
          单号 {{ order.orderNo }} · 下单人 {{ order.memberName }}
        </text>
        <text class="mt-2 block text-2.5 wot-text-text-auxiliary">
          {{ order.payStatus === 1 ? '支付成功 = 已接单，吧台正在制作' : '请尽快支付，超时订单将自动取消' }}
        </text>
      </view>

      <!-- 明细（行级状态） -->
      <view class="mt-3 rounded-3 bg-white p-4">
        <text class="block text-3.5 font-bold wot-text-text-main">
          明细
        </text>
        <view v-for="it in order.items" :key="it.id" class="mt-3 flex items-center justify-between border-b border-solid border-[var(--wot-color-border)] pb-3">
          <view class="flex-1">
            <text class="block text-3 wot-text-text-main">
              {{ it.goodsName }} · {{ it.skuName }} × {{ it.qty }}
            </text>
            <text v-if="it.remark" class="block text-2.5 wot-text-text-auxiliary">
              备注：{{ it.remark }}
            </text>
            <text class="block text-2.5 wot-text-text-auxiliary">
              小计 ¥{{ it.amount.toFixed(2) }}
            </text>
          </view>
          <view class="flex flex-col items-end">
            <text
              class="text-2.5"
              :class="it.makerStatus >= 2 ? 'text-[var(--wot-color-success)]' : 'text-[var(--wot-color-warning)]'"
            >
              {{ makerText(it.makerStatus) }}
            </text>
            <wd-button v-if="it.isWine" size="small" custom-class="mt-1" @click="toDeposit(it)">
              退酒入存
            </wd-button>
            <wd-button
              v-if="order.payStatus === 1 && order.refundStatus === 0"
              size="mini" custom-class="mt-1" @click="openRefund(it)"
            >
              申请退款
            </wd-button>
          </view>
        </view>
        <view class="mt-3 flex items-center justify-between">
          <text class="text-3 wot-text-text-secondary">
            合计
          </text>
          <text class="text-4 font-bold text-[var(--wot-color-danger)]">
            ¥{{ order.payAmount.toFixed(2) }}
          </text>
        </view>
      </view>

      <!-- 操作 -->
      <view class="mt-4">
        <wd-button v-if="order.payStatus === 0" type="primary" block @click="pay">
          立即支付 ¥{{ order.payAmount.toFixed(2) }}（剩余 {{ leftSeconds }}s）
        </wd-button>
        <template v-else>
          <view v-if="order.refundStatus" class="mb-3 rounded-2 bg-[var(--wot-color-warning-light)] p-2">
            <text class="text-2.5 text-[var(--wot-color-warning)]">
              {{ refundText(order.refundStatus) }} · 退款金额 ¥{{ (order.refundAmount || 0).toFixed(2) }}
            </text>
          </view>
          <view class="flex items-center gap-3">
            <wd-button block @click="call('waiter')">
              呼叫服务员
            </wd-button>
            <wd-button block @click="call('urge')">
              催单
            </wd-button>
          </view>
          <wd-button block custom-class="mt-3" @click="router.push({ name: 'menu' })">
            继续加菜
          </wd-button>
          <view class="mt-3 flex items-center gap-3">
            <wd-button block @click="openRefund()">
              整单退款
            </wd-button>
            <wd-button block @click="router.pushTab({ name: 'bottleList' })">
              我的存酒
            </wd-button>
          </view>
        </template>

        <!-- 退款申请弹层 -->
        <wd-popup v-model="refundVisible" position="center" round custom-style="padding: 20px">
          <text class="block text-3.5 font-bold wot-text-text-main">
            申请退款
          </text>
          <text class="mt-2 block text-2.5 wot-text-text-secondary">
            {{ refundItemIds.length ? '按所选明细行退款' : '整单退款' }}；制作前可自助通过，制作后需店员审核。
          </text>
          <wd-input v-model="refundReason" placeholder="请填写退款原因" custom-class="mt-3" />
          <wd-button type="primary" block custom-class="mt-4" @click="submitRefund">
            提交申请
          </wd-button>
        </wd-popup>
      </view>
    </template>
  </view>
</template>
