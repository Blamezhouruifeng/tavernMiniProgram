<script setup lang="ts">
/**
 * 员工工作台（与客人同一小程序，凭会员 is_staff + 角色进入）
 * 先付制要点：支付成功即出票（无「接单」环节），只能 上齐 / 加急 / 退单（走退款）
 */
import type { Goods, Order } from '@/types/tavern'
import {
  closeTableSession,
  getMenu,
  getStaffTables,
  getStaffTickets,
  staffClearGoods,
  staffDepositRegister,
  staffHandleCall,
  staffMoveTable,
  staffServe,
  staffUrgent,
  staffVerifyBottle,
  switchStaffRole,
  type StaffTable,
} from '@/api/tavern'
import { useTavernStore } from '@/store/tavernStore'
import { useGlobalToast } from '@/composables/useGlobalToast'

definePage({
  name: 'staffBoard',
  style: {
    navigationBarTitleText: '员工工作台',
  },
})

const toast = useGlobalToast()
const tavern = useTavernStore()

const tab = ref<'tables' | 'tickets' | 'verify' | 'clear' | 'deposit'>('tickets')

/* ---------- 转台 ---------- */
const moveVisible = ref(false)
const moveTarget = ref('')
const moveSession = ref<StaffTable | null>(null)

function openMove(t: StaffTable) {
  moveSession.value = t
  moveTarget.value = ''
  moveVisible.value = true
}

async function submitMove() {
  if (!moveTarget.value) {
    toast.warning('请输入目标桌号')
    return
  }
  try {
    await staffMoveTable({ sessionId: moveSession.value!.sessionId, toTableNo: moveTarget.value })
    toast.success('转台成功')
    moveVisible.value = false
    await refreshAll()
  }
  catch (e: any) {
    toast.error(e?.message || '转台失败')
  }
}

/* ---------- 响应呼叫 ---------- */
async function handleCall(t: StaffTable) {
  try {
    await staffHandleCall({ sessionId: t.sessionId })
    toast.success('已响应呼叫')
    await refreshAll()
  }
  catch (e: any) {
    toast.error(e?.message || '操作失败')
  }
}

/* ---------- 代存登记（自带酒水） ---------- */
const depForm = ref({ phone: '', goodsName: '', spec: '', degree: '', remainRatio: 100, slotCode: '' })

async function submitDeposit() {
  if (!depForm.value.phone || !depForm.value.goodsName) {
    toast.warning('请填写手机号与酒品名称')
    return
  }
  try {
    const res = await staffDepositRegister({ ...depForm.value })
    toast.success(`已登记，瓶签 ${res.data.labelCode}，储位 ${res.data.slotCode}`)
    depForm.value = { phone: '', goodsName: '', spec: '', degree: '', remainRatio: 100, slotCode: '' }
  }
  catch (e: any) {
    toast.error(e?.message || '登记失败')
  }
}
const tables = ref<StaffTable[]>([])
const tickets = ref<Order[]>([])
const goods = ref<Goods[]>([])
const loading = ref(false)

const verifyCode = ref('')
const verifyResult = ref<{ deposit: any, memberId: number, takeRatio: number } | null>(null)
const takeRatio = ref(50)

onLoad(async () => {
  if (!tavern.member)
    await tavern.bootstrap().catch(() => {})
  await refreshAll()
})

onShow(() => {
  refreshAll()
})

async function refreshAll() {
  loading.value = true
  try {
    const [t, k, m] = await Promise.all([getStaffTables(), getStaffTickets(), getMenu()])
    tables.value = t.data.list
    tickets.value = k.data.list
    goods.value = m.data.list.flatMap(c => c.goods)
  }
  catch (e: any) {
    toast.error(e?.message || '加载失败')
  }
  finally {
    loading.value = false
  }
}

/** 开通员工身份（演示用；真实环境由后台「会员标记为员工」） */
async function enableStaff() {
  try {
    await switchStaffRole({ isStaff: true, role: 'waiter' })
    await tavern.bootstrap()
    toast.success('已切换到员工身份')
    await refreshAll()
  }
  catch (e: any) {
    toast.error(e?.message || '切换失败')
  }
}

