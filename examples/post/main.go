package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_fasthttp"
)

/*
@Time : 2022/6/22 15:50
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/
func main() {
	url := `http://httpbin.org/post?key=123`

	// 填充表单，类似于net/url
	args := &zdpgo_fasthttp.Args{}
	args.Add("name", "test")
	args.Add("age", "18")

	// 发送POST请求
	status, resp, err := zdpgo_fasthttp.Post(nil, url, args)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}
	if status != zdpgo_fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return
	}

	fmt.Println(string(resp))
}
