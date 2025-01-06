package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
)

type Users struct {
	ID                 int       `orm:"column(users_id); pk; auto" json:"id"`
	Username           string    `orm:"size(50); unique" json:"username"`
	Password           string    `orm:"column(userpass); size(255)" json:"-"` // Хранение зашифрованного пароля
	Login              string    `orm:"column(login); size(50); unique" json:"login"`
	Email              string    `orm:"column(working_email); size(250)" json:"email"`
	RolesID            int       `orm:"column(roles_id);" json:"roles_id"`
	Phone              string    `orm:"null" json:"phone"`
	DateOfRegistration time.Time `orm:"column(date_of_registration); null; auto_now_add; type(datetime)" json:"date_of_registration"`
}

func init() {
	orm.RegisterModel(new(Users))
}
