package controllers

import (
	_ "database/sql"
	"github.com/astaxie/beego"
)

type StorageController struct {
	beego.Controller
}

func (storage *StorageController) Get() {
	storage.TplName = "storage.html"
}

/* добавить логику */
