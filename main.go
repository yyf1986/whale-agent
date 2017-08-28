package main

import (
	_ "whale-agent/routers"

	"flag"
	"github.com/astaxie/beego"
)

func main() {
	port := flag.Int("port", 8080, "http port")
	ip := flag.String("ip", "", "docker host")
	flag.Parse()

	beego.BConfig.AppName = "whale-agent"
	beego.BConfig.RunMode = "dev"
	beego.BConfig.ServerName = "whale-agent"
	beego.BConfig.Listen.HTTPPort = *port
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	beego.AppConfig.Set("dockerhost",*ip)

	beego.Run()
}
