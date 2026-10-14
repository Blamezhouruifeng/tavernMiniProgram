<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="酒馆 · 桌位管理">
        桌号即桌码参数（scene = t=桌号）。保存后可在小程序端「扫码进桌」按桌号进入同一会话。
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
      >
        <template #tableTitle>
          <n-button type="primary" class="min-left-space" @click="openEdit(null)">
            <template #icon>
              <n-icon><PlusOutlined /></n-icon>
            </template>
            新增桌位
          </n-button>
        </template>
      </BasicTable>

      <n-modal v-model:show="showEdit" preset="card" :title="editForm.id ? '编辑桌位' : '新增桌位'" style="width: 560px">
        <n-form ref="formRef" :model="editForm" :rules="rules" label-placement="left" label-width="120">
          <n-form-item label="桌号" path="name">
            <n-input v-model:value="editForm.name" placeholder="如 A01" />
          </n-form-item>
          <n-form-item label="所属区域" path="areaId">
            <n-select v-model:value="editForm.areaId" :options="areaOptions" placeholder="选择区域" />
          </n-form-item>
          <n-form-item label="可容纳人数" path="capacity">
            <n-input-number v-model:value="editForm.capacity" :min="1" :max="30" />
          </n-form-item>
          <n-form-item label="排序" path="sort">
            <n-input-number v-model:value="editForm.sort" :min="0" />
          </n-form-item>
          <n-form-item label="状态" path="status">
            <n-switch v-model:value="editForm.status" :checked-value="1" :unchecked-value="2" />
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
  import { h, reactive, ref, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { PlusOutlined } from '@vicons/antd';
  import { adaTableScrollX } from '@/utils/hotgo';
  import { TableList, TableEdit, TableDelete, AreaList } from '@/api/tavern';

  const dialog = useDialog();
  const message = useMessage();
  const actionRef = ref();
  const formRef = ref();
  const showEdit = ref(false);
  const saving = ref(false);
  const scrollX = adaTableScrollX();
  const areaOptions = ref<any[]>([]);

  const schemas = [
    { field: 'name', component: 'NInput', label: '桌号', componentProps: { placeholder: '请输入桌号' } },
    { field: 'areaId', component: 'NSelect', label: '区域', componentProps: { options: areaOptions, placeholder: '全部区域' } },
  ];
  const [register] = useForm({ labelWidth: 120, schemas, showActionButtonGroup: true, showResetButton: true, submitButtonText: '搜索', resetButtonText: '重置' });

  const seatStatusText = (s: number) => ({ 1: '空闲', 2: '用餐中', 3: '待清台' } as Record<number, string>)[s] || '-';

  const columns = [
    { title: 'ID', key: 'id', width: 70 },
    { title: '桌号', key: 'name', width: 110 },
    { title: '区域', key: 'areaName', width: 120 },
    { title: '容纳人数', key: 'capacity', width: 100 },
    { title: '桌码参数', key: 'qrcodeScene', width: 130 },
    { title: '实时状态', key: 'seatStatus', width: 110, render: (row) => seatStatusText(row.seatStatus) },
    { title: '排序', key: 'sort', width: 80 },
    { title: '状态', key: 'status', width: 90, render: (row) => (row.status === 1 ? '正常' : '禁用') },
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
          { label: '编辑', onClick: () => openEdit(record) },
          { label: '删除', type: 'error', onClick: () => handleDelete(record) },
        ],
      });
    },
  });

  const editForm = ref<any>({ id: 0, name: '', areaId: null, capacity: 4, sort: 0, status: 1 });
  const rules = {
    name: { required: true, message: '请输入桌号', trigger: ['blur', 'input'] },
    areaId: { required: true, type: 'number', message: '请选择区域', trigger: ['blur', 'change'] },
  };

  onMounted(async () => {
    const res: any = await AreaList({ page: 1, pageSize: 100 });
    areaOptions.value = (res.data?.list || []).map((a: any) => ({ label: a.name, value: a.id }));
  });

  async function loadDataTable(params: any) {
    const res: any = await TableList(params);
    return { data: res.data?.list || [], total: res.data?.totalCount || 0, pageCount: res.data?.pageCount || 0 };
  }

  function reloadTable() {
    actionRef.value?.reload();
  }

  function openEdit(record: any) {
    editForm.value = record ? { ...record } : { id: 0, name: '', areaId: areaOptions.value[0]?.value || null, capacity: 4, sort: 0, status: 1 };
    showEdit.value = true;
  }

  async function submitEdit() {
    try {
      await formRef.value?.validate();
      saving.value = true;
      await TableEdit(editForm.value);
      message.success('保存成功');
      showEdit.value = false;
      reloadTable();
    } finally {
      saving.value = false;
    }
  }

  function handleDelete(record: any) {
    dialog.warning({
      title: '删除确认',
      content: '确认删除桌位「' + record.name + '」吗？有进行中会话的桌位不可删除。',
      positiveText: '确认删除',
      negativeText: '取消',
      onPositiveClick: async () => {
        await TableDelete({ ids: [record.id] });
        message.success('删除成功');
        reloadTable();
      },
    });
  }
</script>
