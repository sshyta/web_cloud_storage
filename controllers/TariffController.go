package controllers

import (
	_ "database/sql"
	"github.com/astaxie/beego"
)

type TariffController struct {
	beego.Controller
}

func (tariff *TariffController) Get() {
	tariff.TplName = "tariff.html"
}
