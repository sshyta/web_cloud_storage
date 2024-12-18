package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
)

type file_in_storage struct {
	file_in_storage_id int
	users_id           int
	file_size          int
	upload_date        time.Time
	file_type          string
	file_version       int
}

func init() {
	orm.RegisterModel(new(file_in_storage))
}
