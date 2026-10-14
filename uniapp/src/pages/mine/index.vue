<script setup lang="ts">
/** 会员中心：客人与员工共用同一套会员体系，员工额外显示工作台入口 */
import { bindPhone } from '@/api/tavern'
import { useTavernStore } from '@/store/tavernStore'
import { useGlobalToast } from '@/composables/useGlobalToast'

definePage({
  name: 'mine',
  layout: 'tabbar',
  style: {
    navigationBarTitleText: '我的',
  },
})

const router = useRouter()
const toast = useGlobalToast()
const tavern = useTavernStore()

onLoad(async () => {
  if (!tavern.member) {
    try {
      await tavern.bootstrap()
    }
    catch (e: any) {
      toast.error(e?.message || '加载失败')
    }
  }
})

const entries = [
  { icon: 'order', title: '我的订单', desc: '本桌与他人订单', name: 'orderList' },
  { icon: 'gift', title: '我的存酒', desc: '剩余量 · 取酒码', name: 'bottleList' },
  { icon: 'coupon', title: '优惠券', desc: 'P1 功能占位', name: '' },
  { icon: 'setting', title: '设置', desc: '手机号绑定 / 注销', name: '' },
]

function go(name: string) {
  if (!name) {
    toast.warning('该功能在 M2 版本开放')
    return
  }
  if (name === 'bottleList') {
    router.pushTab({ name: 'bottleList' })
    return
  }
  router.push({ name })
}

function goWorkbench() {
  router.push({ name: 'staffBoard' })
}

/* ---------- 手机号绑定（存酒取酒凭证的前置条件） ---------- */
const phoneVisible = ref(false)
const phoneInput = ref('')

async function submitPhone() {
  if (!/^1\d{10}$/.test(phoneInput.value)) {
    toast.warning('请输入正确的手机号')
    return
  }
  try {
    await bindPhone({ phone: phoneInput.value })
    await tavern.bootstrap()
    phoneVisible.value = false
    toast.success('手机号绑定成功')
  }
  catch (e: any) {
    toast.error(e?.message || '绑定失败')
  }
}
</script>

<template>
  <view class="min-h-100vh p-4 bg-[var(--wot-filled-content)]">
    <!-- 会员卡片 -->
    <view class="rounded-3 p-5 wot-bg-filled-oppo">
      <view class="flex items-center gap-3">
        <view class="h-14 w-14 flex items-center justify-center rounded-full bg-[var(--wot-color-theme-light)]">
          <wd-icon name="user" size="28px" custom-class="wot-text-theme" />
        </view>
        <view class="flex-1">
          <text class="block text-4 font-bold wot-text-text-main">
            {{ tavern.member?.nickname || '未登录' }}
          </text>
          <view class="flex items-center gap-2" @click="!tavern.member?.phone && (phoneVisible = true)">
            <text class="block text-2.5 wot-text-text-secondary">
              {{ tavern.member?.phone ? tavern.member.phone.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2') : '未绑定手机号（存酒需绑定）' }}
            </text>
            <text v-if="!tavern.member?.phone" class="text-2.5 wot-text-theme">
              去绑定
            </text>
          </view>
          <text v-if="tavern.isStaff" class="mt-1 inline-block rounded-1 bg-[var(--wot-color-warning-light)] px-1 text-2 text-[var(--wot-color-warning)]">
            员工 · {{ tavern.member?.staffRole }}
          </text>
        </view>
      </view>

      <view class="mt-4 flex items-center justify-around">
        <view class="text-center">
          <text class="block text-4 font-bold wot-text-text-main">
            {{ tavern.member?.points || 0 }}
          </text>
          <text class="block text-2.5 wot-text-text-auxiliary">
            积分
          </text>
        </view>
        <view class="text-center">
          <text class="block text-4 font-bold wot-text-text-main">
            ¥{{ (tavern.member?.balance || 0).toFixed(2) }}
          </text>
          <text class="block text-2.5 wot-text-text-auxiliary">
            储值余额
          </text>
        </view>
        <view class="text-center">
          <text class="block text-4 font-bold wot-text-text-main">
            ¥{{ tavern.tablePaidAmount.toFixed(2) }}
          </text>
          <text class="block text-2.5 wot-text-text-auxiliary">
            本桌已付
          </text>
        </view>
      </view>
    </view>

    <!-- 员工工作台入口（仅 is_staff 可见） -->
    <view v-if="tavern.isStaff" class="mt-3 rounded-3 bg-white p-4" @click="goWorkbench">
      <view class="flex items-center justify-between">
        <view>
          <text class="block text-3.5 font-bold wot-text-text-main">
            员工工作台
          </text>
          <text class="block text-2.5 wot-text-text-auxiliary">
            桌台看板 · 出票区 · 存酒核销
          </text>
        </view>
        <wd-icon name="arrow-right" size="16px" />
      </view>
    </view>

    <!-- 功能入口 -->
    <view class="mt-3 rounded-3 bg-white">
      <view
        v-for="e in entries" :key="e.title"
        class="flex items-center justify-between border-b border-solid border-[var(--wot-color-border)] p-4 last:border-b-0"
        @click="go(e.name)"
      >
        <view class="flex items-center gap-3">
          <wd-icon :name="e.icon" size="20px" />
          <view>
            <text class="block text-3 wot-text-text-main">
              {{ e.title }}
            </text>
            <text class="block text-2.5 wot-text-text-auxiliary">
              {{ e.desc }}
            </text>
          </view>
        </view>
        <wd-icon name="arrow-right" size="16px" />
      </view>
    </view>

    <text class="mt-6 block text-center text-2.5 wot-text-text-auxiliary">
      先付制 · 支付成功即开始制作
    </text>

    <!-- 绑定手机号弹层 -->
    <wd-popup v-model="phoneVisible" position="center" round custom-style="padding: 20px">
      <text class="block text-3.5 font-bold wot-text-text-main">
        绑定手机号
      </text>
      <text class="mt-2 block text-2.5 wot-text-text-secondary">
        手机号是存酒与取酒的凭证（到店需核对后四位），请填写常用号码。
      </text>
      <wd-input v-model="phoneInput" type="number" :maxlength="11" placeholder="请输入 11 位手机号" custom-class="mt-3" />
      <wd-button type="primary" block custom-class="mt-4" @click="submitPhone">
        确认绑定
      </wd-button>
    </wd-popup>
  </view>
</template>
