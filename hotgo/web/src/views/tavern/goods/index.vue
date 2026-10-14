<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="酒馆 · 商品与规格">
        规格即「杯 / 整瓶 / 扎」：勾选「可寄存单位」表示该规格是整瓶酒，客人喝不完可转存；换算杯数用于存酒扣减。
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
            新增商品
          </n-button>
        </template>
      </BasicTable>

      <n-modal v-model:show="showEdit" preset="card" :title="editForm.id ? '编辑商品' : '新增商品'" style="width: 760px">
        <n-form ref="formRef" :model="editForm" :rules="rules" label-placement="left" label-width="120">
          <n-form-item label="商品名称" path="name">
            <n-input v-model:value="editForm.name" placeholder="如 IPA 印度淡色艾尔" />
          </n-form-item>
          <n-form-item label="分类" path="categoryId">
            <n-select v-model:value="editForm.categoryId" :options="categoryOptions" placeholder="选择分类" />
          </n-form-item>
          <n-form-item label="简介" path="brief">
            <n-input v-model:value="editForm.brief" placeholder="口味描述" />
          </n-form-item>
          <n-form-item label="标签" path="tags">
            <n-input v-model:value="editForm.tags" placeholder="逗号分隔，如 招牌,热销" />
          </n-form-item>
          <n-form-item label="预计制作(分钟)" path="makeMinutes">
            <n-input-number v-model:value="editForm.makeMinutes" :min="0" :max="120" />
          </n-form-item>
          <n-form-item label="属性" path="attr">
            <n-space>
              <n-checkbox v-model:checked="editForm.isWine">
                酒水
              </n-checkbox>
              <n-checkbox v-model:checked="editForm.onlyDineIn">
                仅堂食
              </n-checkbox>
              <n-checkbox v-model:checked="editForm.isCleared">
                沽清（暂停售卖）
              </n-checkbox>
            </n-space>
          </n-form-item>
          <n-form-item label="状态" path="status">
            <n-switch v-model:value="editForm.status" :checked-value="1" :unchecked-value="2" :rail-style="() => ({ style: 'width: 60px' })" />
          </n-form-item>
          <n-form-item label="排序" path="sort">
            <n-input-number v-model:value="editForm.sort" :min="0" />
          </n-form-item>

          <n-divider>规格（杯 / 整瓶 / 扎）</n-divider>
          <div v-for="(sku, index) in editForm.skus" :key="index" class="mb-3 flex items-center gap-2">
            <n-input v-model:value="sku.name" placeholder="规格名，如 整瓶 (700ml)" style="width: 200px" />
            <n-input-number v-model:value="sku.price" :min="0" placeholder="价格" style="width: 120px" />
            <n-input-number v-model:value="sku.unitCups" :min="0.1" :step="0.1" placeholder="换算杯数" style="width: 130px" />
            <n-checkbox v-model:checked="sku.isDepositUnit">
              可寄存
            </n-checkbox>
            <n-button size="small" type="error" @click="editForm.skus.splice(index, 1)">
              删除
            </n-button>
          </div>
          <n-button size="small" @click="editForm.skus.push({ id: 0, name: '', price: 0, stock: -1, unitCups: 1, isDepositUnit: false })">
            + 添加规格
          </n-button>
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
  import { GoodsList, GoodsView, GoodsEdit, GoodsDelete, GoodsStatus, CategoryList } from '@/api/tavern';

  const dialog = useDialog();
  const message = useMessage();
  const actionRef = ref();
  const formRef = ref();
  const showEdit = ref(false);
  const saving = ref(false);
  const scrollX = adaTableScrollX();
  const categoryOptions = ref<any[]>([]);

  const schemas = [
    { field: 'name', component: 'NInput', label: '商品名称', componentProps: { placeholder: '请输入名称' } },
    { field: 'categoryId', component: 'NSelect', label: '分类', componentProps: { options: categoryOptions, placeholder: '全部' } },
    {
      field: 'status',
      component: 'NSelect',
      label: '状态',
      componentProps: {
        placeholder: '全部',
        options: [
          { label: '上架', value: 1 },
          { label: '下架', value: 2 },
        ],
      },
    },
  ];
  const [register] = useForm({ labelWidth: 120, schemas, showActionButtonGroup: true, showResetButton: true, submitButtonText: '搜索', resetButtonText: '重置' });

  const columns = [
    { title: 'ID', key: 'id', width: 70 },
    { title: '商品名称', key: 'name', width: 200 },
    { title: '分类', key: 'categoryName', width: 130 },
    { title: '展示价', key: 'price', width: 100, render: (row) => '¥' + Number(row.price || 0).toFixed(2) },
    { title: '酒水', key: 'isWine', width: 80, render: (row) => (row.isWine ? '是' : '否') },
    { title: '仅堂食', key: 'onlyDineIn', width: 90, render: (row) => (row.onlyDineIn ? '是' : '否') },
    { title: '沽清', key: 'isCleared', width: 90, render: (row) => (row.isCleared ? '已沽清' : '正常') },
    { title: '制作(分钟)', key: 'makeMinutes', width: 110 },
    {
      title: '上架状态',
      key: 'status',
      width: 110,
      render(row: any) {
        return h(
          'span',
          { style: row.status === 1 ? 'color:#18a058' : 'color:#d03050' },
          row.status === 1 ? '上架' : '下架'
        );
      },
    },
  ];

  const actionColumn = reactive({
    width: 200,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record: any) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          { label: '编辑', onClick: () => openEdit(record) },
          {
            label: record.status === 1 ? '下架' : '上架',
            onClick: () => toggleStatus(record),
          },
          { label: '删除', type: 'error', onClick: () => handleDelete(record) },
        ],
      });
    },
  });

  const editForm = ref<any>({ id: 0, categoryId: null, name: '', brief: '', tags: '', isWine: false, onlyDineIn: false, isCleared: false, makeMinutes: 10, status: 1, sort: 0, skus: [] });
  const rules = {
    name: { required: true, message: '请输入商品名称', trigger: ['blur', 'input'] },
    categoryId: { required: true, type: 'number', message: '请选择分类', trigger: ['blur', 'change'] },
  };

  onMounted(async () => {
    const res: any = await CategoryList({ page: 1, pageSize: 100 });
    categoryOptions.value = (res.data?.list || []).map((c: any) => ({ label: c.name, value: c.id }));
  });

  async function loadDataTable(params: any) {
    const res: any = await GoodsList(params);
    return { data: res.data?.list || [], total: res.data?.totalCount || 0, pageCount: res.data?.pageCount || 0 };
  }

  function reloadTable() {
    actionRef.value?.reload();
  }

  async function openEdit(record: any) {
    if (!record) {
      editForm.value = { id: 0, categoryId: categoryOptions.value[0]?.value || null, name: '', brief: '', tags: '', isWine: false, onlyDineIn: false, isCleared: false, makeMinutes: 10, status: 1, sort: 0, skus: [{ id: 0, name: '杯', price: 0, stock: -1, unitCups: 1, isDepositUnit: false }] };
      showEdit.value = true;
      return;
    }
    const res: any = await GoodsView({ id: record.id });
    const data = res.data || record;
    editForm.value = {
      ...data,
      tags: data.tags || '',
      skus: (data.skus && data.skus.length ? data.skus : [{ id: 0, name: '标准份', price: data.price || 0, stock: -1, unitCups: 1, isDepositUnit: false }]).map((s: any) => ({ ...s, unitCups: s.unitCups || 1, stock: s.stock === 0 ? -1 : s.stock })),
    };
    showEdit.value = true;
  }

  async function submitEdit() {
    try {
      await formRef.value?.validate();
      if (!editForm.value.skus.length) {
        message.warning('请至少配置一个规格');
        return;
      }
      saving.value = true;
      await GoodsEdit(editForm.value);
      message.success('保存成功');
      showEdit.value = false;
      reloadTable();
    } finally {
      saving.value = false;
    }
  }

  async function toggleStatus(record: any) {
    await GoodsStatus({ id: record.id, status: record.status === 1 ? 2 : 1 });
    message.success('操作成功');
    reloadTable();
  }

  function handleDelete(record: any) {
    dialog.warning({
      title: '删除确认',
      content: '确认删除商品「' + record.name + '」及其全部规格吗？',
      positiveText: '确认删除',
      negativeText: '取消',
      onPositiveClick: async () => {
        await GoodsDelete({ ids: [record.id] });
        message.success('删除成功');
        reloadTable();
      },
    });
  }
</script>
