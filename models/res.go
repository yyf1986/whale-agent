package models

import (
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

func GetRes() map[string]interface{} {
	res := make(map[string]interface{})
	MemTotal, _ := exec.Command("/bin/bash", "-c", `cat /proc/meminfo |grep MemTotal |awk '{print $2/1024}'`).Output()
	res["MemTotal"] = strings.Replace(string(MemTotal), "\n", "", -1) + "M"

	MemAvailable, _ := exec.Command("/bin/bash", "-c", `cat /proc/meminfo |grep MemAvailable |awk '{print $2/1024}'`).Output()
	res["MemAvailable"] = strings.Replace(string(MemAvailable), "\n", "", -1) + "M"

	CpuTotal, _ := exec.Command("/bin/bash", "-c", `cat /proc/cpuinfo | grep pr |wc -l`).Output()
	res["CpuTotal"] = strings.Replace(string(CpuTotal), "\n", "", -1)

	return res
}

func setCpu(do string, num float64) {
	var newnum float64
	if IsExistInFie("cpu") {
		old := GetCacheFromFile("cpu")
		if do == "sub" {
			newnum = old.(float64) - num
			setCache2File("cpu", newnum, 0)
		} else if do == "add" {
			newnum = old.(float64) + num
			setCache2File("cpu", newnum, 0)
		}
	} else {
		setCache2File("cpu", num, 0)
	}
}

func setMem(do string, num int) {
	var newnum int
	if IsExistInFie("mem") {
		old := GetCacheFromFile("mem")
		if do == "sub" {
			newnum = old.(int) - num
			setCache2File("mem", newnum, 0)
		} else if do == "add" {
			newnum = old.(int) + num
			setCache2File("mem", newnum, 0)
		}
	} else {
		setCache2File("mem", num, 0)
	}
}

//生成一个大于45000小于50000的数
func randNum() int {
	min := 45000
	//max := 65535

SUIJI:
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(50000)
	if num < min {
		goto SUIJI
	}
	return num
}

func setPort(do string, port int) {
	var newPorts []int
	if IsExistInFie("port") {
		oldPorts := GetCacheFromFile("port")
		if do == "sub" {
			for _, p := range oldPorts.([]int) {
				if p != port {
					newPorts = append(newPorts, p)
				}
			}
			setCache2File("port", newPorts, 0)
		} else if do == "add" {
			newPorts = append(oldPorts.([]int), port)
			setCache2File("port", newPorts, 0)
		}
	} else {
		newPorts = []int{port}
		setCache2File("port", newPorts, 0)
	}
}
