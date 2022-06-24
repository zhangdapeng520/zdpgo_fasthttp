package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_fasthttp"
)

/*
@Time : 2022/6/24 11:11
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	url := `http://httpbin.org/post?key=123`

	// 创建请求对象
	req := &zdpgo_fasthttp.Request{}
	req.SetRequestURI(url)
	requestBody := []byte(`{"request":"test"}`)
	req.SetBody(requestBody)

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	// 创建响应对象
	resp := &zdpgo_fasthttp.Response{}

	// 创建客户端
	client := &zdpgo_fasthttp.Client{}

	// 发送请求
	if err := client.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	// 请求体
	b := resp.Body()
	fmt.Println("result:\r\n", string(b))
}
