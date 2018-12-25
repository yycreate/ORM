package ftpTool

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/dutchcoders/goftp"
)

var (
	ftp_ip       string
	ftp_port     string
	ftp_user     string
	ftp_password string
)

func InitFtp() {
	var err error
	var ftp *goftp.FTP

	ftp_ip = beego.AppConfig.String("ftp.ip")
	ftp_port = beego.AppConfig.String("ftp.port")
	ftp_user = beego.AppConfig.String("ftp.user")
	ftp_password = beego.AppConfig.String("ftp.password")

	// For debug messages: goftp.ConnectDbg("ftp.server.com:21")
	if ftp, err = goftp.Connect(ftp_ip + ":" + ftp_port); err != nil {
		panic(err)
	}

	defer ftp.Close()
	fmt.Println("Successfully connected !!")

	// Username / password authentication
	if err = ftp.Login(ftp_user, ftp_password); err != nil {
		panic(err)
	}

	if err = ftp.Cwd("/home/ftp"); err != nil {
		panic(err)
	}

	var curpath string
	if curpath, err = ftp.Pwd(); err != nil {
		panic(err)
	}

	fmt.Printf("Current path: %s", curpath)

	// Get directory listing
	var files []string
	if files, err = ftp.List(""); err != nil {
		panic(err)
	}
	fmt.Println("Directory listing:/n", files)

	// Upload a file
	var file *os.File
	if file, err = os.Open("E://6楼花名册.xlsx"); err != nil {
		panic(err)
	}

	if err := ftp.Stor("/home/ftp/6楼花名册.xlsx", file); err != nil {
		panic(err)
	}

}
