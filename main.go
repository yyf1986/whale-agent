package main

import (
	"whale-agent/models"
	_ "whale-agent/routers"

	"flag"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/astaxie/beego/context"

)

var FilterIp = func(ctx *context.Context) {
	isIn := false
	allowall := beego.AppConfig.String("allowall")
	if allowall == "N" {
		clientip := ctx.Input.IP()
		allow_ip := strings.Split(beego.AppConfig.String("whaleserver"),":")[0]
		if clientip == allow_ip {
			isIn = true
		}
		if !isIn {
			ctx.Redirect(403, "")
		}
	}
}
func main() {
	port := flag.Int("port", 12345, "http port")
	ip := flag.String("ip", "10.11.20.127", "docker host")
	whaleserver := flag.String("whaleserver", "10.21.38.118:8080", "whale server host:port")
	allowall := flag.String("allowall", "N", "是否允许whaleserver以为的机器访问")
	flag.Parse()
	//默认配置
	beego.BConfig.AppName = "whale-agent"
	beego.BConfig.RunMode = "prod"
	beego.BConfig.ServerName = "whale-agent"
	beego.BConfig.Listen.HTTPPort = *port
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	//启动输入参数，设置为配置
	p := strconv.Itoa(*port)
	beego.AppConfig.Set("port", p)
	beego.AppConfig.Set("dockerhost", *ip)
	beego.AppConfig.Set("whaleserver", *whaleserver)
	beego.AppConfig.Set("allowall", *allowall)

	beego.InsertFilter("/v1/*", beego.BeforeExec, FilterIp)

	toolbox.StartTask()
	models.Reg()
	beego.Run()

}
