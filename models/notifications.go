package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
)

type notifications struct {
	NotificationsID      int `orm:"column(notifications_id);pk;auto"`
	notifications_text   string
	date_of_notification time.Time
	status               bool
	file_in_storage_id   int
}

func init() {
	orm.RegisterModel(new(notifications))
}
