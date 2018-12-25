package form

type Role struct {
	Id       string `json:"id" orm:"column(id);pk"` // 设置主键
	RoleName string `json:"roleName" orm:"colunm(roleName);not null"`
	RoleType string `json:"roleType" orm:"column(roleType)"`
	Delflag  int    `json:"delflag" orm:"column(delflag);default(1)"`
}
