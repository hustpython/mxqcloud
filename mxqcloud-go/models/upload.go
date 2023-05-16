package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

// Upload 表，用于前端交互和数据存储
type Upload struct {
	Id       string `orm:"pk"`
	FileName string
	Device   *Device `orm:"null;rel(fk)"`
	FileType string
	Url      string
	Size     int32
	Date     time.Time
}

var ormOpr orm.Ormer

func init() {
	logs.SetLogger(logs.AdapterConsole)
	orm.RegisterModel(new(Device), new(Upload), new(Download))
	orm.RegisterDataBase("default", "sqlite3", "./mxqcloud.db")
	orm.RunSyncdb("default", false, true)
	ormOpr = orm.NewOrm()
}

func AddUploadData(u *Upload) error {
	qs := ormOpr.QueryTable(new(Upload))
	if !qs.Filter("Id", u.Id).Exist() {
		_, e := ormOpr.Insert(u)
		fmt.Println("insert", u.Id)
		return e
	} else {
		return errors.New("Exist")
	}

}
