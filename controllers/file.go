package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

func uploadFile(this *FileController) {
	f, h, a := this.GetFile("myfile") //获取上传的文件
	fmt.Println(f, h, a)
}
