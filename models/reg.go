package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

type AgentInfo struct {
	IP           string
	Port         int
	TotalCpu     int
	TotalMem     int
	AverCpu      int
	AverMem      int
	DockerStatus string
}

func GetAgentinfo() AgentInfo {
	memtotal, memavailable, cpuTotal := GetRes()
	var ai AgentInfo
	port := beego.AppConfig.String("port")
	p, _ := strconv.Atoi(port)
	ai = AgentInfo{beego.AppConfig.String("dockerhost"), p,
		cpuTotal, memtotal,
		cpuTotal, memavailable, "Y"}
	return ai
}
func Reg() {
	req := httplib.Post("http://" + beego.AppConfig.String("whaleserver") + "/v1/reg")
	req.SetTimeout(1*time.Second, 2*time.Second)
	req.JSONBody(GetAgentinfo())
	_, err := req.String()
	if err != nil {
		logs.Error(err.Error())
	} else {
		logs.Info(beego.AppConfig.String("dockerhost") + " reg sucess")
	}
}

func addTask4Reg() {
	f := func() error { Reg(); return nil }
	tk := toolbox.NewTask("reg", "0 */1 * * * *", f)
	toolbox.AddTask("reg", tk)
}

func init() {
	addTask4Reg()
}
