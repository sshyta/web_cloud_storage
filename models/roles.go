package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type roles struct {
	RolesID      int `orm:"column(roles_id);pk;auto"`
	type_of_role string
}

func init() {
	orm.RegisterModel(new(roles))
}
