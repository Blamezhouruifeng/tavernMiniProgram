<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="酒馆 · 优惠券">
        勾选「新人礼」的券会在会员注册时自动发放；也可按手机号定向发放或对全体会员发放。
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
            新增优惠券
          </n-button>
        </template>
      </BasicTable>

      <n-modal v-model:show="showEdit" preset="card" :title="editForm.id ? '编辑优惠券' : '新增优惠券'" style="width: 620px">
        <n-form ref="formRef" :model="editForm" :rules="rules" label-placement="left" label-width="140">
          <n-form-item label="券名称" path="name">
            <n-input v-model:value="editForm.name" placeholder="如 新人首单礼" />
          </n-form-item>
          <n-form-item label="券类型" path="type">
            <n-select v-model:value="editForm.type" :options="[{ label: '满减', value: 1 }, { label: '折扣', value: 2 }]" />
          </n-form-item>
          <n-form-item v-if="editForm.type === 1" label="减免金额" path="amount">
            <n-input-number v-model:value="editForm.amount" :min="0" :step="5" />
          </n-form-item>
          <n-form-item v-else label="折扣（0.88=八八折）" path="discount">
            <n-input-number v-model:value="editForm.discount" :min="0.1" :max="1" :step="0.01" />
          </n-form-item>
          <n-form-item label="使用门槛(订单满)" path="minAmount">
            <n-input-number v-model:value="editForm.minAmount" :min="0" :step="10" />
          </n-form-item>
          <n-form-item label="领取后有效天数" path="validDays">
            <n-input-number v-model:value="editForm.validDays" :min="1" :max="365" />
          </n-form-item>
          <n-form-item label="发放总量(-1不限)" path="totalLimit">
            <n-input-number v-model:value="editForm.totalLimit" :min="-1" />
          </n-form-item>
          <n-form-item label="每人限领" path="perMemberLimit">
            <n-input-number v-model:value="editForm.perMemberLimit" :min="1" :max="100" />
          </n-form-item>
          <n-form-item label="新人礼" path="isNewcomer">
            <n-switch v-model:value="editForm.isNewcomer" />
          </n-form-item>
          <n-form-item label="状态" path="status">
            <n-switch v-model:value="editForm.status" :checked-value="1" :unchecked-value="2" />
          </n-form-item>
          <n-form-item label="排序" path="sort">
            <n-input-number v-model:value="editForm.sort" :min="0" />
          </n-form-item>
          <n-form-item label="备注" path="remark">
            <n-input v-model:value="editForm.remark" placeholder="使用说明" />
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

      <n-modal v-model:show="showGrant" preset="card" title="发放优惠券" style="width: 520px">
        <n-form label-placement="left" label-width="120">
          <n-form-item label="优惠券">
            <n-input :value="grantForm.name" disabled />
          </n-form-item>
          <n-form-item label="发放对象">
            <n-radio-group v-model:value="grantForm.mode">
              <n-radio value="phone">
                指定手机号
              </n-radio>
              <n-radio value="all">
                全体会员
              </n-radio>
            </n-radio-group>
          </n-form-item>
          <n-form-item v-if="grantForm.mode === 'phone'" label="手机号">
            <n-input v-model:value="grantForm.phone" placeholder="会员手机号" />
          </n-form-item>
        </n-form>
        <template #footer>
          <n-space justify="end">
            <n-button @click="showGrant = false">
              取消
            </n-button>
            <n-button type="primary" :loading="saving" @click="submitGrant">
              确认发放
            </n-button>
          </n-space>
        </template>
      </n-modal>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { NButton, NTag, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { PlusOutlined } from '@vicons/antd';
  import { adaTableScrollX } from '@/utils/hotgo';
  import { CouponList, CouponEdit, CouponGrant } from '@/api/tavern';

  const message = useMessage();
  const actionRef = ref();
  const formRef = ref();
  const showEdit = ref(false);
  const showGrant = ref(false);
  const saving = ref(false);
  const scrollX = adaTableScrollX();

  const schemas = [{ field: 'name', component: 'NInput', label: '券名称', componentProps: { placeholder: '请输入券名称' } }];
  const [register] = useForm({ labelWidth: 120, schemas, showActionButtonGroup: true, showResetButton: true, submitButtonText: '搜索', resetButtonText: '重置' });

  const columns = [
    { title: 'ID', key: 'id', width: 70 },
    { title: '券名称', key: 'name', width: 200 },
    {
      title: '优惠',
      key: 'benefit',
      width: 140,
      render(row: any) {
        return row.type === 2 ? (row.discount * 10).toFixed(1) + ' 折' : '减 ¥' + Number(row.amount).toFixed(2);
      },
    },
    { title: '门槛', key: 'minAmount', width: 110, render: (row) => '满 ¥' + Number(row.minAmount).toFixed(0) },
    { title: '有效期(天)', key: 'validDays', width: 110 },
    { title: '每人限领', key: 'perMemberLimit', width: 100 },
    {
      title: '新人礼',
      key: 'isNewcomer',
      width: 100,
      render(row: any) {
        return h(NTag, { type: row.isNewcomer ? 'success' : 'default', bordered: false }, { default: () => (row.isNewcomer ? '是' : '否') });
      },
    },
    { title: '已发放', key: 'grantedCount', width: 90 },
    { title: '已使用', key: 'usedCount', width: 90 },
    { title: '状态', key: 'status', width: 90, render: (row) => (row.status === 1 ? '启用' : '停用') },
  ];

  const actionColumn = reactive({
    width: 180,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record: any) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          { label: '编辑', onClick: () => openEdit(record) },
          { label: '发放', onClick: () => openGrant(record) },
        ],
      });
    },
  });

  const editForm = ref<any>({ id: 0, name: '', type: 1, amount: 20, minAmount: 100, discount: 0.88, validDays: 30, totalLimit: -1, perMemberLimit: 1, isNewcomer: false, status: 1, sort: 0, remark: '' });
  const grantForm = ref<any>({ couponId: 0, name: '', mode: 'phone', phone: '' });
  const rules = { name: { required: true, message: '请输入券名称', trigger: ['blur', 'input'] } };

  async function loadDataTable(params: any) {
    const res: any = await CouponList(params);
    return { data: res.data?.list || [], total: res.data?.totalCount || 0, pageCount: res.data?.pageCount || 0 };
  }

  function reloadTable() {
    actionRef.value?.reload();
  }

  function openEdit(record: any) {
    editForm.value = record ? { ...record } : { id: 0, name: '', type: 1, amount: 20, minAmount: 100, discount: 0.88, validDays: 30, totalLimit: -1, perMemberLimit: 1, isNewcomer: false, status: 1, sort: 0, remark: '' };
    showEdit.value = true;
  }

  async function submitEdit() {
    try {
      await formRef.value?.validate();
      saving.value = true;
      await CouponEdit(editForm.value);
      message.success('保存成功');
      showEdit.value = false;
      reloadTable();
    } finally {
      saving.value = false;
    }
  }

  function openGrant(record: any) {
    grantForm.value = { couponId: record.id, name: record.name, mode: 'phone', phone: '' };
    showGrant.value = true;
  }

  async function submitGrant() {
    saving.value = true;
    try {
      const payload: any = { couponId: grantForm.value.couponId };
      if (grantForm.value.mode === 'all') {
        payload.all = true;
      } else {
        payload.phone = grantForm.value.phone;
      }
      const res: any = await CouponGrant(payload);
      message.success('已发放 ' + (res.data?.count || 0) + ' 张');
      showGrant.value = false;
      reloadTable();
    } finally {
      saving.value = false;
    }
  }
</script>
