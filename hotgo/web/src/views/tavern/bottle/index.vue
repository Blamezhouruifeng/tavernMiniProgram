<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="酒馆 · 存酒台账">
        实物与台账必须一致：储位号对应酒柜位置，瓶签码用于到店扫码认领，到期前 30/7/1 天提醒客人。
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard">
      <BasicForm @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable" />
      <BasicTable
        ref="actionRef"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        :action-column="actionColumn"
        :scroll-x="scrollX"
      />
      <n-modal v-model:show="showEdit" preset="card" title="存酒维护" style="width: 560px">
        <n-form :model="editForm" label-placement="left" label-width="120">
          <n-form-item label="客人">
            <n-input :value="editForm.memberName + '（尾号 ' + editForm.phoneTail + '）'" disabled />
          </n-form-item>
          <n-form-item label="酒品">
            <n-input :value="editForm.goodsName + ' ' + editForm.spec" disabled />
          </n-form-item>
          <n-form-item label="剩余比例">
            <n-slider v-model:value="editForm.remainRatio" :min="0" :max="100" :step="5" />
          </n-form-item>
          <n-form-item label="储位号">
            <n-input v-model:value="editForm.slotCode" placeholder="如 A-2-3" />
          </n-form-item>
          <n-form-item label="状态">
            <n-select v-model:value="editForm.status" :options="statusOptions" />
          </n-form-item>
          <n-form-item label="到期时间">
            <n-date-picker v-model:value="expireTs" type="datetime" clearable />
          </n-form-item>
          <n-form-item label="备注">
            <n-input v-model:value="editForm.remark" type="textarea" placeholder="如 瓶身破损报损原因" />
          </n-form-item>
        </n-form>
        <template #footer>
          <n-space justify="end">
            <n-button @click="showEdit = false">
              取消
            </n-button>
            <n-button type="primary" :loading="saving" @click="submitEdit">
              保存
            </n-button>
          </n-space>
        </template>
      </n-modal>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { adaTableScrollX } from '@/utils/hotgo';
  import { BottleList, BottleUpdate } from '@/api/tavern';

  const message = useMessage();
  const actionRef = ref();
  const showEdit = ref(false);
  const saving = ref(false);
  const scrollX = adaTableScrollX();
  const expireTs = ref<number | null>(null);

  const statusOptions = [
    { label: '在库', value: 1 },
    { label: '已取完', value: 2 },
    { label: '已过期', value: 3 },
    { label: '已转让', value: 4 },
    { label: '报损', value: 5 },
  ];
  const sourceText = (s: number) => ({ 1: '自带代存', 2: '购酒寄存', 3: '退酒入存' } as Record<number, string>)[s] || '-';
  const statusText = (s: number) => ({ 1: '在库', 2: '已取完', 3: '已过期', 4: '已转让', 5: '报损' } as Record<number, string>)[s] || '-';

  const schemas = [
    { field: 'keyword', component: 'NInput', label: '关键词', componentProps: { placeholder: '会员名 / 手机尾号 / 酒品 / 瓶签码' } },
    { field: 'status', component: 'NSelect', label: '状态', componentProps: { placeholder: '全部', options: statusOptions } },
    {
      field: 'sourceType',
      component: 'NSelect',
      label: '来源',
      componentProps: {
        placeholder: '全部',
        options: [
          { label: '自带代存', value: 1 },
          { label: '购酒寄存', value: 2 },
          { label: '退酒入存', value: 3 },
        ],
      },
    },
  ];
  const [register] = useForm({ labelWidth: 120, schemas, showActionButtonGroup: true, showResetButton: true, submitButtonText: '搜索', resetButtonText: '重置' });

  const columns = [
    { title: 'ID', key: 'id', width: 70 },
    { title: '存酒单号', key: 'depositNo', width: 190 },
    { title: '瓶签码', key: 'labelCode', width: 180 },
    { title: '客人', key: 'memberName', width: 120 },
    { title: '手机尾号', key: 'phoneTail', width: 100 },
    { title: '酒品', key: 'goodsName', width: 200 },
    { title: '规格', key: 'spec', width: 110 },
    { title: '剩余', key: 'remainRatio', width: 100, render: (row) => row.remainRatio + '%' },
    { title: '储位', key: 'slotCode', width: 100 },
    { title: '来源', key: 'sourceType', width: 110, render: (row) => sourceText(row.sourceType) },
    { title: '状态', key: 'status', width: 100, render: (row) => statusText(row.status) },
    { title: '存入时间', key: 'depositAt', width: 180 },
    { title: '到期时间', key: 'expireAt', width: 180 },
  ];

  const actionColumn = reactive({
    width: 160,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record: any) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          { label: '维护', onClick: () => openEdit(record) },
          {
            label: '报损',
            type: 'error',
            onClick: () => quickUpdate(record, { status: 5, remark: '报损（瓶身破损/丢失）' }),
          },
        ],
      });
    },
  });

  const editForm = ref<any>({ id: 0, memberName: '', phoneTail: '', goodsName: '', spec: '', remainRatio: 100, slotCode: '', status: 1, remark: '' });

  async function loadDataTable(params: any) {
    const res: any = await BottleList(params);
    return { data: res.data?.list || [], total: res.data?.totalCount || 0, pageCount: res.data?.pageCount || 0 };
  }

  function reloadTable() {
    actionRef.value?.reload();
  }

  function openEdit(record: any) {
    editForm.value = { ...record };
    expireTs.value = record.expireAt ? new Date(record.expireAt.replace(/-/g, '/')).getTime() : null;
    showEdit.value = true;
  }

  async function submitEdit() {
    saving.value = true;
    try {
      const payload: any = { ...editForm.value };
      if (expireTs.value) {
        const d = new Date(expireTs.value);
        const p = (n: number) => String(n).padStart(2, '0');
        payload.expireAt = d.getFullYear() + '-' + p(d.getMonth() + 1) + '-' + p(d.getDate()) + ' ' + p(d.getHours()) + ':' + p(d.getMinutes()) + ':' + p(d.getSeconds());
      }
      await BottleUpdate(payload);
      message.success('保存成功');
      showEdit.value = false;
      reloadTable();
    } finally {
      saving.value = false;
    }
  }

  async function quickUpdate(record: any, patch: any) {
    await BottleUpdate({ id: record.id, status: record.status, remainRatio: record.remainRatio, ...patch });
    message.success('操作成功');
    reloadTable();
  }
</script>
