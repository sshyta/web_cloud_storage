package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"web_cloud_storage/models"
	"web_cloud_storage/utils"
)

type UserController struct {
	beego.Controller
}

func (user *UserController) Get() {
	user.TplName = "user.html"
}

func (c *UserController) AddUser() {
	if c.Ctx.Input.Method() == "POST" {
		// Получение данных из формы
		username := c.GetString("username")
		password := c.GetString("password")
		login := c.GetString("login")
		email := c.GetString("email")
		rolesID, _ := c.GetInt("roles_id") // Получаем роль как целое число

		// Проверка обязательных полей
		if username == "" || password == "" || login == "" || email == "" || rolesID == 0 {
			c.Data["json"] = map[string]string{"error": "All fields are required"}
			c.ServeJSON()
			return
		}

		// Шифрование пароля
		hashedPassword := utils.HashPassword(password)

		// Создание нового пользователя
		user := models.Users{
			Username:           username,
			Password:           hashedPassword,
			Login:              login,
			Email:              email,
			RolesID:            rolesID, // Присваиваем значение роли
			DateOfRegistration: time.Now(),
		}

		// Сохранение пользователя в базу данных
		o := orm.NewOrm()
		_, err := o.Insert(&user)
		if err != nil {
			c.Data["json"] = map[string]string{"error": "Failed to save user: " + err.Error()}
			c.ServeJSON()
			return
		}

		c.Data["json"] = map[string]string{"success": "User added successfully"}
		c.ServeJSON()
	}
}
