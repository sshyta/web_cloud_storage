package main

import (
	"log"
	"os"

	"github.com/astaxie/beego/logs"
)

func setLogPath(logFilePath string) {
	logDir := logFilePath[:len(logFilePath)-len("/"+getFileName(logFilePath))]
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Не удалось создать папку: %v", err)
		}
	}
	logs.SetLevel(logs.LevelDebug)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(7)
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+logFilePath+`"}`)
	logs.SetLogger(logs.AdapterConsole, `{"filename":"logs/app.log", "daily":true, "maxdays":10}`)
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

/* сделать нормальную глубину логов */
