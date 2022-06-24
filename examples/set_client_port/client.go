package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_fasthttp"
	"github.com/zhangdapeng520/zdpgo_log"
	"os"
)

var (
	targetUrl = `http://httpbin.org/get`
)

func main() {
	// 创建对象
	f := zdpgo_fasthttp.New(zdpgo_log.Tmp)

	for i := 0; i < 10000; i++ {
		// 获取客户端对象
		client := f.GetClient()

		// 发送GET请求
		req := zdpgo_fasthttp.AcquireRequest()
		req.SetRequestURI(targetUrl)
		req.Header.SetMethod(zdpgo_fasthttp.MethodGet)
		resp := zdpgo_fasthttp.AcquireResponse()
		err := client.Do(req, resp)
		zdpgo_fasthttp.ReleaseRequest(req)
		if err == nil {
			fmt.Printf("DEBUG Response: %s\n", resp.Body())
		} else {
			fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
		}
		zdpgo_fasthttp.ReleaseResponse(resp)

		// 查看客户端端口号
		fmt.Println("客户端端口号：", f.ClientPort)
	}
}
