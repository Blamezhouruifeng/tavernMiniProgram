package genrouter

import "hotgo/internal/controller/admin/sys"

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, sys.TavernAdmin) // 酒馆管理模块
}
