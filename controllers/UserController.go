package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
	"web_cloud_storage/models"
	"web_cloud_storage/utils"
)

type UserController struct {
	beego.Controller
}

// Get handles the GET request for the user page
func (c *UserController) Get() {
	o := orm.NewOrm()

	username := c.GetSession("username")
	if username != nil {
		var currentUser models.Users
		err := o.QueryTable("users").Filter("login", username).One(&currentUser)
		if err == nil {
			c.Data["Username"] = currentUser.Username
			c.Data["Login"] = currentUser.Login
			c.Data["WorkingEmail"] = currentUser.WorkingEmail
			c.Data["RolesID"] = currentUser.RolesID

			// Fetch tariff name
			var tariff models.Tariff
			err = o.QueryTable("tariff").Filter("tariff_id", currentUser.TariffID).One(&tariff)
			if err == nil {
				c.Data["TariffName"] = tariff.TariffName
			} else {
				beego.Error("Ошибка при получении тарифа:", err)
			}
		} else {
			beego.Error("Ошибка при получении данных пользователя:", err)
		}
	}

	var users []models.Users
	_, err := o.QueryTable("users").All(&users)
	if err != nil {
		beego.Error("Ошибка при получении пользователей:", err)
		c.Data["Users"] = []models.Users{}
	} else {
		c.Data["Users"] = users
	}

	c.TplName = "user.html"
}

// GetUsers returns a list of all users
func (c *UserController) GetUsers() {
	o := orm.NewOrm()
	var users []models.Users
	_, err := o.QueryTable("users").All(&users)
	if err != nil {
		beego.Error("Ошибка при получении пользователей:", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Не удалось получить пользователей"}
	} else {
		c.Data["json"] = users
	}
	c.ServeJSON()
}

// AddUser handles the creation of a new user
func (c *UserController) AddUser() {
	if c.Ctx.Input.Method() == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		login := c.GetString("login")
		email := c.GetString("email")
		rolesIDStr := c.GetString("roles_id")
		tariffIDStr := c.GetString("tariff_id")

		if username == "" || password == "" || login == "" || email == "" || rolesIDStr == "" || tariffIDStr == "" {
			c.Data["json"] = map[string]string{"error": "Все поля обязательны для заполнения"}
			c.ServeJSON()
			return
		}

		rolesID, err := strconv.Atoi(rolesIDStr)
		if err != nil || rolesID < 1 {
			c.Data["json"] = map[string]string{"error": "Неверный ID роли"}
			c.ServeJSON()
			return
		}

		tariffID, err := strconv.Atoi(tariffIDStr)
		if err != nil || tariffID < 1 {
			c.Data["json"] = map[string]string{"error": "Неверный ID тарифа"}
			c.ServeJSON()
			return
		}

		hashedPassword := utils.HashPassword(password)

		user := models.Users{
			Username:           username,
			Userpass:           hashedPassword,
			Login:              login,
			WorkingEmail:       email,
			RolesID:            rolesID,
			TariffID:           tariffID,
			DateOfRegistration: time.Now(),
		}

		o := orm.NewOrm()
		_, err = o.Insert(&user)
		if err != nil {
			beego.Error("Ошибка при сохранении пользователя:", err)
			c.Data["json"] = map[string]string{"error": "Не удалось сохранить пользователя: " + err.Error()}
			c.ServeJSON()
			return
		}

		c.Data["json"] = map[string]string{"success": "Пользователь успешно добавлен"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]string{"error": "Неверный метод запроса"}
		c.ServeJSON()
	}
}

// UpdateTariff updates the tariff for a specific user
func (c *UserController) UpdateTariff() {
	userID, _ := c.GetInt("user_id")
	tariffID, _ := c.GetInt("tariff_id")

	o := orm.NewOrm()
	user := models.Users{UsersID: userID}
	if err := o.Read(&user); err == nil {
		user.TariffID = tariffID
		if _, err := o.Update(&user, "TariffID"); err == nil {
			c.Data["json"] = map[string]interface{}{"status": "success"}
		} else {
			beego.Error("Ошибка при обновлении тарифа пользователя:", err)
			c.Data["json"] = map[string]interface{}{"status": "error", "message": err.Error()}
		}
	} else {
		beego.Error("Пользователь не найден:", err)
		c.Data["json"] = map[string]interface{}{"status": "error", "message": "Пользователь не найден"}
	}
	c.ServeJSON()
}

// CheckAndUpdateUserTariffs checks and updates user tariffs if necessary
func (c *UserController) CheckAndUpdateUserTariffs() {
	o := orm.NewOrm()
	var users []*models.Users
	_, err := o.QueryTable("users").All(&users)
	if err != nil {
		beego.Error("Ошибка при получении пользователей:", err)
		return
	}

	for _, user := range users {
		if user.TariffID == 0 || user.TariffID > 3 {
			user.TariffID = 1 // Устанавливаем базовый тариф по умолчанию
			_, err := o.Update(user, "TariffID")
			if err != nil {
				beego.Error("Ошибка при обновлении тарифа пользователя:", err)
			} else {
				beego.Info(fmt.Sprintf("Обновлен тариф для пользователя %s: %d", user.Username, user.TariffID))
			}
		}
	}
}
