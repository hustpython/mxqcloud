package main

import (
	_ "mxqcloud-go/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.BConfig.MaxUploadSize = 1 << 40
	beego.BConfig.MaxMemory = 1 << 30
	beego.Run()
}
