package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"

	apiTavern "hotgo/internal/controller/api/tavern"
	tavernLogic "hotgo/internal/logic/tavern"
)

// Tavern 酒馆小程序端路由
// 说明：
//   - 与客人共用同一套会员体系；鉴权由 tavernLogic.Auth 负责（公开路径见其中 publicPaths）。
//   - 员工工作台接口在此之上再校验会员的「是员工」标记（logic/tavern/staff.go）。
func Tavern(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/tavern", func(group *ghttp.RouterGroup) {
		// H5 端跨域调试：小程序端不发送 Origin，不受影响；生产可按需收紧 allowOrigin
		group.Middleware(ghttp.MiddlewareCORS)
		group.Middleware(tavernLogic.Auth)
		group.Bind(
			apiTavern.Store,
			apiTavern.Member,
			apiTavern.Menu,
			apiTavern.Session,
			apiTavern.Cart,
			apiTavern.Order,
			apiTavern.Bottle,
			apiTavern.Staff,
		)
	})
}