async function serveItem(order: Order, itemId: number) {
  try {
    await staffServe({ orderNo: order.orderNo, itemIds: [itemId] })
    toast.success('已标记上齐')
    await refreshAll()
  }
  catch (e: any) {
    toast.error(e?.message || '操作失败')
  }
}

async function serveAll(order: Order) {
  try {
    await staffServe({ orderNo: order.orderNo })
    toast.success('整单已上齐')
    await refreshAll()
  }
  catch (e: any) {
    toast.error(e?.message || '操作失败')
  }
}

async function urgent(order: Order) {
  try {
    await staffUrgent({ orderNo: order.orderNo })
    toast.warning('已加急')
  }
  catch (e: any) {
    toast.error(e?.message || '操作失败')
  }
}

async function closeTable(t: StaffTable) {
  if (!t.sessionId)
    return
  try {
    await closeTableSession({ sessionId: t.sessionId })
    toast.success(`${t.tableName} 已清台`)
    await refreshAll()
  }
  catch (e: any) {
    toast.error(e?.message || '清台失败')
  }
}

async function verify() {
  if (!verifyCode.value) {
    toast.warning('请输入取酒码')
    return
  }
  try {
    const res = await staffVerifyBottle({ code: verifyCode.value, takeRatio: takeRatio.value })
    verifyResult.value = res.data
    toast.success('核销成功')
  }
  catch (e: any) {
    verifyResult.value = null
    toast.error(e?.message || '核销失败')
  }
}

async function toggleClear(g: Goods) {
  try {
    await staffClearGoods({ goodsId: g.id, isCleared: !g.isCleared })
    toast.success(g.isCleared ? `${g.name} 已恢复` : `${g.name} 已沽清`)
    await refreshAll()
  }
  catch (e: any) {
    toast.error(e?.message || '操作失败')
  }
}

const seatText = (s: number) => ({ 1: '空闲', 2: '用餐中', 3: '待清台' } as Record<number, string>)[s] || ''
const makerText = (s: number) => ({ 0: '待制作', 1: '制作中', 2: '已上齐' } as Record<number, string>)[s] || ''
function minutesFrom(at: string) {
  if (!at)
    return 0
  return Math.max(0, Math.floor((Date.now() - new Date(at.replace(/-/g, '/')).getTime()) / 60000))
}
</script>

