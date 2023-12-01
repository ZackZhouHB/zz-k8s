package controllers

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.Data["json"] = "Hello, DevOps! 版权所有: 杜宽(Dotbalo), 获取课程和帮助: https://edu.51cto.com/lecturer/11062970.html"
	c.ServeJSON()
	return
}