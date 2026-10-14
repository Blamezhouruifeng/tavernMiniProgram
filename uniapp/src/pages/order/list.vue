<script setup lang="ts">
/** 我的订单 / 本桌订单：先付制下每轮酒水都是独立订单 */
import type { Order } from '@/types/tavern'
import { getOrderList } from '@/api/tavern'
import { useTavernStore } from '@/store/tavernStore'
import { useGlobalToast } from '@/composables/useGlobalToast'

definePage({
  name: 'orderList',
  style: {
    navigationBarTitleText: '我的订单',
  },
})

const router = useRouter()
const toast = useGlobalToast()
const tavern = useTavernStore()

const tab = ref<'table' | 'mine'>('table')
const orders = ref<Order[]>([])
const loading = ref(true)

onLoad(async () => {
  await load()
})

async function load() {
  loading.value = true
  try {
    const res = await getOrderList({ sessionId: tavern.session?.id, mine: tab.value === 'mine' })
    orders.value = res.data.list
  }
  catch (e: any) {
    toast.error(e?.message || '订单加载失败')
  }
  finally {
    loading.value = false
  }
}

async function switchTab(t: 'table' | 'mine') {
  tab.value = t
  await load()
}

const payText = (s: number) => ({ 0: '待支付', 1: '已支付', 2: '部分退款', 3: '已退款' } as Record<number, string>)[s] || ''
const makingText = (s: number) => ({ 0: '', 1: '制作中', 2: '已上齐', 3: '已完成' } as Record<number, string>)[s] || ''

function goDetail(o: Order) {
  router.push({ name: 'orderDetail', query: { orderNo: o.orderNo } })
}

function summary(o: Order) {
  return o.items.map(i => `${i.goodsName}×${i.qty}`).join('、')
}
</script>

<template>
  <view class="min-h-100vh p-4 bg-[var(--wot-filled-content)]">
    <view class="flex items-center gap-4">
      <text
        class="text-3.5" :class="tab === 'table' ? 'font-bold wot-text-theme' : 'wot-text-text-secondary'"
        @click="switchTab('table')"
      >
        本桌订单
      </text>
      <text
        class="text-3.5" :class="tab === 'mine' ? 'font-bold wot-text-theme' : 'wot-text-text-secondary'"
        @click="switchTab('mine')"
      >
        我下的单
      </text>
    </view>

    <view v-if="loading" class="py-10 text-center text-3 wot-text-text-auxiliary">
      加载中…
    </view>
    <view v-else-if="!orders.length" class="py-10 text-center text-3 wot-text-text-auxiliary">
      暂无订单
    </view>

    <view v-for="o in orders" :key="o.id" class="mt-3 rounded-3 bg-white p-4" @click="goDetail(o)">
      <view class="flex items-center justify-between">
        <text class="text-3.5 font-bold wot-text-text-main">
          {{ o.tableName }} · 第 {{ o.orderSeq }} 单
        </text>
        <text class="text-2.5" :class="o.payStatus === 1 ? 'text-[var(--wot-color-success)]' : 'text-[var(--wot-color-warning)]'">
          {{ payText(o.payStatus) }}{{ makingText(o.makingStatus) ? ' · ' + makingText(o.makingStatus) : '' }}
        </text>
      </view>
      <text class="mt-2 block text-2.5 wot-text-text-secondary">
        {{ summary(o) }}
      </text>
      <view class="mt-2 flex items-center justify-between">
        <text class="text-2.5 wot-text-text-auxiliary">
          {{ o.memberName }} · {{ o.createdAt }}
        </text>
        <text class="text-3.5 font-bold text-[var(--wot-color-danger)]">
          ¥{{ o.payAmount.toFixed(2) }}
        </text>
      </view>
    </view>
  </view>
</template>
