package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic("error occurred...")
	}
	defer resp.Body.Close()
	bytes, e := httputil.DumpResponse(resp, true)
	if e != nil {
		panic("error...")
	}
	fmt.Printf("%s\n", bytes)
}
