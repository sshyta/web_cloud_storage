package main

import (
	_ "database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "web_cloud_storage/routers"
)

func main() {
	beego.Run()

	orm.RegisterDriver("db", orm.DRMySQL)
	// строка подключения: пользователь:пароль@tcp(хост:порт)/имя_базы_данных
	orm.RegisterDataBase("default", "pq", "postgres:467912@tcp(127.0.0.1:5432)/web_cloud_storage")
}
