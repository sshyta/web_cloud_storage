package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Users struct {
	UsersID            int       `orm:"column(users_id); pk; auto"`
	Username           string    `orm:"column(username); size(50); unique"`
	Userpass           string    `orm:"column(userpass); size(64)"` // SHA-256 хэш имеет длину 64 символа
	Login              string    `orm:"column(login); size(50)"`
	WorkingEmail       string    `orm:"column(working_email); size(250);null"`
	RolesID            int       `orm:"column(roles_id)"`
	DateOfRegistration time.Time `orm:"column(date_of_registration); type(date); null"`
	TariffID           int       `orm:"column(tariff_id);null"`
}

func init() {
	orm.RegisterModel(new(Users))
}
