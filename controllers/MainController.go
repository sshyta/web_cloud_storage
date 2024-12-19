package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
	"log"
	"net/http"
)

type MainController struct {
	beego.Controller
}

func (form *MainController) Get() {
	form.TplName = "form_login.html"
}

func (form *MainController) Post() {
	login := form.GetString("login")
	password := form.GetString("password")

	// Проверка логина и пароля
	if isValidUser(login, password) {
		form.Redirect("/storage", http.StatusFound)
		return
	}

	// Сообщение об ошибке
	form.Data["Error"] = "Invalid login or password"
	form.TplName = "form_login.html"
}

func isValidUser(login, password string) bool {
	connStr := "user=postgres password=467912 host=127.0.0.1 port=5432 dbname=web_cloud_storage sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Database connection error:", err)
		return false
	}
	defer db.Close()

	var storedPassword string
	// SQL-запрос ищет по полю login и возвращает userpass
	err = db.QueryRow("SELECT userpass FROM users WHERE login = $1", login).Scan(&storedPassword)
	if err != nil {
		log.Println("Error fetching user:", err)
		return false
	}

	// Проверка пароля
	return storedPassword == password
}
