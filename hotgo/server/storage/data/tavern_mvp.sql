-- ============================================================
-- 酒馆小程序 MVP 建表脚本
-- 项目：tavernMiniProgram    库：tavernMiniProgram    表前缀：hg_
-- 约定：遵循 hotgo/AGENTS.md 建表约定
--   id / created_at / updated_at / deleted_at / created_by / updated_by / status / sort
-- 说明：所有业务表预留 store_id，便于后续多门店
-- ============================================================

SET NAMES utf8mb4;

-- ----------------------------
-- 1. 门店
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_store`;
CREATE TABLE `hg_tavern_store` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '门店名称',
  `brief` varchar(255) NOT NULL DEFAULT '' COMMENT '简介',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '地址',
  `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '联系电话',
  `business_open` varchar(8) NOT NULL DEFAULT '18:00' COMMENT '开市时间',
  `business_close` varchar(8) NOT NULL DEFAULT '02:00' COMMENT '打烊时间（可跨天）',
  `last_order_gap` int(11) NOT NULL DEFAULT 30 COMMENT '打烊前多少分钟停止下单',
  `pay_first` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否先付制：1是 0否',
  `pay_expire_minutes` int(11) NOT NULL DEFAULT 5 COMMENT '待支付超时分钟数',
  `member_qrcode` varchar(255) NOT NULL DEFAULT '' COMMENT '员工码（兜底入口）路径',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1正常 2停用',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-门店';

