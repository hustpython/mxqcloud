package models

import "fmt"

const WindowPC = "Windows PC"

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

var ServerSavePath = ""

func initServerSavePath() {
	var d Device
	err := ormOpr.QueryTable("device").Filter("is_server", true).One(&d)
	if err != nil {
		fmt.Println(err)
		return
	}
	ServerSavePath = d.SavePath
}

func (d *Device) GetDeviceByName() error {
	err := ormOpr.QueryTable("device").Filter("name", d.Name).One(d)
	return err
}

func GetAllUpload() ([]*Device, error) {
	var devices []*Device
	_, err := ormOpr.QueryTable("device").All(&devices)
	if err != nil {
		return nil, err
	}
	for i := range devices {
		_, err := ormOpr.LoadRelated(devices[i], "Uploads")
		if err != nil {
			return nil, err
		}
	}
	return devices, nil
}
func (d *Device) AddOrUpdateDevice() error {
	qs := ormOpr.QueryTable(new(Device))
	var err error
	if !qs.Filter("Name", d.Name).Exist() {
		_, err = ormOpr.Insert(d)
	} else {
		_, err = ormOpr.Update(d)
	}
	if err != nil {
		return err
	}
	if d.Name == WindowPC {
		ServerSavePath = d.SavePath
	}
	return nil
}
