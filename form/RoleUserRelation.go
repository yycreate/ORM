package form

type RoleUserRelation struct {
	Id      string `json:"id" orm:"column(id);pk"` // 设置主键
	UserId  string `json:"userId" orm:"column(userId)"`
	RoleId  string `json:"roleId" orm:"column(roleId)"`
	Delflag int    `json:"delflag" orm:"column(delflag);default(1)"`
}
