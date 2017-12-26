package controllers

import (
	"github.com/astaxie/beego"

	"net/http"
	"io/ioutil"
)

type AddController struct {
	beego.Controller
}

func(c *AddController)Get(){
	c.TplName = "add.html"
}


func(c *AddController)Post(){
	input := c.Input()
	r, err := http.PostForm("http://127.0.0.1:8080/request", input)
	if err != nil{
		beego.Error(err.Error())
	}else{
		beego.Info(ioutil.ReadAll(r.Body))
		r.Body.Close()
	}
	c.TplName = "add.html"
}
