# tavernMiniProgram

酒馆小程序 —— 由 **HotGo V2 全栈框架**（Go 后端 + Vue3 后台）与 **uni-app 小程序端**组成的 monorepo。

当前处于**框架就绪、业务未实现**阶段：两端都是可运行的上游骨架，酒馆业务（门店 / 酒水 / 订单等）尚未开发。

## 目录结构

| 目录 | 说明 | 技术栈 |
|---|---|---|
| `hotgo/server/` | 后端主程序（多应用入口：admin / api / home / websocket / TCP） | Go 1.24+、GoFrame v2、JWT + casbin |
| `hotgo/web/` | 后台管理前端 | Vue3、Vite5、TypeScript、Naive UI、Pinia |
| `hotgo/docs/guide-zh-CN/` | 后端官方中文文档 | — |
| `uniapp/` | 小程序端（H5 / 微信小程序） | uni-app、Vue3、Wot UI、Alova、Pinia、UnoCSS |

## 环境要求

- Go >= 1.24（后端）
- Node >= 20.19 / 22.12 / 24 与 pnpm（两端前端）
- MySQL 8（库名 `tavernMiniProgram`，表前缀 `hg_`）
- Redis（配置为 db 2，缓存 / 队列）

## 本地启动

### 1. 后端

```bash
# 首次：复制配置模板并填写数据库连接
cp hotgo/server/manifest/config/config.example.yaml hotgo/server/manifest/config/config.yaml

# 初始化数据库（表前缀 hg_）
mysql -uroot -p tavernMiniProgram < hotgo/server/storage/data/hotgo.sql

cd hotgo/server && go run main.go
```

- 服务地址：http://127.0.0.1:8000
- Swagger：http://127.0.0.1:8000/swagger/
- OpenAPI：http://127.0.0.1:8000/api.json
- TCP 服务：:8099

> 数据库连接串形如 `mysql:用户名:密码@tcp(127.0.0.1:3306)/tavernMiniProgram?loc=Local&parseTime=true&charset=utf8mb4`，
> 用户名与密码之间是 `:`；写成 `root@:root@tcp(...)` 会被解析成用户 `root@` 并报 `Error 1045 Access denied`。

### 2. 后台管理前端

```bash
cd hotgo/web && pnpm install && pnpm dev   # http://localhost:8001（代理 /admin → :8000）
```

### 3. 小程序端

```bash
cd uniapp && pnpm install
pnpm dev:h5          # H5 调试
pnpm dev:mp-weixin   # 微信小程序
```

> 小程序端接口地址由 `uniapp/.env.*` 的 `VITE_API_BASE_URL` 控制，当前仍指向示例地址，接入后端时改为 `http://127.0.0.1:8000/api`。

## 开发规范（写代码前必读）

- 后端：[`hotgo/AGENTS.md`](./hotgo/AGENTS.md) —— 分层契约 `api → controller → service → logic → dao`、字典规范、插件机制、错误处理、建表约定。
- 小程序端：[`uniapp/AGENTS.md`](./uniapp/AGENTS.md) —— 目录边界、OpenSpec 变更流程、跨端约束。

## 注意

`hotgo/server/manifest/config/config.yaml` 含本地数据库凭据，已被 `hotgo/server/.gitignore` 排除，请以 `config.example.yaml` 为准自行创建。
