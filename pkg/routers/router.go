package routers

import (
	"github.com/astaxie/beego"
	"github.com/skdong/zeus/pkg/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/winds", &controllers.WindController{}, "*:Get")
	beego.Router("/ws/winds", &controllers.WindWSController{}, "*:Get")
}
