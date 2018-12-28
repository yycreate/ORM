package main

import (
	_ "ORM/routers"

	"ORM/models"
	_ "ORM/util/ftpTool"
	"ORM/util/httpTool"
	"ORM/util/redisTool"
	_ "fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	dbInit()              //数据库初始化
	redisTool.RedisLink() //初始化redis的链接
	//	ftpTool.copyFile("e://s.txt", "/home/ftp")
	var params = map[string]string{}

	httpTool.Post("https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE", params)
	beego.Run()
}

func dbInit() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@/beeadmin?charset=utf8", 30)
	// register model
	orm.RegisterModel(
		new(models.User),
		new(models.Role),
		new(models.RoleUserRelation))
	// create table
	orm.RunSyncdb("default", false, true)
}
