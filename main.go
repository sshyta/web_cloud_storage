package main

import (
	_ "database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"web_cloud_storage/controllers"
	_ "web_cloud_storage/routers"
)

func init() {
	/* Инициализация базы данных*/
	orm.RegisterDataBase("default", "postgres", "user=postgres password=467912 "+
		"host=127.0.0.1 port=5432 dbname=web_cloud_storage sslmode=disable")
	orm.RegisterDriver("db", orm.DRMySQL)
}

func main() {
	// Загрузка конфигурации
	beego.LoadAppConfig("ini", "conf/app.conf")
	// Настройка сессий
	beego.BConfig.WebConfig.Session.SessionOn = true

	// Остальные настройки
	beego.SetStaticPath("/static", "static")
	beego.BConfig.Listen.HTTPAddr = "localhost"
	beego.BConfig.Listen.HTTPPort = 8181
	userController := &controllers.UserController{}
	userController.CheckAndUpdateUserTariffs()

	setLogPath("logs/app.log") // Вызов функции логов
	beego.Run()
}
