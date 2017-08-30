package controllers

import (
	"github.com/astaxie/beego"
	"whale-agent/models"
)

type ImageController struct {
	beego.Controller
}

func (d *ImageController) GetImages() {
	var images []string
	images = models.GetImages()
	d.Data["json"] = images
	d.ServeJSON()
}