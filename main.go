package main

import (
	"whale-agent/models"
	_ "whale-agent/routers"

	"flag"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	port := flag.Int("port", 12345, "http port")
	ip := flag.String("ip", "10.11.20.127", "docker host")
	whaleserver := flag.String("whaleserver", "10.21.38.118:8080", "whale server host:port")
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

	toolbox.StartTask()
	models.Reg()
	beego.Run()

}
