package controllers

import (
	"fmt"
	"mxqcloud-go/models"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type UploadController struct {
	ErrorController
}

// GET /v1/upload
func (m *UploadController) GetUpload() {
	name := m.GetString("Name")
	if len(name) == 0 {
		m.setParamInvalid("Name is null")
		return
	}
	us, err := models.GetUpload(name)
	if err != nil {
		m.setParamInvalid(err.Error())
		return
	}
	m.Data["json"] = us
	m.ServeJSON()
}

// POST /v1/upload
func (m *UploadController) Upload() {
	device := m.Ctx.Request.Header.Get("Device")
	if len(device) < 1 {
		m.setInternalError("device invalid")
		return
	}

	d := models.Device{Name: device}
	err := d.GetDeviceByName()
	if err != nil {
		m.setInternalError(err.Error())
		return
	}

	if len(models.ServerSavePath) == 0 {
		m.setInternalError("invalid savePath")
		return
	}

	f, h, _ := m.GetFile("file")
	defer f.Close()
	url := models.ServerSavePath + "/" + device + "/" + h.Filename
	m.SaveToFile("file", url)
	nowDate := time.Now()
	uploadData := &models.Upload{
		Id:       device + h.Filename,
		FileName: h.Filename,
		Device: &models.Device{
			Name: device,
		},
		FileType: filepath.Ext(h.Filename),
		Url:      strings.Replace(url, ":", "", -1),
		Size:     h.Size,
		SizeStr:  ByteSizeTransform(h.Size),
		DateStr:  nowDate.Format("2006-01-02 15:04"),
		Date:     nowDate,
	}
	if err := models.AddUploadData(uploadData); err != nil {
		m.setInternalError(err.Error())
		return
	}
	m.Data["json"] = uploadData
	m.ServeJSON()
}

// DELETE v1/upload
func (m *UploadController) DelUpload() {
	id := m.GetString("Id")
	if len(id) == 0 {
		m.setParamInvalid("Id is null")
		return
	}
	url, err := models.DeleteUpload(id)
	if err != nil {
		m.setInternalError(err.Error())
		return
	}
	url = url[:1] + ":" + url[1:]
	_, err = os.Stat(url)
	if err == nil {
		if err = os.Remove(url); err != nil {
			m.setInternalError(err.Error())
			return
		}
	}
	m.ServeJSON()
}

// 以1024作为基数
func ByteSizeTransform(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %c",
		float64(b)/float64(div), "KMGTPE"[exp])
}
