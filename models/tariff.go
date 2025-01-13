package models

import (
	"github.com/astaxie/beego/orm"
)

type Tariff struct {
	TariffID               int    `orm:"column(tariff_id); pk"`
	TariffName             string `orm:"column(tariff_name); size(50)"`
	MaximumStorageCapacity int    `orm:"column(maximum_storage_capacity)"`
	Price                  int    `orm:"column(price)"`
	Description            string `orm:"column(description); type(text)"`
}

func init() {
	orm.RegisterModel(new(Tariff))
}
