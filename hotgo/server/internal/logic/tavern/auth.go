package tavern

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"

	"hotgo/internal/library/contexts"
	"hotgo/internal/library/response"
	"hotgo/internal/library/token"
)

// publicPaths 无需登录即可访问的路径（相对 /api 前缀）
var publicPaths = []string{
	"/tavern/store/info",
	"/tavern/menu/list",
	"/tavern/member/login",
}

// isPublicPath 是否免登录路径
func isPublicPath(path string) bool {
	for _, p := range publicPaths {
		if gstr.Equal(path, p) || gstr.HasSuffix(path, p) {
			return true
		}
	}
	return false
}

// Auth 酒馆小程序鉴权中间件
// 说明：客人与员工共用同一套会员体系，登录态来自会员 token；
//
//	员工接口在此基础上再校验会员的「是员工」标记（见 logic/tavern/staff.go）。
func Auth(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		path = r.URL.Path
	)

	if isPublicPath(path) {
		// 已登录时也把身份放进上下文，便于个性化
		if user, err := token.ParseLoginUser(r); err == nil {
			contexts.SetUser(ctx, user)
		}
		r.Middleware.Next()
		return
	}

	user, err := token.ParseLoginUser(r)
	if err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
		return
	}
	contexts.SetUser(ctx, user)

	r.Middleware.Next()
}
