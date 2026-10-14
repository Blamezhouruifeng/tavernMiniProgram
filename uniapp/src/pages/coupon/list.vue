<script setup lang="ts">
/** 我的优惠券：新人礼自动发放，过期与门槛状态一目了然 */
import type { Coupon } from '@/types/tavern'
import { getCouponList } from '@/api/tavern'
import { useGlobalToast } from '@/composables/useGlobalToast'

definePage({
  name: 'couponList',
  style: {
    navigationBarTitleText: '我的优惠券',
  },
})

const router = useRouter()
const toast = useGlobalToast()
const tab = ref(1)
const list = ref<Coupon[]>([])
const loading = ref(true)

const tabs = [
  { value: 1, label: '未使用' },
  { value: 2, label: '已使用' },
  { value: 3, label: '已过期' },
]

onLoad(async () => {
  await load()
})

async function load() {
  loading.value = true
  try {
    const res = await getCouponList({ status: tab.value })
    list.value = res.data.list
  }
  catch (e: any) {
    toast.error(e?.message || '加载失败')
  }
  finally {
    loading.value = false
  }
}

async function switchTab(v: number) {
  tab.value = v
  await load()
}

function amountText(c: Coupon) {
  return c.type === 2 ? `${(c.discount * 10).toFixed(1)} 折` : `¥${c.amount.toFixed(0)}`
}

function goOrder() {
  router.push({ name: 'menu' })
}
</script>

<template>
  <view class="min-h-100vh p-4 bg-[var(--wot-filled-content)]">
    <view class="flex items-center gap-4">
      <text
        v-for="t in tabs" :key="t.value"
        class="text-3.5" :class="tab === t.value ? 'font-bold wot-text-theme' : 'wot-text-text-secondary'"
        @click="switchTab(t.value)"
      >
        {{ t.label }}
      </text>
    </view>

    <view v-if="loading" class="py-10 text-center text-3 wot-text-text-auxiliary">
      加载中…
    </view>
    <view v-else-if="!list.length" class="py-10 text-center text-3 wot-text-text-auxiliary">
      暂无优惠券
    </view>

    <view v-for="c in list" :key="c.id" class="mt-3 flex rounded-3 overflow-hidden bg-white">
      <view class="w-28 shrink-0 flex flex-col items-center justify-center bg-[var(--wot-color-theme-light)] py-4">
        <text class="text-4 font-bold wot-text-theme">
          {{ amountText(c) }}
        </text>
        <text class="mt-1 text-2 text-[var(--wot-color-theme)]">
          满 {{ c.minAmount.toFixed(0) }} 可用
        </text>
      </view>
      <view class="flex-1 p-3">
        <text class="block text-3.5 font-bold wot-text-text-main">
          {{ c.name }}
        </text>
        <text class="mt-1 block text-2.5 wot-text-text-auxiliary">
          有效期至 {{ (c.expireAt || '').slice(0, 10) }}
        </text>
        <text v-if="c.status === 1 && !c.usable" class="mt-1 block text-2.5 text-[var(--wot-color-warning)]">
          {{ c.unusableReason || '暂不可用' }}
        </text>
        <view class="mt-2 flex justify-end">
          <wd-button v-if="c.status === 1" size="mini" type="primary" @click="goOrder">
            去使用
          </wd-button>
          <text v-else class="text-2.5 wot-text-text-auxiliary">
            {{ c.status === 2 ? '已使用' : '已过期' }}
          </text>
        </view>
      </view>
    </view>

    <text class="mt-6 block text-center text-2.5 wot-text-text-auxiliary">
      下单时可在购物车中选择使用
    </text>
  </view>
</template>
