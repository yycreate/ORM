package httpTool

import (
	"fmt"
	"net/http"
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
	return nil
}
