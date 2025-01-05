package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
)

type action_log struct {
	ActionlogID        int `orm:"column(action_log_id);pk;auto"`
	action_info        string
	file_in_storage_id int
	date_action        time.Time
}

func init() {
	orm.RegisterModel(new(action_log))
}
