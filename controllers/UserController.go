package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
	"web_cloud_storage/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.TplName = "user.html"
}

func (c *UserController) AddUser() {
	if c.Ctx.Input.Method() == "POST" {
		// Получение данных из формы
		username := c.GetString("username")
		password := c.GetString("password")
		login := c.GetString("login")
		email := c.GetString("email")
		rolesIDStr := c.GetString("roles_id")

		// Проверка обязательных полей
		if username == "" || password == "" || login == "" || email == "" || rolesIDStr == "" {
			c.Data["json"] = map[string]string{"error": "All fields are required"}
			c.ServeJSON()
			return
		}

		// Преобразование roles_id в число
		rolesID, err := strconv.Atoi(rolesIDStr)
		if err != nil || rolesID < 1 {
			c.Data["json"] = map[string]string{"error": "Invalid role ID"}
			c.ServeJSON()
			return
		}

		// Создание нового пользователя
		user := models.Users{
			Username:           username,
			Userpass:           password,
			Login:              login,
			WorkingEmail:       email,
			RolesID:            rolesID,
			DateOfRegistration: time.Now(),
		}

		// Сохранение пользователя в базу данных
		o := orm.NewOrm()
		_, err = o.Insert(&user)
		if err != nil {
			c.Data["json"] = map[string]string{"error": "Failed to save user: " + err.Error()}
			c.ServeJSON()
			return
		}

		c.Data["json"] = map[string]string{"success": "User added successfully"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]string{"error": "Invalid request method"}
		c.ServeJSON()
	}
}
