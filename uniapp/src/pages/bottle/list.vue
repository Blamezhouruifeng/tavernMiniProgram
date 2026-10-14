<script setup lang="ts">
/** 我的存酒：剩余量、储位、到期提醒、动态取酒码（到店由店员扫码核销） */
import type { BottleDeposit } from '@/types/tavern'
import { createTakeCode, getBottleList } from '@/api/tavern'
import { useGlobalToast } from '@/composables/useGlobalToast'

definePage({
  name: 'bottleList',
  layout: 'tabbar',
  style: {
    navigationBarTitleText: '我的存酒',
  },
})

const toast = useGlobalToast()
const list = ref<BottleDeposit[]>([])
const loading = ref(true)

const takeVisible = ref(false)
const takeCode = ref('')
const takeLeft = ref(0)
const takeDeposit = ref<BottleDeposit | null>(null)
let timer: any = null

const statusText = (s: number) => ({ 1: '在库', 2: '已取完', 3: '已过期', 4: '已转让', 5: '报损' } as Record<number, string>)[s] || ''
const sourceText = (s: number) => ({ 1: '自带代存', 2: '购酒寄存', 3: '退酒入存' } as Record<number, string>)[s] || ''

onLoad(async () => {
  await load()
})

onUnload(() => {
  if (timer)
    clearInterval(timer)
})

async function load() {
  try {
    const res = await getBottleList()
    list.value = res.data.list
  }
  catch (e: any) {
    toast.error(e?.message || '存酒加载失败')
  }
  finally {
    loading.value = false
  }
}

async function openTakeCode(dep: BottleDeposit) {
  takeDeposit.value = dep
  try {
    const res = await createTakeCode({ depositIds: [dep.id] })
    takeCode.value = res.data.code
    takeLeft.value = res.data.expireSeconds
    takeVisible.value = true
    if (timer)
      clearInterval(timer)
    timer = setInterval(async () => {
      takeLeft.value -= 1
      if (takeLeft.value <= 0) {
        clearInterval(timer)
        takeVisible.value = false
        // 动态码 30 秒有效；过期自动刷新一版，防止截图冒领
        const r = await createTakeCode({ depositIds: [dep.id] })
        takeCode.value = r.data.code
        takeLeft.value = r.data.expireSeconds
        takeVisible.value = true
        timer = setInterval(() => { takeLeft.value -= 1 }, 1000)
      }
    }, 1000)
  }
  catch (e: any) {
    toast.error(e?.message || '生成取酒码失败')
  }
}

function daysLeft(dep: BottleDeposit) {
  const t = new Date((dep.expireAt || '').replace(/-/g, '/')).getTime()
  return Math.max(0, Math.ceil((t - Date.now()) / 86400000))
}
</script>

<template>
  <view class="min-h-100vh p-4 bg-[var(--wot-filled-content)]">
    <view class="rounded-3 p-4 wot-bg-filled-oppo">
      <text class="block text-4 font-bold wot-text-text-main">
        我的存酒
      </text>
      <text class="mt-1 block text-2.5 wot-text-text-secondary">
        共 {{ list.length }} 条 · 到店出示取酒码由店员核销
      </text>
    </view>

    <view v-if="loading" class="py-10 text-center text-3 wot-text-text-auxiliary">
      加载中…
    </view>

    <view v-else-if="!list.length" class="py-10 text-center">
      <text class="block text-3 wot-text-text-auxiliary">
        还没有存酒记录
      </text>
      <text class="mt-2 block text-2.5 wot-text-text-auxiliary">
        买整瓶喝不完可在订单里点「退酒入存」，自带酒水可由店员代存
      </text>
    </view>

    <view v-for="dep in list" :key="dep.id" class="mt-3 rounded-3 bg-white p-4">
      <view class="flex items-start justify-between">
        <view class="flex-1">
          <text class="block text-3.5 font-bold wot-text-text-main">
            {{ dep.goodsName }}
          </text>
          <text class="mt-1 block text-2.5 wot-text-text-secondary">
            {{ dep.spec }} {{ dep.degree }} · {{ sourceText(dep.sourceType) }}
          </text>
        </view>
        <text
          class="text-2.5"
          :class="dep.status === 1 ? 'text-[var(--wot-color-success)]' : 'wot-text-text-auxiliary'"
        >
          {{ statusText(dep.status) }}
        </text>
      </view>

      <!-- 剩余量 -->
      <view class="mt-3">
        <view class="flex items-center justify-between text-2.5 wot-text-text-auxiliary">
          <text>剩余 {{ dep.remainRatio }}%</text>
          <text>储位 {{ dep.slotCode }}</text>
        </view>
        <view class="mt-1 h-2 w-full overflow-hidden rounded-full bg-[var(--wot-color-border)]">
          <view
            class="h-full rounded-full bg-[var(--wot-color-theme)]"
            :style="{ width: dep.remainRatio + '%' }"
          />
        </view>
      </view>

      <view class="mt-3 flex items-center justify-between">
        <text class="text-2.5" :class="daysLeft(dep) <= 7 ? 'text-[var(--wot-color-danger)]' : 'wot-text-text-auxiliary'">
          {{ daysLeft(dep) }} 天后到期（{{ dep.expireAt.slice(0, 10) }}）
        </text>
        <wd-button size="small" type="primary" :disabled="dep.status !== 1" @click="openTakeCode(dep)">
          出示取酒码
        </wd-button>
      </view>

      <text class="mt-2 block text-2.5 wot-text-text-auxiliary">
        瓶签 {{ dep.labelCode }} · {{ dep.phoneTail }}（手机号后四位）
      </text>
    </view>

    <!-- 取酒码弹层 -->
    <wd-popup v-model="takeVisible" position="center" round custom-style="padding: 24px; text-align: center">
      <text class="block text-3.5 font-bold wot-text-text-main">
        取酒码
      </text>
      <text class="mt-1 block text-2.5 wot-text-text-secondary">
        {{ takeDeposit?.goodsName }} · 剩余 {{ takeDeposit?.remainRatio }}%
      </text>
      <text class="mt-4 block text-8 font-bold tracking-widest wot-text-theme">
        {{ takeCode }}
      </text>
      <text class="mt-2 block text-2.5 text-[var(--wot-color-warning)]">
        {{ takeLeft }} 秒后自动刷新，请勿截图外传
      </text>
      <text class="mt-3 block text-2.5 wot-text-text-auxiliary">
        核销需「手机号 + 取酒码 + 店员确认」，三者缺一不可
      </text>
    </wd-popup>
  </view>
</template>
