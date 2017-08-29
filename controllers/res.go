package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"whale-agent/models"
)

type ResController struct {
	beego.Controller
}

func (d *ResController) Get() {
	d.Data["json"] = models.GetAgentinfo()
	d.ServeJSON()
}

func (d *ResController) CreatePort() {
	d.Data["json"] = map[string]int{"port": models.CreatePort()}
	d.ServeJSON()
}

func (d *ResController) DelPort() {
	p := d.GetString("port")
	port, _ := strconv.Atoi(p)
	models.DelPort(port)
	d.Data["json"] = map[string]int{"status": 200}
	d.ServeJSON()
}

func (d *ResController) GetAllPorts() {
	d.Data["json"] = models.GetAllPorts()
	d.ServeJSON()
}

func (d *ResController) DelAllPorts() {
	models.DelAllPorts()
	d.Data["json"] = map[string]int{"status": 200}
	d.ServeJSON()
}
