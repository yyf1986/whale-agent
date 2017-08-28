package models

import (
	"math/rand"
	"os/exec"
	"strings"
	"time"
	"strconv"
)
//内存换算的时候去掉了小数点，内存单位为M
func GetRes() (int,int,int) {
	MemTotal, _ := exec.Command("/bin/bash", "-c", `cat /proc/meminfo |grep MemTotal |awk '{print $2/1024}'`).Output()
	memtotal,_  := strconv.Atoi(strings.Split(strings.Replace(string(MemTotal), ".*\n", "", -1),".")[0])

	MemAvailable, _ := exec.Command("/bin/bash", "-c", `cat /proc/meminfo |grep MemAvailable |awk '{print $2/1024}'`).Output()
	memavailable,_ := strconv.Atoi(strings.Split(strings.Replace(string(MemAvailable), ".*\n", "", -1),".")[0])

	CpuTotal, _ := exec.Command("/bin/bash", "-c", `cat /proc/cpuinfo | grep pr |wc -l`).Output()
	cpuTotal,_ := strconv.Atoi(strings.Replace(string(CpuTotal), "\n", "", -1))

	return memtotal, memavailable, cpuTotal
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
//生成一个45000-50000的端口，端口必须不是已经在使用的
func GetPort() int {
pp:
	newport := randNum()
	if ! checkPort(newport) {
		goto pp
	}
	setPort("add", newport)
	return newport
}

func DelPort(port int) {
	if ! checkPort(port) {
		setPort("sub", port)
	}
}
//不存在为true，存在为false
func checkPort(port int) bool {
	isnew := true
	if IsExistInFie("ResPort") {
		ports := GetCacheFromFile("ResPort")
		for _, p := range ports.([]int) {
			if p == port {
				isnew = false
			}
		}
		if ! isnew {
			return false
		} else {
			return true
		}
	} else {
		return true
	}
}
func setPort(do string, port int) {
	var newPorts []int
	if IsExistInFie("ResPort") {
		oldPorts := GetCacheFromFile("ResPort")
		if do == "sub" {
			for _, p := range oldPorts.([]int) {
				if p != port {
					newPorts = append(newPorts, p)
				}
			}
			setCache2File("ResPort", newPorts, 0)
		} else if do == "add" {
			newPorts = append(oldPorts.([]int), port)
			setCache2File("ResPort", newPorts, 0)
		}
	} else {
		newPorts = []int{port}
		setCache2File("ResPort", newPorts, 0)
	}
}
