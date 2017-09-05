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
	Url          map[string]string
}

func GetAgentinfo() AgentInfo {
	memtotal, memavailable, cpuTotal := GetRes()
	var ai AgentInfo
	port := beego.AppConfig.String("port")
	p, _ := strconv.Atoi(port)
	url := map[string]string{"CreateContainer":"/v1/container/create",
								"DelContainer":"/v1/container/del",
								"StartContainer":"/v1/container/start",
								"StopContainer":"/v1/container/stop",
								"GetAllContainers":"/v1/container/getall",
								"GetImages":"/v1/image/getall",
								"CreatePort":"/v1/res/createport",
								"DelPort":"/v1/res/delport",
								"GetAllPorts":"/v1/res/getallports",
								"DelAllPorts":"/v1/res/delallports"}
	ai = AgentInfo{beego.AppConfig.String("dockerhost"), p,
					cpuTotal, memtotal,
					cpuTotal, memavailable, "Y",url}
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
