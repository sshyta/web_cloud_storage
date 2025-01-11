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
	beego.Router("/storage/view", &controllers.StorageController{}, "get:View")
	beego.Router("/storage/edit", &controllers.StorageController{}, "post:Edit")
	beego.Router("/storage/download", &controllers.StorageController{}, "get:Download")

	beego.Router("/tariff", &controllers.TariffController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "post:AddUser")
	beego.Router("/user/list", &controllers.UserController{}, "get:GetUsers") // Роут для отображения пользователей

	beego.InsertFilter("/storage/*", beego.BeforeRouter, filters.AuthMiddleware)

}
