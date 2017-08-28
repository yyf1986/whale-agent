package controllers

import (
	"github.com/astaxie/beego"
	"whale-agent/models"
	"strconv"
)

type ResController struct {
	beego.Controller
}

func (d *ResController) Get() {
	d.Data["json"] = models.GetAgentinfo()
	d.ServeJSON()
}

func (d *ResController) GetPort() {
	d.Data["json"] = map[string]int{"port": models.GetPort()}
	d.ServeJSON()
}

func (d *ResController) DelPort() {
	p := d.GetString("port")
	port, _ := strconv.Atoi(p)
	models.DelPort(port)
	d.Data["json"] = map[string]int{"status": 200}
	d.ServeJSON()
}

