package main

import (
	_ "ORM/routers"

	"ORM/models"
	"ORM/util/redisTool"
	_ "fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	dbInit()              //数据库初始化
	redisTool.RedisLink() //初始化redis的链接
	beego.Run()
}

func dbInit() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:YYMySql2012@/beeadmin?charset=utf8", 30)
	// register model
	orm.RegisterModel(
		new(models.User),
		new(models.Role),
		new(models.RoleUserRelation))
	// create table
	orm.RunSyncdb("default", false, true)
}
