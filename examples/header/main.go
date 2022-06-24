package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_fasthttp"
)

/*
@Time : 2022/6/22 15:52
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	url := `http://httpbin.org/post?key=123`

	req := &zdpgo_fasthttp.Request{}
	req.SetRequestURI(url)

	requestBody := []byte(`{"request":"test"}`)
	req.SetBody(requestBody)

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := &zdpgo_fasthttp.Response{}

	client := &zdpgo_fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	b := resp.Body()

	fmt.Println("result:\r\n", string(b))
}
