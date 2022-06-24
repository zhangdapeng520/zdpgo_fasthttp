package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_fasthttp"
)

/*
@Time : 2022/6/22 15:53
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	url := `http://httpbin.org/post?key=123`

	req := zdpgo_fasthttp.AcquireRequest()
	resp := zdpgo_fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		zdpgo_fasthttp.ReleaseResponse(resp)
		zdpgo_fasthttp.ReleaseRequest(req)
	}()

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)

	requestBody := []byte(`{"request":"test"}`)
	req.SetBody(requestBody)

	if err := zdpgo_fasthttp.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	b := resp.Body()

	fmt.Println("result:\r\n", string(b))
}
