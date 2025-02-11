package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
)

type payments struct {
	PaymentsID         int `orm:"column(payments_id);pk;auto"`
	amount             float64
	date_of_payment    time.Time
	status             bool
	transaction_number int32
	tariff_id          int
}

func init() {
	orm.RegisterModel(new(payments))
}
