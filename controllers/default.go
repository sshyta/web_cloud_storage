package controllers

import (
	_ "database/sql"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (form *MainController) Get() {
	form.TplName = "form_login.tpl"
}

func (storage *MainController) Post() {
	storage.TplName = "storage.tpl"
}
