package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       string `json:"id" orm:"column(id);pk"` // 设置主键
	Username string `json:"username" orm:"colunm(username);not null"`
	Password string `json:"password" orm:"colunm(password);not null"`
	Tel      string `json:"tel" orm:"column(tel)"`
	Delflag  int    `json:"delflag" orm:"column(delflag);default(1)"`
}

type Role struct {
	Id       string `json:"id" orm:"column(id);pk"` // 设置主键
	RoleName string `json:"roleName" orm:"colunm(roleName);not null"`
	RoleType string `json:"roleType" orm:"column(roleType)"`
	Delflag  int    `json:"delflag" orm:"column(delflag);default(1)"`
}

type RoleUserRelation struct {
	Id      string `json:"id" orm:"column(id);pk"` // 设置主键
	UserId  string `json:"userId" orm:"column(userId)"`
	RoleId  string `json:"roleId" orm:"column(roleId)"`
	Delflag int    `json:"delflag" orm:"column(delflag);default(1)"`
}
