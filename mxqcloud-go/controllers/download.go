package controllers

import "mxqcloud-go/models"

type DownloadController struct {
	ErrorController
}

func (d *DownloadController) GetDownload() {
	ds, err := models.GetAllUpload()
	if err != nil {
		d.setInternalError(err.Error())
		return
	}
	d.Data["json"] = ds
	d.ServeJSON()
}
