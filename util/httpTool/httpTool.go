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

func Post(reqUrl string, params map[string]string) interface{} {
	client := &http.Client{}
	mjson, _ := json.Marshal(params)
	paramsStr := string(mjson)
	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(paramsStr))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return nil
	}
	return string(body)
}
