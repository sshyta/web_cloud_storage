package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type access_to_files struct {
	AccesToFileID   int    `orm:"column(acces_to_file_id);pk;auto"` // Указание первичного ключа
	FileInStorageID int    `orm:"column(file_in_storage_id);null"`  // Внешний ключ, если необходимо
	AccessType      string `orm:"column(access_type);size(50);null"`
}

func init() {
	// Регистрация модели без префикса
	orm.RegisterModel(new(access_to_files))
}
