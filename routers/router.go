package routers

import (
	"github.com/astaxie/beego"
	"web_cloud_storage/controllers"
	"web_cloud_storage/filters"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:Post")
	beego.Router("/user", &controllers.UserController{}, "get:Get")
	beego.Router("/user/add", &controllers.UserController{}, "post:AddUser")
	beego.Router("/user/update-tariff", &controllers.UserController{}, "post:UpdateTariff")
	beego.Router("/user/list", &controllers.UserController{}, "get:GetUsers") // Роут для отображения пользователей

	beego.Router("/logout", &controllers.MainController{}, "post:Logout")

	beego.Router("/storage", &controllers.StorageController{}, "get:Get")
	beego.Router("/storage/upload", &controllers.StorageController{}, "post:Upload")
	beego.Router("/storage/files", &controllers.StorageController{}, "get:ListFiles")
	beego.Router("/storage/delete", &controllers.StorageController{}, "post:Delete")
	beego.Router("/storage/view", &controllers.StorageController{}, "get:View")
	beego.Router("/storage/edit", &controllers.StorageController{}, "post:Edit")
	beego.Router("/storage/download", &controllers.StorageController{}, "get:Download")
	beego.Router("/storage/info", &controllers.StorageController{}, "get:GetStorageInfo")

	beego.Router("/tariff", &controllers.TariffController{}, "get:Get")

	beego.Router("/admin", &controllers.AdminController{}, "get:Get")
	beego.Router("/admin/user-info", &controllers.AdminController{}, "get:GetUserInfo")
	beego.Router("/admin/storage-report", &controllers.AdminController{}, "get:GetStorageReport")
	beego.Router("/admin/file-list", &controllers.AdminController{}, "get:GetFileList")

	beego.InsertFilter("/admin/*", beego.BeforeRouter, filters.AdminMiddleware)
	beego.InsertFilter("/storage/*", beego.BeforeRouter, filters.AuthMiddleware)
}
