package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Roles struct {
	RolesID    int    `orm:"column(roles_id); pk; auto"`
	TypeOfRole string `orm:"column(type_of_role); size(128);"`
}

func (r *Roles) TableName() string {
	return "roles"
}

func init() {
	orm.RegisterModel(new(Roles))
}
