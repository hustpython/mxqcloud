package controllers

import (
	"mxqcloud-go/models"
	"path/filepath"
	"time"
)

type UploadController struct {
	ErrorController
}

// POST /v1/upload
func (m *UploadController) Upload() {
	device := m.Ctx.Request.Header.Get("Device")
	if len(device) < 1 {
		m.setInternalError("device invalid")
		return
	}
	f, h, _ := m.GetFile("file")
	defer f.Close()
	//m.SaveToFile("file", "D:\\1.png")
	if err := models.AddUploadData(&models.Upload{
		Id:       device + h.Filename,
		FileName: h.Filename,
		Device: &models.Device{
			Name: device,
		},
		FileType: filepath.Ext(h.Filename),
		Url:      "",
		Size:     1,
		Date:     time.Now(),
	}); err != nil {
		m.setInternalError(err.Error())
		return
	}
	m.Data["json"] = &models.Upload{
		Id:       device + h.Filename,
		FileName: h.Filename,
		Device: &models.Device{
			Name: device,
		},
		FileType: filepath.Ext(h.Filename),
		Url:      "",
		Size:     1,
		Date:     time.Now(),
	}
	m.ServeJSON()
}
