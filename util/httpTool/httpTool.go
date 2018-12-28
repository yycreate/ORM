package httpTool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	body := resp.Body
	defer resp.Body.Close()
	if err != nil {
		// handle error
	}
	fmt.Print(body)
	return ""
}

func Post(reqUrl string, params map[string]string) string {

	mjson, _ := json.Marshal(params)
	paramsStr := string(mjson)
	res, err := http.NewRequest("POST", reqUrl, strings.NewReader(paramsStr))
	defer res.Body.Close()

	fmt.Println(res.Body)
	fmt.Println(res.Context())
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(body)
	if err != nil {
		return ""
	}
	return ""
}
