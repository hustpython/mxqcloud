package controllers

import (
	"mxqcloud-go/models"
	"os"
)

type DeviceController struct {
	ErrorController
}

func (md *DeviceController) AddOrUpdateDevice() {
	var d models.Device
	if err := md.BindJSON(&d); err != nil {
		md.setInternalError(err.Error())
		return
	}
	if len(models.ServerSavePath) == 0 && d.Name != models.WindowPC {
		md.setInternalError("serverPath in invalid")
		return
	}
	if err := d.AddOrUpdateDevice(); err != nil {
		md.setInternalError(err.Error())
		return
	}
	dirPath := models.ServerSavePath + "/" + d.Name

	_, err := os.Stat(dirPath)

	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dirPath, os.ModePerm)
		}
	}
	md.ServeJSON()
}

func (md *DeviceController) GetDevice() {
	name := md.GetString("Name")
	if len(name) == 0 {
		md.setParamInvalid("Name is null")
		return
	}
	d := models.Device{Name: name}
	err := d.GetDeviceByName()
	if err != nil {
		md.setInternalError(err.Error())
		return
	}
	md.Data["json"] = d
	md.ServeJSON()
}
