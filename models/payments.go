package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	"time"
)

type payments struct {
	payments_id        int
	amount             float64
	date_of_payment    time.Time
	status             bool
	transaction_number int32
	tariff_id          int
}

func init() {
	orm.RegisterModel(new(payments))
}
