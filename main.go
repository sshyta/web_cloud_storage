package main

import (
	_ "database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"log"
	"os"
	_ "web_cloud_storage/routers"
)

func main() {
	setLogPath("logs/app.log")
	beego.BConfig.Listen.HTTPAddr = "localhost"
	beego.BConfig.Listen.HTTPPort = 8181
	orm.RegisterDriver("db", orm.DRMySQL)
	// строка подключения: пользователь:пароль@tcp(хост:порт)/имя_базы_данных
	orm.RegisterDataBase("default", "postgres", "user=postgres password=467912 host=127.0.0.1 port=5432 dbname=web_cloud_storage sslmode=disable")

	beego.Run()
}

func setLogPath(string) {
	LogPath := "logs"
	if _, err := os.Stat(LogPath); os.IsNotExist(err) {
		err := os.Mkdir(LogPath, os.ModePerm)
		if err != nil {
			log.Fatal("Не удалось создать папку %v", err)
		}
	}
	beego.BConfig.Log.FileLineNum = true
	beego.BConfig.Log.Outputs = map[string]string{
		"console": "",
		"file":    LogPath + "/app.log",
	}

	log.Printf(LogPath)
}