-- ----------------------------
-- 2. 会员（客人与员工共用同一套会员体系）
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_member`;
CREATE TABLE `hg_tavern_member` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID（0为全店通用）',
  `openid` varchar(64) NOT NULL DEFAULT '' COMMENT '微信openid',
  `unionid` varchar(64) NOT NULL DEFAULT '' COMMENT '微信unionid',
  `nickname` varchar(64) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号（存酒取酒凭证）',
  `real_name` varchar(32) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `is_staff` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否员工：1是 0否',
  `staff_role` varchar(32) NOT NULL DEFAULT '' COMMENT '员工角色：waiter服务员/bar吧台/kitchen后厨/manager店长',
  `staff_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '员工状态：1在职 2停用',
  `points` int(11) NOT NULL DEFAULT 0 COMMENT '积分',
  `balance` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '储值余额',
  `total_amount` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '累计消费',
  `last_login_at` datetime DEFAULT NULL COMMENT '最后登录时间',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1正常 2禁用',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_openid` (`openid`),
  KEY `idx_phone` (`phone`),
  KEY `idx_staff` (`is_staff`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-会员';

-- ----------------------------
-- 3. 区域
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_area`;
CREATE TABLE `hg_tavern_area` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '区域名称',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1正常 2禁用',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_store` (`store_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-区域';

-- ----------------------------
-- 4. 桌位
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_table`;
CREATE TABLE `hg_tavern_table` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `area_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '区域ID',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '桌号',
  `capacity` int(11) NOT NULL DEFAULT 4 COMMENT '可容纳人数',
  `qrcode` varchar(255) NOT NULL DEFAULT '' COMMENT '桌码文件路径',
  `qrcode_scene` varchar(64) NOT NULL DEFAULT '' COMMENT '桌码scene参数',
  `seat_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '桌位状态：1空闲 2用餐中 3待清台',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1正常 2禁用',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_store_area` (`store_id`,`area_id`),
  UNIQUE KEY `uk_scene` (`qrcode_scene`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-桌位';

-- ----------------------------
-- 5. 桌台会话（一桌客人一次落座）
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_table_session`;
CREATE TABLE `hg_tavern_table_session` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `table_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '桌位ID',
  `session_no` varchar(32) NOT NULL DEFAULT '' COMMENT '会话号',
  `owner_member_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '开台会员ID',
  `guest_count` int(11) NOT NULL DEFAULT 0 COMMENT '就餐人数',
  `member_count` int(11) NOT NULL DEFAULT 0 COMMENT '参与会员数',
  `order_count` int(11) NOT NULL DEFAULT 0 COMMENT '已支付订单数',
  `total_paid_amount` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '累计已付金额',
  `guest_total_amount` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '全桌合计金额',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '会话状态：1进行中 2已归档',
  `open_at` datetime DEFAULT NULL COMMENT '开台时间',
  `close_at` datetime DEFAULT NULL COMMENT '关台时间',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_session_no` (`session_no`),
  KEY `idx_table_status` (`table_id`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-桌台会话';

-- ----------------------------
-- 6. 会话成员
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_session_member`;
CREATE TABLE `hg_tavern_session_member` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `session_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '会话ID',
  `member_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '会员ID',
  `nickname` varchar(64) NOT NULL DEFAULT '' COMMENT '昵称快照',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像快照',
  `is_owner` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否开台人：1是 0否',
  `join_at` datetime DEFAULT NULL COMMENT '加入时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_session_member` (`session_id`,`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-会话成员';

-- ----------------------------
-- 7. 商品分类
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_category`;
CREATE TABLE `hg_tavern_category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '分类名称',
  `icon` varchar(255) NOT NULL DEFAULT '' COMMENT '图标',
  `time_slots` varchar(64) NOT NULL DEFAULT '' COMMENT '可见时段，如 18:00-23:59，空为全天',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1正常 2禁用',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_store` (`store_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-商品分类';

-- ----------------------------
-- 8. 商品
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_goods`;
CREATE TABLE `hg_tavern_goods` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `category_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '分类ID',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '商品名称',
  `cover` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图',
  `brief` varchar(255) NOT NULL DEFAULT '' COMMENT '简介',
  `price` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '最低售价（展示用）',
  `tags` varchar(128) NOT NULL DEFAULT '' COMMENT '标签，逗号分隔：招牌,热销,新品',
  `is_wine` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否酒水：1是 0否',
  `only_dine_in` tinyint(1) NOT NULL DEFAULT 0 COMMENT '仅堂食：1是 0否',
  `is_cleared` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否沽清：1是 0否',
  `make_minutes` int(11) NOT NULL DEFAULT 10 COMMENT '预计制作分钟数',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1上架 2下架',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_category` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-商品';

-- ----------------------------
-- 9. 商品规格（杯/瓶/扎）
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_goods_sku`;
CREATE TABLE `hg_tavern_goods_sku` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `goods_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '商品ID',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '规格名称：杯/瓶/扎',
  `price` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '售价',
  `stock` int(11) NOT NULL DEFAULT -1 COMMENT '库存，-1为不限量',
  `unit_cups` decimal(6,2) NOT NULL DEFAULT 1.00 COMMENT '换算杯数（1瓶=N杯）',
  `is_deposit_unit` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否为整瓶可寄存单位：1是 0否',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1正常 2禁用',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_goods` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-商品规格';

-- ----------------------------
-- 10. 购物车（每人一个篮子，全桌可见）
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_cart_item`;
CREATE TABLE `hg_tavern_cart_item` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `session_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '会话ID',
  `member_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '加购会员ID（归属）',
  `goods_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '商品ID',
  `sku_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '规格ID',
  `qty` int(11) NOT NULL DEFAULT 1 COMMENT '数量',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注（忌口/做法）',
  `order_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '下单后回填订单ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_session` (`session_id`,`deleted_at`),
  KEY `idx_member` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-购物车明细';

-- ----------------------------
-- 11. 订单（先付制：支付即制作）
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_order`;
CREATE TABLE `hg_tavern_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `session_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '会话ID',
  `table_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '桌位ID',
  `table_name` varchar(32) NOT NULL DEFAULT '' COMMENT '桌号快照',
  `order_no` varchar(32) NOT NULL DEFAULT '' COMMENT '订单号',
  `order_seq` int(11) NOT NULL DEFAULT 1 COMMENT '本桌第几单',
  `member_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '下单会员ID',
  `payer_member_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '付款会员ID（可代付）',
  `nickname_snapshot` varchar(64) NOT NULL DEFAULT '' COMMENT '下单人昵称快照',
  `total_amount` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '订单总额',
  `pay_amount` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '实付金额',
  `discount_amount` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '优惠金额',
  `pay_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '支付状态：0未支付 1已支付 2部分退款 3已退款',
  `making_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '制作状态：0未开始 1制作中 2已上齐 3已完成',
  `refund_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '退款状态：0无 1申请中 2已通过 3已拒绝 4已退款',
  `pay_channel` varchar(16) NOT NULL DEFAULT '' COMMENT '支付渠道：wxpay/mock/balance',
  `out_trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '商户订单号',
  `expire_at` datetime DEFAULT NULL COMMENT '待支付超时时间',
  `pay_time` datetime DEFAULT NULL COMMENT '支付时间',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_order_no` (`order_no`),
  KEY `idx_session` (`session_id`),
  KEY `idx_member` (`member_id`),
  KEY `idx_pay_status` (`pay_status`,`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-订单';

-- ----------------------------
-- 12. 订单明细（行级制作状态）
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_order_item`;
CREATE TABLE `hg_tavern_order_item` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `order_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '订单ID',
  `goods_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '商品ID',
  `sku_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '规格ID',
  `goods_name` varchar(64) NOT NULL DEFAULT '' COMMENT '商品名称快照',
  `sku_name` varchar(32) NOT NULL DEFAULT '' COMMENT '规格名称快照',
  `unit_price` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '单价',
  `qty` int(11) NOT NULL DEFAULT 1 COMMENT '数量',
  `amount` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '小计',
  `is_wine` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否酒水（出票分单用）',
  `maker_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '行状态：0待制作 1制作中 2已上齐',
  `served_at` datetime DEFAULT NULL COMMENT '上齐时间',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_order` (`order_id`),
  KEY `idx_maker` (`maker_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-订单明细';

-- ----------------------------
-- 13. 支付流水（幂等与对账）
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_order_pay`;
CREATE TABLE `hg_tavern_order_pay` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `order_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '订单ID',
  `out_trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '商户订单号（唯一，防重复回调）',
  `transaction_id` varchar(64) NOT NULL DEFAULT '' COMMENT '第三方交易号',
  `pay_channel` varchar(16) NOT NULL DEFAULT '' COMMENT '支付渠道',
  `amount` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '支付金额',
  `pay_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态：0待支付 1已支付 2失败',
  `paid_at` datetime DEFAULT NULL COMMENT '支付时间',
  `raw_callback` text COMMENT '回调原文（对账用）',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_out_trade_no` (`out_trade_no`),
  KEY `idx_order` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-支付流水';

-- ----------------------------
-- 14. 储位（存酒柜位）
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_storage_slot`;
CREATE TABLE `hg_tavern_storage_slot` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `area` varchar(16) NOT NULL DEFAULT '' COMMENT '储区，如 A',
  `layer` varchar(16) NOT NULL DEFAULT '' COMMENT '层，如 2',
  `cell` varchar(16) NOT NULL DEFAULT '' COMMENT '格位，如 3',
  `code` varchar(32) NOT NULL DEFAULT '' COMMENT '储位号，如 A-2-3',
  `is_used` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否占用：1是 0否',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1正常 2禁用',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-存酒储位';

-- ----------------------------
-- 15. 存酒（代存 / 购酒寄存 / 退酒入存）
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_bottle_deposit`;
CREATE TABLE `hg_tavern_bottle_deposit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `deposit_no` varchar(32) NOT NULL DEFAULT '' COMMENT '存酒单号',
  `label_code` varchar(32) NOT NULL DEFAULT '' COMMENT '瓶签码（扫码认领）',
  `member_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '会员ID',
  `member_name` varchar(64) NOT NULL DEFAULT '' COMMENT '会员姓名快照',
  `phone_tail` varchar(8) NOT NULL DEFAULT '' COMMENT '手机号后四位（瓶签用）',
  `source_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '来源：1自带代存 2购酒寄存 3退酒入存',
  `source_order_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '来源订单ID（购酒/退酒入存）',
  `session_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '来源桌台会话ID',
  `goods_name` varchar(64) NOT NULL DEFAULT '' COMMENT '酒品名称',
  `spec` varchar(32) NOT NULL DEFAULT '' COMMENT '规格，如 750ml',
  `degree` varchar(16) NOT NULL DEFAULT '' COMMENT '酒精度',
  `total_ratio` decimal(5,2) NOT NULL DEFAULT 100.00 COMMENT '存入时剩余比例(%)',
  `remain_ratio` decimal(5,2) NOT NULL DEFAULT 100.00 COMMENT '当前剩余比例(%)',
  `slot_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '储位ID',
  `slot_code` varchar(32) NOT NULL DEFAULT '' COMMENT '储位号快照',
  `photo` varchar(255) NOT NULL DEFAULT '' COMMENT '封签照片',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1在库 2已取完 3已过期 4已转让 5报损',
  `deposit_at` datetime DEFAULT NULL COMMENT '存入时间',
  `expire_at` datetime DEFAULT NULL COMMENT '到期时间',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者（店员）',
  `updated_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_deposit_no` (`deposit_no`),
  UNIQUE KEY `uk_label_code` (`label_code`),
  KEY `idx_member` (`member_id`),
  KEY `idx_status` (`status`,`expire_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-存酒记录';

-- ----------------------------
-- 16. 取酒/核销流水
-- ----------------------------
DROP TABLE IF EXISTS `hg_tavern_bottle_take_log`;
CREATE TABLE `hg_tavern_bottle_take_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `store_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '门店ID',
  `deposit_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '存酒ID',
  `member_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '会员ID',
  `take_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '类型：1店内续饮 2整瓶带走 3报损',
  `take_ratio` decimal(5,2) NOT NULL DEFAULT 0.00 COMMENT '本次取用比例(%)',
  `remain_ratio` decimal(5,2) NOT NULL DEFAULT 100.00 COMMENT '取后剩余比例(%)',
  `session_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '关联桌台会话ID',
  `operator_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '操作人员工会员ID',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `take_at` datetime DEFAULT NULL COMMENT '取用时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_deposit` (`deposit_id`),
  KEY `idx_member` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='酒馆-取酒流水';

-- ============================================================
-- 追加：订单退款字段（先付制部分退款按明细行退）
-- ============================================================
ALTER TABLE `hg_tavern_order`
  ADD COLUMN `refund_amount` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '退款金额' AFTER `refund_status`,
  ADD COLUMN `refund_reason` varchar(255) NOT NULL DEFAULT '' COMMENT '退款原因' AFTER `refund_amount`,
  ADD COLUMN `refund_at` datetime DEFAULT NULL COMMENT '退款时间' AFTER `refund_reason`,
  ADD COLUMN `refund_audit_by` bigint(20) NOT NULL DEFAULT 0 COMMENT '退款审核人' AFTER `refund_at`;
