import { http } from '@/utils/http/axios';

/** 酒馆门店配置 */
export function StoreView() {
  return http.request({ url: '/tavern/store/view', method: 'get' });
}
export function StoreEdit(params) {
  return http.request({ url: '/tavern/store/edit', method: 'POST', params });
}

/** 区域 */
export function AreaList(params) {
  return http.request({ url: '/tavern/area/list', method: 'get', params });
}
export function AreaEdit(params) {
  return http.request({ url: '/tavern/area/edit', method: 'POST', params });
}
export function AreaDelete(params) {
  return http.request({ url: '/tavern/area/delete', method: 'POST', params });
}

/** 桌位 */
export function TableList(params) {
  return http.request({ url: '/tavern/table/list', method: 'get', params });
}
export function TableEdit(params) {
  return http.request({ url: '/tavern/table/edit', method: 'POST', params });
}
export function TableDelete(params) {
  return http.request({ url: '/tavern/table/delete', method: 'POST', params });
}

/** 分类 */
export function CategoryList(params) {
  return http.request({ url: '/tavern/category/list', method: 'get', params });
}
export function CategoryEdit(params) {
  return http.request({ url: '/tavern/category/edit', method: 'POST', params });
}
export function CategoryDelete(params) {
  return http.request({ url: '/tavern/category/delete', method: 'POST', params });
}

/** 商品（含规格） */
export function GoodsList(params) {
  return http.request({ url: '/tavern/goods/list', method: 'get', params });
}
export function GoodsView(params) {
  return http.request({ url: '/tavern/goods/view', method: 'get', params });
}
export function GoodsEdit(params) {
  return http.request({ url: '/tavern/goods/edit', method: 'POST', params });
}
export function GoodsDelete(params) {
  return http.request({ url: '/tavern/goods/delete', method: 'POST', params });
}
export function GoodsStatus(params) {
  return http.request({ url: '/tavern/goods/status', method: 'POST', params });
}

/** 订单查询 */
export function OrderList(params) {
  return http.request({ url: '/tavern/order/list', method: 'get', params });
}

/** 存酒台账 */
export function BottleList(params) {
  return http.request({ url: '/tavern/bottle/list', method: 'get', params });
}
export function BottleUpdate(params) {
  return http.request({ url: '/tavern/bottle/update', method: 'POST', params });
}
