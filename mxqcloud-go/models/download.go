package models

import "time"

// Download 表，用于前端交互和数据存储
type Download struct {
	Id       string `orm:"pk"`
	FileName string
	Device   *Device `orm:"null;rel(fk)"`
	FileType string
	Url      string
	Size     int64
	Date     time.Time
}
