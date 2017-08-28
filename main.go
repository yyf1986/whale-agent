package main

import (
	_ "whale-agent/routers"

	"flag"
	"github.com/astaxie/beego"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "http port")
	ip := flag.String("ip", "", "docker host")
	whaleserver := flag.String("whaleserver", "", "whale server host:port")
	flag.Parse()
	//默认配置
	beego.BConfig.AppName = "whale-agent"
	beego.BConfig.RunMode = "dev"
	beego.BConfig.ServerName = "whale-agent"
	beego.BConfig.Listen.HTTPPort = *port
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	//启动输入参数，设置为配置
	p := strconv.Itoa(*port)
	beego.AppConfig.Set("port", p)
	beego.AppConfig.Set("dockerhost", *ip)
	beego.AppConfig.Set("whaleserver", *whaleserver)

	beego.Run()
}
