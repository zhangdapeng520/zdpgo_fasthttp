package main

import (
	"github.com/zhangdapeng520/zdpgo_fasthttp"
	"github.com/zhangdapeng520/zdpgo_log"
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

	f := zdpgo_fasthttp.New(zdpgo_log.Tmp)

	// 发送GET请求
	f.Get(url)
}
