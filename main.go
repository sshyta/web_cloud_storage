package main

import (
	_ "database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "web_cloud_storage/routers"
)

func main() {
	beego.LoadAppConfig("init", "conf/app.conf")
	setLogPath("logs/app.log")
	beego.SetStaticPath("/static", "static")
	beego.BConfig.Listen.HTTPAddr = "localhost"
	beego.BConfig.Listen.HTTPPort = 8181
	orm.RegisterDriver("db", orm.DRMySQL)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=467912 "+
		"host=127.0.0.1 port=5432 dbname=web_cloud_storage sslmode=disable") /* Подключение к бд*/

	beego.Run()
}
