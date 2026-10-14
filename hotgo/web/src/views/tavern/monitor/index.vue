<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="酒馆 · 异常订单与对账">
        先付制下「钱已收但流程没走完」的单子都在这里：支付成功未出票可一键补出票，退款待审可在订单页处理。
      </n-card>
    </div>

    <n-card :bordered="false" class="proCard">
      <n-tabs type="line" v-model:value="tab">
        <n-tab-pane name="abnormal" tab="异常订单">
          <n-space justify="end" class="mb-3">
            <n-button size="small" @click="loadAbnormal">
              刷新
            </n-button>
          </n-space>
          <n-data-table :columns="abnormalColumns" :data="abnormalList" :loading="loading" :row-key="(row) => row.id" />
        </n-tab-pane>

        <n-tab-pane name="reconcile" tab="支付对账">
          <n-space align="center" class="mb-3">
            <n-date-picker v-model:value="reconcileDate" type="date" />
            <n-button type="primary" size="small" @click="loadReconcile">
              生成对账
            </n-button>
          </n-space>

          <n-grid :cols="4" :x-gap="12" class="mb-3">
            <n-gi>
              <n-statistic label="已支付订单数" :value="reconcile.orderCount" />
            </n-gi>
            <n-gi>
              <n-statistic label="订单实收合计" :value="'¥' + Number(reconcile.orderAmount || 0).toFixed(2)" />
            </n-gi>
            <n-gi>
              <n-statistic label="支付流水合计" :value="'¥' + Number(reconcile.payRecordAmount || 0).toFixed(2)" />
            </n-gi>
            <n-gi>
              <n-statistic label="差异" :value="'¥' + Number(reconcile.diff || 0).toFixed(2)" :value-style="reconcile.diff ? 'color:#d03050' : ''" />
            </n-gi>
          </n-grid>

          <n-data-table :columns="reconcileColumns" :data="reconcile.items || []" :loading="loading" :row-key="(row) => row.orderNo" />
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { h, ref, onMounted } from 'vue';
  import { NButton, NTag, useMessage } from 'naive-ui';
  import { OrderAbnormal, OrderReissue, Reconcile } from '@/api/tavern';

  const message = useMessage();
  const tab = ref('abnormal');
  const loading = ref(false);
  const abnormalList = ref<any[]>([]);
  const reconcile = ref<any>({ items: [] });
  const reconcileDate = ref<number>(Date.now());

  const abnormalColumns = [
    { title: '订单号', key: 'orderNo', width: 210 },
    { title: '桌号', key: 'tableName', width: 90 },
    { title: '下单人', key: 'memberName', width: 120 },
    { title: '实付', key: 'payAmount', width: 100, render: (row) => '¥' + Number(row.payAmount || 0).toFixed(2) },
    { title: '支付时间', key: 'payTime', width: 180 },
    {
      title: '异常',
      key: 'issue',
      width: 260,
      render(row: any) {
        return h(NTag, { type: 'warning', bordered: false }, { default: () => row.issue });
      },
    },
    {
      title: '操作',
      key: 'action',
      width: 140,
      render(row: any) {
        if (row.payStatus === 1 && row.makingStatus === 0) {
          return h(
            NButton,
            { size: 'small', type: 'primary', onClick: () => reissue(row) },
            { default: () => '补出票' }
          );
        }
        return h('span', { style: 'color:#999' }, row.refundStatus === 1 ? '去订单页审核退款' : '无需处理');
      },
    },
  ];

  const reconcileColumns = [
    { title: '订单号', key: 'orderNo', width: 210 },
    { title: '订单实收', key: 'payAmount', width: 120, render: (row) => '¥' + Number(row.payAmount || 0).toFixed(2) },
    { title: '流水金额', key: 'recordAmt', width: 120, render: (row) => '¥' + Number(row.recordAmt || 0).toFixed(2) },
    {
      title: '是否匹配',
      key: 'hasRecord',
      width: 120,
      render(row: any) {
        return h(NTag, { type: row.hasRecord ? 'success' : 'error', bordered: false }, { default: () => (row.hasRecord ? '有流水' : '缺流水') });
      },
    },
    { title: '渠道', key: 'payChannel', width: 110 },
    { title: '支付时间', key: 'payTime', width: 180 },
  ];

  onMounted(() => {
    loadAbnormal();
    loadReconcile();
  });

  async function loadAbnormal() {
    loading.value = true;
    try {
      const res: any = await OrderAbnormal({});
      abnormalList.value = res.data?.list || [];
    } finally {
      loading.value = false;
    }
  }

  async function reissue(row: any) {
    await OrderReissue({ orderNo: row.orderNo });
    message.success('已补出票');
    loadAbnormal();
  }

  async function loadReconcile() {
    loading.value = true;
    try {
      const d = new Date(reconcileDate.value);
      const p = (n: number) => String(n).padStart(2, '0');
      const date = d.getFullYear() + '-' + p(d.getMonth() + 1) + '-' + p(d.getDate());
      const res: any = await Reconcile({ date });
      reconcile.value = res.data || { items: [] };
    } finally {
      loading.value = false;
    }
  }
</script>
