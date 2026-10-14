<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="酒馆 · 门店配置">
        先付制与营业时间在这里配置：支付成功即进入制作，未支付订单按超时时间自动取消。
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard">
      <BasicForm @register="register" @submit="handleSubmit" />
      <div class="mt-4">
        <n-button type="primary" :loading="saving" @click="handleSubmit">
          保存配置
        </n-button>
      </div>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { StoreView, StoreEdit } from '@/api/tavern';

  const message = useMessage();
  const saving = ref(false);

  const schemas = [
    { field: 'name', component: 'NInput', label: '门店名称', componentProps: { placeholder: '如 晚风民谣小酒馆' } },
    { field: 'brief', component: 'NInput', label: '简介', componentProps: { placeholder: '一句话介绍' } },
    { field: 'address', component: 'NInput', label: '地址', componentProps: { placeholder: '门店地址' } },
    { field: 'phone', component: 'NInput', label: '联系电话', componentProps: { placeholder: '门店电话' } },
    { field: 'businessOpen', component: 'NInput', label: '开市时间', componentProps: { placeholder: '18:00' } },
    { field: 'businessClose', component: 'NInput', label: '打烊时间（可跨天）', componentProps: { placeholder: '02:00' } },
    { field: 'lastOrderGap', component: 'NInputNumber', label: '打烊前停止下单(分钟)', componentProps: { min: 0, max: 240 } },
    { field: 'payExpireMinutes', component: 'NInputNumber', label: '待支付超时(分钟)', componentProps: { min: 1, max: 60 } },
    { field: 'payFirst', component: 'NSwitch', label: '先付制', componentProps: { checkedValue: true, uncheckedValue: false } },
  ];

  const [register, { setFieldsValue, validate }] = useForm({
    labelWidth: 180,
    schemas,
    showActionButtonGroup: false,
  });

  onMounted(async () => {
    const res: any = await StoreView();
    if (res?.data) {
      setFieldsValue(res.data);
    }
  });

  async function handleSubmit() {
    try {
      const values: any = await validate();
      saving.value = true;
      await StoreEdit(values);
      message.success('保存成功');
    } finally {
      saving.value = false;
    }
  }
</script>
