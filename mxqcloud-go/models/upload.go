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
	Size     int64
	SizeStr  string
	Date     time.Time
	DateStr  string
	Check    bool `orm:"-"`
}

var ormOpr orm.Ormer

func init() {
	logs.SetLogger(logs.AdapterConsole)
	orm.RegisterModel(new(Device), new(Upload), new(Download))
	orm.RegisterDataBase("default", "sqlite3", "./mxqcloud.db")
	orm.RunSyncdb("default", false, true)
	ormOpr = orm.NewOrm()
	initServerSavePath()
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

func GetUpload(name string) ([]*Upload, error) {
	var device = Device{Name: name}
	_, err := ormOpr.LoadRelated(&device, "Uploads")
	if err != nil {
		return nil, err
	}
	return device.Uploads, nil
}

func DeleteUpload(id string) (string, error) {
	var u Upload
	err := ormOpr.QueryTable("upload").Filter("id", id).One(&u)
	if err != nil {
		return "", err
	}
	if _, e := ormOpr.Delete(&u); e != nil {
		return "", e
	}
	return u.Url, nil
}
