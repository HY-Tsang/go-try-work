package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"try-work/app/api"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/hello", api.Hello)

		group.Group("/admin/", func(group *ghttp.RouterGroup) {
			group.GET("/", api.Admin.Index)
			group.POST("/signup", api.Admin.SignUp)
			group.POST("/sign_in", api.Admin.SignIn)
		})
	})
}