<template>
  <view class="min-h-100vh p-4 bg-[var(--wot-filled-content)]">
    <!-- 未开通员工身份 -->
    <view v-if="!tavern.isStaff" class="rounded-3 p-5 wot-bg-filled-oppo">
      <text class="block text-4 font-bold wot-text-text-main">
        员工工作台
      </text>
      <text class="mt-2 block text-2.5 wot-text-text-secondary">
        当前会员未开通员工身份。正式环境下由后台在会员上勾选「是员工」并分配角色。
      </text>
      <wd-button type="primary" block custom-class="mt-4" @click="enableStaff">
        演示：切换为员工身份（服务员）
      </wd-button>
    </view>

    <template v-else>
      <!-- 顶部 tab -->
      <view class="flex items-center justify-between rounded-3 bg-white p-3">
        <text
          v-for="t in [{ k: 'tickets', n: '出票区' }, { k: 'tables', n: '桌台' }, { k: 'verify', n: '核销' }, { k: 'deposit', n: '代存' }, { k: 'clear', n: '沽清' }]"
          :key="t.k"
          class="text-3" :class="tab === t.k ? 'font-bold wot-text-theme' : 'wot-text-text-secondary'"
          @click="tab = t.k as any"
        >
          {{ t.n }}
        </text>
      </view>

      <!-- 出票区：支付即出票 -->
      <template v-if="tab === 'tickets'">
        <view v-if="!tickets.length" class="py-10 text-center text-3 wot-text-text-auxiliary">
          暂无待上齐的订单
        </view>
        <view v-for="o in tickets" :key="o.id" class="mt-3 rounded-3 bg-white p-4">
          <view class="flex items-center justify-between">
            <text class="text-3.5 font-bold wot-text-text-main">
              {{ o.tableName }} · 第 {{ o.orderSeq }} 单
            </text>
            <view class="flex items-center gap-2">
              <text class="text-2.5 wot-text-text-auxiliary">
                {{ o.memberName }}
              </text>
              <wd-button size="mini" @click="urgent(o)">
                加急
              </wd-button>
            </view>
          </view>
          <text class="mt-1 block text-2.5 wot-text-text-auxiliary">
            {{ o.orderNo }} · 已支付 ¥{{ o.payAmount.toFixed(2) }}
          </text>

          <view v-for="it in o.items" :key="it.id" class="mt-3 flex items-center justify-between border-b border-solid border-[var(--wot-color-border)] pb-2">
            <view class="flex-1">
              <text class="block text-3 wot-text-text-main">
                {{ it.goodsName }} · {{ it.skuName }} × {{ it.qty }}
              </text>
              <text v-if="it.remark" class="block text-2.5 text-[var(--wot-color-warning)]">
                {{ it.remark }}
              </text>
            </view>
            <view class="flex items-center gap-2">
              <text class="text-2.5" :class="it.makerStatus >= 2 ? 'text-[var(--wot-color-success)]' : 'text-[var(--wot-color-warning)]'">
                {{ makerText(it.makerStatus) }}
              </text>
              <wd-button v-if="it.makerStatus < 2" size="mini" type="primary" @click="serveItem(o, it.id)">
                已上齐
              </wd-button>
            </view>
          </view>

          <wd-button block size="small" custom-class="mt-3" @click="serveAll(o)">
            整单已上齐
          </wd-button>
        </view>
      </template>

      <!-- 桌台看板 -->
      <template v-else-if="tab === 'tables'">
        <view v-for="t in tables" :key="t.tableId" class="mt-3 rounded-3 bg-white p-4">
          <view v-if="t.called" class="mb-2 flex items-center justify-between rounded-2 bg-[var(--wot-color-warning-light)] p-2">
            <text class="text-2.5 text-[var(--wot-color-warning)]">
              客人呼叫中
            </text>
            <wd-button size="mini" type="warning" @click="handleCall(t)">
              已响应
            </wd-button>
          </view>
          <view class="flex items-center justify-between">
            <text class="text-3.5 font-bold wot-text-text-main">
              {{ t.tableName }}
            </text>
            <text
              class="text-2.5"
              :class="t.seatStatus === 2 ? 'text-[var(--wot-color-success)]' : t.seatStatus === 3 ? 'text-[var(--wot-color-warning)]' : 'wot-text-text-auxiliary'"
            >
              {{ seatText(t.seatStatus) }}
            </text>
          </view>
          <text class="mt-1 block text-2.5 wot-text-text-secondary">
            {{ t.guestCount }} 人 · 已开台 {{ minutesFrom(t.openAt) }} 分钟 · 已付 ¥{{ t.paidAmount.toFixed(2) }}
          </text>
          <text class="mt-1 block text-2.5 wot-text-text-auxiliary">
            订单 {{ t.orderCount }} 笔 · 待支付 {{ t.pendingPayCount }} · 待上齐 {{ t.makingCount }}
          </text>
          <view v-if="t.sessionId" class="mt-2 flex items-center gap-2">
            <wd-button size="small" @click="openMove(t)">
              转台
            </wd-button>
            <wd-button size="small" @click="closeTable(t)">
              清台
            </wd-button>
          </view>
        </view>
      </template>

      <!-- 存酒核销 -->
      <template v-else-if="tab === 'verify'">
        <view class="mt-3 rounded-3 bg-white p-4">
          <text class="block text-3 font-bold wot-text-text-main">
            存酒核销
          </text>
          <text class="mt-1 block text-2.5 wot-text-text-auxiliary">
            需核对「手机号后四位 + 取酒码 + 店员确认」三者
          </text>
          <wd-input v-model="verifyCode" placeholder="输入客人出示的 6 位取酒码" custom-class="mt-3" />
          <view class="mt-3 flex items-center justify-between">
            <text class="text-3">本次取用比例</text>
            <view class="flex items-center gap-3">
              <wd-button size="small" @click="takeRatio = Math.max(5, takeRatio - 5)">
                -
              </wd-button>
              <text>{{ takeRatio }}%</text>
              <wd-button size="small" @click="takeRatio = Math.min(100, takeRatio + 5)">
                +
              </wd-button>
            </view>
          </view>
          <wd-button type="primary" block custom-class="mt-4" @click="verify">
            确认核销
          </wd-button>

          <view v-if="verifyResult" class="mt-4 rounded-3 p-3 wot-bg-filled-content">
            <text class="block text-3 font-bold">
              {{ verifyResult.deposit.goodsName }} {{ verifyResult.deposit.spec }}
            </text>
            <text class="mt-1 block text-2.5 wot-text-text-secondary">
              储位 {{ verifyResult.deposit.slotCode }} · 本次取 {{ verifyResult.takeRatio }}%
            </text>
            <text class="mt-1 block text-2.5 wot-text-text-secondary">
              剩余 {{ verifyResult.deposit.remainRatio }}%
            </text>
          </view>
        </view>
      </template>

      <!-- 代存登记（自带酒水） -->
      <template v-else-if="tab === 'deposit'">
        <view class="mt-3 rounded-3 bg-white p-4">
          <text class="block text-3 font-bold wot-text-text-main">
            代存登记（自带酒水）
          </text>
          <text class="mt-1 block text-2.5 wot-text-text-auxiliary">
            客人需先扫码登录并绑定手机号；登记后贴瓶签入库，记好储位
          </text>
          <wd-input v-model="depForm.phone" type="number" :maxlength="11" placeholder="客人手机号" custom-class="mt-3" />
          <wd-input v-model="depForm.goodsName" placeholder="酒品名称（如 山崎12年）" custom-class="mt-2" />
          <wd-input v-model="depForm.spec" placeholder="规格（如 700ml）" custom-class="mt-2" />
          <wd-input v-model="depForm.degree" placeholder="酒精度（如 43%）" custom-class="mt-2" />
          <wd-input v-model="depForm.slotCode" placeholder="储位号（留空自动分配）" custom-class="mt-2" />
          <view class="mt-3 flex items-center justify-between">
            <text class="text-3">
              存入时剩余比例
            </text>
            <view class="flex items-center gap-3">
              <wd-button size="small" @click="depForm.remainRatio = Math.max(5, depForm.remainRatio - 5)">
                -
              </wd-button>
              <text>{{ depForm.remainRatio }}%</text>
              <wd-button size="small" @click="depForm.remainRatio = Math.min(100, depForm.remainRatio + 5)">
                +
              </wd-button>
            </view>
          </view>
          <wd-button type="primary" block custom-class="mt-4" @click="submitDeposit">
            登记入库
          </wd-button>
        </view>
      </template>

      <!-- 沽清 -->
      <template v-else>
        <view class="mt-3 rounded-3 bg-white">
          <view v-for="g in goods" :key="g.id" class="flex items-center justify-between border-b border-solid border-[var(--wot-color-border)] p-4 last:border-b-0">
            <view>
              <text class="block text-3 wot-text-text-main">
                {{ g.name }}
              </text>
              <text class="block text-2.5 wot-text-text-auxiliary">
                ¥{{ g.price }} · 约 {{ g.makeMinutes }} 分钟
              </text>
            </view>
            <wd-button size="mini" :type="g.isCleared ? 'warning' : 'danger'" @click="toggleClear(g)">
              {{ g.isCleared ? '恢复供应' : '沽清' }}
            </wd-button>
          </view>
        </view>
      </template>
    </template>

    <!-- 转台弹层 -->
    <wd-popup v-model="moveVisible" position="center" round custom-style="padding: 20px">
      <text class="block text-3.5 font-bold wot-text-text-main">
        转台 / 并台
      </text>
      <text class="mt-2 block text-2.5 wot-text-text-secondary">
        {{ moveSession?.tableName }} → 目标桌（订单与存酒归属随会话迁移）
      </text>
      <wd-input v-model="moveTarget" placeholder="输入目标桌号，如 B01" custom-class="mt-3" />
      <wd-button type="primary" block custom-class="mt-4" @click="submitMove">
        确认转台
      </wd-button>
    </wd-popup>
  </view>
</template>
