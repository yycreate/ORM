package controllers

import (
	"test01/models"
	"test01/service"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

type JSONInsertObj struct {
	Success bool
	Id      int64
	User    models.User
}

var log *logs.BeeLogger

func init() {
	log = logs.NewLogger(10000)
	log.SetLogger("console", "")
	log.SetLevel(logs.LevelDebug)
}

//新增用户测试
func (c *UserController) InsertUser() {
	user := models.User{Username: "yy", Password: "pwd", Tel: "16512342232", Delflag: 1}
	id, user := service.InsertUser(user)
	json := JSONInsertObj{Success: true, Id: id, User: user}
	c.Data["json"] = json
	c.ServeJSON()
}

//查询用户
func (c *UserController) InfoUser() {
	user := models.User{Id: "7dcceef0-036a-11e9-8002-000000000000"}
	user = service.InfoUser(user)
	c.Data["json"] = user
	c.ServeJSON()
}

func (c *UserController) ListUser() {
	user := service.ListUser()
	c.Data["json"] = user
	c.ServeJSON()
}
