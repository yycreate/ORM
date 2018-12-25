package form

type User struct {
	Id       string `json:"id" orm:"column(id);pk"` // 设置主键
	Username string `json:"username" orm:"colunm(username);not null"`
	Password string `json:"password" orm:"colunm(password);not null"`
	Tel      string `json:"tel" orm:"column(tel)"`
	Delflag  int    `json:"delflag" orm:"column(delflag);default(1)"`
}
