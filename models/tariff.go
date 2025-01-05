package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type tariff struct {
	TariffID                 int `orm:"column(tariff_id);pk;auto"`
	tariff_name              string
	maximum_storage_capacity int
	price                    int
	description              string
	user_id                  int
}

func init() {
	orm.RegisterModel(new(tariff))
}
