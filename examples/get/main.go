package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_fasthttp"
)

/*
@Time : 2022/6/24 11:05
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	url := `http://httpbin.org/get`

	// 发送GET请求
	status, resp, err := zdpgo_fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	// 查看请求状态
	if status != zdpgo_fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return
	}
	fmt.Println(string(resp))
}
