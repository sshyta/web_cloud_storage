package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
	"web_cloud_storage/models"
	"web_cloud_storage/utils"
	_ "web_cloud_storage/utils"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.TplName = "user.html"
}

func (c *UserController) AddUser() {
	if c.Ctx.Input.Method() == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		login := c.GetString("login")
		email := c.GetString("email")
		rolesIDStr := c.GetString("roles_id")

		if username == "" || password == "" || login == "" || email == "" || rolesIDStr == "" {
			c.Data["json"] = map[string]string{"error": "All fields are required"}
			c.ServeJSON()
			return
		}

		rolesID, err := strconv.Atoi(rolesIDStr)
		if err != nil || rolesID < 1 {
			c.Data["json"] = map[string]string{"error": "Invalid role ID"}
			c.ServeJSON()
			return
		}

		// Хеширование пароля
		hashedPassword := utils.HashPassword(password)

		user := models.Users{
			Username:           username,
			Userpass:           hashedPassword,
			Login:              login,
			WorkingEmail:       email,
			RolesID:            rolesID,
			DateOfRegistration: time.Now(),
		}

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
