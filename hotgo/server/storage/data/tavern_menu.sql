-- 酒馆管理菜单（后台侧边栏）
INSERT INTO hg_admin_menu (id, pid, level, tree, title, name, path, icon, type, redirect, permissions, permission_name, component, always_show, active_menu, is_root, is_frame, keep_alive, hidden, affix, sort, remark, status, created_at, updated_at) VALUES
(2432, 0, 1, '', '酒馆管理', 'Tavern', '/tavern', 'ShopOutlined', 1, '', '', '', 'LAYOUT', 0, '', 0, 0, 1, 0, 0, 60, '酒馆小程序业务', 1, NOW(), NOW()),
(2433, 2432, 2, '2432', '门店配置', 'TavernStore', 'store', '', 2, '', '', '', '/tavern/store/index', 0, '', 0, 0, 1, 0, 0, 10, '', 1, NOW(), NOW()),
(2434, 2432, 2, '2432', '桌位管理', 'TavernTable', 'table', '', 2, '', '', '', '/tavern/table/index', 0, '', 0, 0, 1, 0, 0, 20, '', 1, NOW(), NOW()),
(2435, 2432, 2, '2432', '商品分类', 'TavernCategory', 'category', '', 2, '', '', '', '/tavern/category/index', 0, '', 0, 0, 1, 0, 0, 30, '', 1, NOW(), NOW()),
(2436, 2432, 2, '2432', '商品与规格', 'TavernGoods', 'goods', '', 2, '', '', '', '/tavern/goods/index', 0, '', 0, 0, 1, 0, 0, 40, '', 1, NOW(), NOW()),
(2437, 2432, 2, '2432', '订单查询', 'TavernOrder', 'order', '', 2, '', '', '', '/tavern/order/index', 0, '', 0, 0, 1, 0, 0, 50, '', 1, NOW(), NOW()),
(2438, 2432, 2, '2432', '存酒台账', 'TavernBottle', 'bottle', '', 2, '', '', '', '/tavern/bottle/index', 0, '', 0, 0, 1, 0, 0, 60, '', 1, NOW(), NOW());
INSERT INTO hg_admin_menu (id, pid, level, tree, title, name, path, icon, type, redirect, permissions, permission_name, component, always_show, active_menu, is_root, is_frame, keep_alive, hidden, affix, sort, remark, status, created_at, updated_at)
VALUES (2439, 2432, 2, '2432', '异常订单与对账', 'TavernMonitor', 'monitor', '', 2, '', '', '', '/tavern/monitor/index', 0, '', 0, 0, 1, 0, 0, 70, '支付成功未出票补单、退款待审、订单与流水对账', 1, NOW(), NOW())
ON DUPLICATE KEY UPDATE title=VALUES(title), component=VALUES(component);
