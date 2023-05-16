package models

// Device 表，用于前端交互和数据存储
type Device struct {
	Name      string `orm:"pk"`
	IsPublic  bool   // 是否共享
	IsServer  bool   // 仅PC作为服务端
	SavePath  string // 仅作为服务端时取值有效
	MaxSize   int32
	Uploads   []*Upload   `orm:"reverse(many)"`
	Downloads []*Download `orm:"reverse(many)"`
}
