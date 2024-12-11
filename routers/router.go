package routers

import (
	"github.com/astaxie/beego"
	"web_cloud_storage/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/form", &controllers.MainController{})
	beego.Router("/storage", &controllers.MainController{})
}
