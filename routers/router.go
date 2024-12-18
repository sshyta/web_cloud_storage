package routers

import (
	"github.com/astaxie/beego"
	"web_cloud_storage/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/storage", &controllers.StorageController{})
	beego.Router("/tariff", &controllers.TariffController{})
	beego.Router("/user", &controllers.UserController{})
}
