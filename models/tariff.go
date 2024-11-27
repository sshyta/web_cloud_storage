package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
)

type tariff struct {
	tariff_id                int
	tariff_name              string
	maximum_storage_capacity int
	price                    int
	description              string
	user_id                  int
}

func init() {
	orm.RegisterModel(new(tariff))
}
