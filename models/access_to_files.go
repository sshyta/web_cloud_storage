package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type access_to_files struct {
	acces_to_file_id   int
	file_in_storage_id int
	access_type        string
}

func init() {
	orm.RegisterModel(new(access_to_files))
}
