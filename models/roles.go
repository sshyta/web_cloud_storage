package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type roles struct {
	roles_id     int
	type_of_role string
}

func init() {
	orm.RegisterModel(new(roles))
}
