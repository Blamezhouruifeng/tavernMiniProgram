<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="酒馆 · 商品分类">
        支持分时段菜单：填写可见时段（如 18:00-23:59）即只在该时段展示。
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
            新增分类
          </n-button>
        </template>
      </BasicTable>

      <n-modal v-model:show="showEdit" preset="card" :title="editForm.id ? '编辑分类' : '新增分类'" style="width: 560px">
        <n-form ref="formRef" :model="editForm" :rules="rules" label-placement="left" label-width="120">
          <n-form-item label="分类名称" path="name">
            <n-input v-model:value="editForm.name" placeholder="如 精酿啤酒" />
          </n-form-item>
          <n-form-item label="可见时段" path="timeSlots">
            <n-input v-model:value="editForm.timeSlots" placeholder="如 18:00-23:59，留空为全天" />
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
  import { h, reactive, ref } from 'vue';
  import { NButton, useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { PlusOutlined } from '@vicons/antd';
  import { adaTableScrollX } from '@/utils/hotgo';
  import { CategoryList, CategoryEdit, CategoryDelete } from '@/api/tavern';

  const dialog = useDialog();
  const message = useMessage();
  const actionRef = ref();
  const formRef = ref();
  const showEdit = ref(false);
  const saving = ref(false);
  const scrollX = adaTableScrollX();

  const schemas = [
    { field: 'name', component: 'NInput', label: '分类名称', componentProps: { placeholder: '请输入分类名称' } },
  ];
  const [register] = useForm({ labelWidth: 120, schemas, showActionButtonGroup: true, showResetButton: true, submitButtonText: '搜索', resetButtonText: '重置' });

  const columns = [
    { title: 'ID', key: 'id', width: 80 },
    { title: '分类名称', key: 'name', width: 180 },
    { title: '可见时段', key: 'timeSlots', width: 160, render: (row) => row.timeSlots || '全天' },
    { title: '排序', key: 'sort', width: 80 },
    { title: '状态', key: 'status', width: 100, render: (row) => (row.status === 1 ? '正常' : '禁用') },
    { title: '创建时间', key: 'createdAt', width: 180 },
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

  const editForm = ref<any>({ id: 0, name: '', timeSlots: '', sort: 0, status: 1 });
  const rules = { name: { required: true, message: '请输入分类名称', trigger: ['blur', 'input'] } };

  async function loadDataTable(params: any) {
    const res: any = await CategoryList(params);
    return { data: res.data?.list || [], total: res.data?.totalCount || 0, pageCount: res.data?.pageCount || 0 };
  }

  function reloadTable() {
    actionRef.value?.reload();
  }

  async function openEdit(record: any) {
    editForm.value = record ? { ...record } : { id: 0, name: '', timeSlots: '', sort: 0, status: 1 };
    showEdit.value = true;
  }

  async function submitEdit() {
    try {
      await formRef.value?.validate();
      saving.value = true;
      await CategoryEdit(editForm.value);
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
      content: `确认删除分类「${record.name}」吗？`,
      positiveText: '确认删除',
      negativeText: '取消',
      onPositiveClick: async () => {
        await CategoryDelete({ ids: [record.id] });
        message.success('删除成功');
        reloadTable();
      },
    });
  }
</script>
