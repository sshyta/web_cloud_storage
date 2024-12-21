package routers

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"web_cloud_storage/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:Post")
	beego.Router("/storage", &controllers.StorageController{}, "get:Get;post:Post;delete:Delete")
	beego.Router("/tariff", &controllers.TariffController{})
	beego.Router("/user", &controllers.UserController{})
}
