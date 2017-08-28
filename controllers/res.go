package controllers

import (
	"github.com/astaxie/beego"
	"whale-agent/models"
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

