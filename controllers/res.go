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

func (d *ResController) Get2() {
	a, b, c := models.GetRes()
	d.Data["json"] = a + b + c
	d.ServeJSON()
}
