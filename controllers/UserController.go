package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

type UserController struct {
	beego.Controller
}

func (user *UserController) Get() {
	user.TplName = "user.html"
}
