package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
	"log"
	"net/http"
	"web_cloud_storage/utils"
)

type MainController struct {
	beego.Controller
}

// Отображение страницы входа
func (form *MainController) Get() {
	form.TplName = "form_login.html" // Отображаем форму входа
}

// Обработка логина
func (form *MainController) Post() {
	login := form.GetString("login")
	password := form.GetString("password")

	if login == "" || password == "" {
		form.SetSession("username", login) // После успешного логина
		form.Data["Error"] = "Login and password cannot be empty"
		form.TplName = "form_login.html"
		return
	}

	if isValidUser(login, password) {
		// Сохранение логина в сессии
		form.SetSession("username", login)
		form.Redirect("/storage", http.StatusFound)
		return
	}

	form.Data["Error"] = "Invalid login or password"
	form.TplName = "form_login.html"
}

// Отображение страницы хранилища
func (form *MainController) GetStorage() {
	username := form.GetSession("username")
	if username == nil {
		form.Redirect("/", http.StatusFound) // Если нет сессии, перенаправляем на логин
		return
	}

	form.Data["Username"] = username
	form.TplName = "storage.html"
}

// Выход из системы

func (c *MainController) Logout() {
	// Удаляем сессию
	c.DestroySession()

	// Отправляем успешный ответ
	c.Data["json"] = map[string]interface{}{
		"status":  "success",
		"message": "Logged out successfully",
	}
	c.ServeJSON()
}

// Функция проверки логина и пароля
func isValidUser(login, password string) bool {
	connStr := "user=postgres password=467912 host=127.0.0.1 port=5432 dbname=web_cloud_storage sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Database connection error:", err)
		return false
	}
	defer db.Close()

	var storedPassword string
	err = db.QueryRow("SELECT userpass FROM users WHERE login = $1", login).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		log.Println("No user found with login:", login)
		return false
	} else if err != nil {
		log.Println("Error fetching user:", err)
		return false
	}

	// Проверка хэша пароля
	hashedPassword := utils.HashPassword(password)
	return hashedPassword == storedPassword
}
