-- 酒馆定时任务（与 internal/crons/tavern.go 的 GetName() 一致）
INSERT INTO hg_sys_cron (id, group_id, title, name, params, pattern, policy, count, sort, remark, status, created_at, updated_at) VALUES
(101, 0, '酒馆-待支付订单超时取消', 'tavern_order_expire', '', '0 */1 * * * *', 1, 0, 10, '先付制：未支付订单每分钟检查一次，超时自动取消（不出票、不占库存）', 1, NOW(), NOW()),
(102, 0, '酒馆-存酒到期处理', 'tavern_bottle_expire', '', '0 0 3 * * *', 1, 0, 20, '每天 03:00 将到期存酒标记为已过期（到期前 30/7/1 天提醒由小程序端展示）', 1, NOW(), NOW());
