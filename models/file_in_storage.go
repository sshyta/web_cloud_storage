package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type FileInStorage struct {
	FileID      int       `orm:"column(file_in_storage_id);pk;auto"`
	UserID      int       `orm:"column(users_id)"`
	FileSize    int       `orm:"column(file_size)"`
	UploadDate  time.Time `orm:"column(upload_date);type(timestamp)"`
	FileType    string    `orm:"column(file_type);size(30)"`
	FileVersion int       `orm:"column(file_version)"`
}

func init() {
	orm.RegisterModel(new(FileInStorage))
}
