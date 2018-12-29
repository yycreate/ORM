package wx

import (
	"ORM/util/httpTool"
)

func GetToken() interface{} {
	result := httpTool.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET")
	return result
}

func UploadMedia() interface{} {
	var params = map[string]string{}
	result := httpTool.Post("https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE", params)
	return result
}
