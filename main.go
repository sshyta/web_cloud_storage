package main

import (
	_ "database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"log"
	"os"
	_ "web_cloud_storage/routers"
)

func main() {
	beego.LoadAppConfig("ini", "conf/app.conf")
	setLogPath("logs/app.log")

	beego.SetStaticPath("/static", "static")

	beego.BConfig.Listen.HTTPAddr = "localhost"
	beego.BConfig.Listen.HTTPPort = 8181
	orm.RegisterDriver("db", orm.DRMySQL)
	// строка подключения: пользователь:пароль@tcp(хост:порт)/имя_базы_данных
	orm.RegisterDataBase("default", "postgres", "user=postgres password=467912 host=127.0.0.1 port=5433 dbname=web_cloud_storage sslmode=disable")

	beego.Run()
}

func setLogPath(logFilePath string) {
	logDir := logFilePath[:len(logFilePath)-len("/"+getFileName(logFilePath))]
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			log.Fatal("Не удалось создать папку %v", err)
		}
	}

	logs.SetLogger(logs.AdapterFile, `{"filename":"`+logFilePath+`"}`)
	logs.SetLogger(logs.AdapterConsole, `{"filename":"logs/app.log", "level":7, "daily":true, "maxdays":10}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	logs.SetLevel(logs.LevelDebug)
}

func getFileName(filePath string) string {
	if len(filePath) == 0 {
		return ""
	}

	for i := len(filePath) - 1; i >= 0; i-- {
		if filePath[i] == '/' || filePath[i] == '\\' {
			return filePath[i+1:]
		}
	}
	return filePath
}
