package routers

import (
	"github.com/astaxie/beego"
	"github.com/skdong/zeus/pkg/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.SetStaticPath("/static", "static")
	//beego.Get("*", func(ctx *context.Context) {
	//	ctx.Redirect(302, indexUrl)
	//})
	beego.Router("/api/winds", &controllers.WindController{}, "*:Get")
	beego.Router("/api/ws/winds", &controllers.WindWSController{}, "*:Get")
}
