package routers

import (
	"fmt"
	"test01/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/user/insertUser", &controllers.UserController{}, "get:InsertUser")
	beego.Router("/api/user/infoUser", &controllers.UserController{}, "get:InfoUser")
	beego.Router("/api/user/listUser", &controllers.UserController{}, "get:ListUser")

	beego.Get("/api/hello", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})

	beego.InsertFilter("/api/*.*", beego.BeforeExec, func(ctx *context.Context) {
		fmt.Print("过滤器")
	})
}
