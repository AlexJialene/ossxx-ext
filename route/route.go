package route

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"ossxx-ext/app/controller"
	"ossxx-ext/app/controller/stor"
	"ossxx-ext/app/interceptor"
)

func init() {
	server := g.Server()

	server.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(interceptor.DefaultInterceptor)

		group.ALL("/", controller.Index)
		group.GET("/auth", controller.Auth)
		group.GET("/logout", controller.Logout)
	})

	server.Group("/stor", func(group *ghttp.RouterGroup) {
		group.Middleware(interceptor.DefaultInterceptor)

		group.GET("/list", stor.List)
	})

}
