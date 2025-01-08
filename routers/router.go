package routers

import (
	"github.com/astaxie/beego"
	"web_cloud_storage/controllers"
	"web_cloud_storage/filters"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:Post")
	beego.Router("/logout", &controllers.MainController{}, "post:Logout")

	beego.Router("/storage", &controllers.StorageController{})
	beego.Router("/storage/upload", &controllers.StorageController{}, "post:Upload")
	beego.Router("/storage/files", &controllers.StorageController{}, "get:ListFiles")
	beego.Router("/storage/delete", &controllers.StorageController{}, "post:Delete")

	beego.Router("/tariff", &controllers.TariffController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "post:AddUser")

	// Применяем middleware для маршрутов
	beego.InsertFilter("/storage/*", beego.BeforeRouter, filters.AuthMiddleware)

}
