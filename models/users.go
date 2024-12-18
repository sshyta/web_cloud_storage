package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
)

type Users struct {
	user_id              int
	foto                 []byte
	username             string
	userpass             string
	login                string
	working_email        string
	phone                int
	date_of_registration time.Time
}

func init() {
	orm.RegisterModel(new(Users))
}
