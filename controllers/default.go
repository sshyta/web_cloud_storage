package controllers

import (
	_ "database/sql"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (form *MainController) Get() {
	form.TplName = "form_registration.tpl"
}
