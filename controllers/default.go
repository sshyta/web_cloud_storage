package controllers

import (
	_ "database/sql"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type StorageController struct {
	beego.Controller
}

func (form *MainController) Get() {
	form.TplName = "form_login.tpl"
}

func (storage *StorageController) Get() {
	storage.TplName = "storage.tpl"
}
