<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="酒馆 · 订单查询">
        先付制：支付状态为「已支付」的订单才会进入制作；退款状态反映申请与审核进度。
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard">
      <BasicForm @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable" />
      <BasicTable
        ref="actionRef"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        :scroll-x="scrollX"
      />
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { BasicTable } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { adaTableScrollX } from '@/utils/hotgo';
  import { defRangeShortcuts } from '@/utils/dateUtil';
  import { OrderList } from '@/api/tavern';

  const actionRef = ref();
  const scrollX = adaTableScrollX();

  const schemas = [
    { field: 'orderNo', component: 'NInput', label: '订单号', componentProps: { placeholder: '支持模糊搜索' } },
    { field: 'tableName', component: 'NInput', label: '桌号', componentProps: { placeholder: '如 A01' } },
    {
      field: 'payStatus',
      component: 'NSelect',
      label: '支付状态',
      componentProps: {
        placeholder: '全部',
        options: [
          { label: '待支付', value: 0 },
          { label: '已支付', value: 1 },
          { label: '部分退款', value: 2 },
          { label: '已退款', value: 3 },
        ],
      },
    },
    { field: 'dateRange', component: 'NDatePicker', label: '下单时间', componentProps: { type: 'datetimerange', clearable: true, shortcuts: defRangeShortcuts() } },
  ];
  const [register] = useForm({ labelWidth: 120, schemas, showActionButtonGroup: true, showResetButton: true, submitButtonText: '搜索', resetButtonText: '重置' });

  const payText = (s: number) => ({ 0: '待支付', 1: '已支付', 2: '部分退款', 3: '已退款' } as Record<number, string>)[s] || '-';
  const makingText = (s: number) => ({ 0: '未开始', 1: '制作中', 2: '已上齐', 3: '已完成' } as Record<number, string>)[s] || '-';
  const refundText = (s: number) => ({ 0: '-', 1: '审核中', 2: '已通过', 3: '已驳回', 4: '已退款' } as Record<number, string>)[s] || '-';

  const columns = [
    { title: 'ID', key: 'id', width: 70 },
    { title: '订单号', key: 'orderNo', width: 210 },
    { title: '桌号', key: 'tableName', width: 90 },
    { title: '第几单', key: 'orderSeq', width: 90 },
    { title: '下单人', key: 'memberName', width: 120 },
    { title: '支付状态', key: 'payStatus', width: 110, render: (row) => payText(row.payStatus) },
    { title: '制作状态', key: 'makingStatus', width: 110, render: (row) => makingText(row.makingStatus) },
    { title: '退款', key: 'refundStatus', width: 110, render: (row) => refundText(row.refundStatus) },
    { title: '实付', key: 'payAmount', width: 100, render: (row) => '¥' + Number(row.payAmount || 0).toFixed(2) },
    { title: '退款金额', key: 'refundAmount', width: 110, render: (row) => '¥' + Number(row.refundAmount || 0).toFixed(2) },
    { title: '支付渠道', key: 'payChannel', width: 110 },
    { title: '支付时间', key: 'payTime', width: 180 },
    { title: '下单时间', key: 'createdAt', width: 180 },
  ];

  async function loadDataTable(params: any) {
    const res: any = await OrderList(params);
    return { data: res.data?.list || [], total: res.data?.totalCount || 0, pageCount: res.data?.pageCount || 0 };
  }

  function reloadTable() {
    actionRef.value?.reload();
  }
</script>
