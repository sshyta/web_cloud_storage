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

type TariffController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (form *MainController) Get() {
	form.TplName = "form_login.html"
}

func (storage *StorageController) Get() {
	storage.TplName = "storage.html"
}

func (tariff *TariffController) Get() {
	tariff.TplName = "tariff.html"
}

func (user *UserController) Get() {
	user.TplName = "user.html"
}
