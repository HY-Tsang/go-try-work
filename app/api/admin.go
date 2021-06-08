package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"try-work/app/model"
	"try-work/app/service"
	"try-work/library/response"
)

var Admin = adminApi{}

type adminApi struct{}

func (*adminApi) SignUp(r *ghttp.Request) {
	var (
		apiReq     *model.AdministratorApiSignUpReq
		serviceReq *model.AdministratorServiceSignUpReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Administrator.SignUp(serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

func (*adminApi) SignIn(r *ghttp.Request) {
	var (
		data *model.AdministratorApiSignInReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Administrator.SignIn(r.Context(), data.Username, data.Password); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// Index is a demonstration route handler for output "Hello World!".
func (*adminApi) Index(r *ghttp.Request) {
	r.Response.Writeln("Hello Admin!")
}
