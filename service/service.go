package service

import (
	"fmt"
	"test01/models"
	"test01/util"

	"github.com/astaxie/beego/orm"
)

//插入数据
func InsertUser(user models.User) (int64, models.User) {
	orm.Debug = true
	o := orm.NewOrm()
	user.Id = util.TimeUUID().String()
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	return id, user
}

//查询单条用户记录
func InfoUser(user models.User) models.User {
	orm.Debug = true
	o := orm.NewOrm()
	err := o.Read(&user)
	if err != nil {
		fmt.Println("err:", err)
		return user
	}
	return user
}

//用户列表
func ListUser() []models.User {
	var users []models.User
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	_, err := qs.Filter("delflag", 1).All(&users)
	if err != nil {
		fmt.Println("err:", err)
		return users
	}
	return users
}
