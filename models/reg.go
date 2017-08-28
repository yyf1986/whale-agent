package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	 "strconv"
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
	p,_ := strconv.Atoi(port)
	ai = AgentInfo{beego.AppConfig.String("dockerhost"), p,
		cpuTotal, memtotal,
		cpuTotal, memavailable, "Y"}
	fmt.Println(ai)
	fmt.Println(beego.AppConfig.String("whaleserver"))
	return ai
}
func reg() {
	req := httplib.Post("http://" + beego.AppConfig.String("whaleserver") + "/v1/reg")
	req.Body(GetAgentinfo())
	str, err := req.String()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(str)

	//fmt.Println(GetAgentinfo())
}
