package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_fasthttp"
	"net"
	"os"
	"time"
)

var (
	client    *zdpgo_fasthttp.Client
	targetUrl = `http://httpbin.org/get`
)

func main() {
	// 设置请求超时
	readTimeout, _ := time.ParseDuration("500ms")
	writeTimeout, _ := time.ParseDuration("500ms")
	maxIdleConnDuration, _ := time.ParseDuration("1h")

	// 创建客户端对象
	netAddr := &net.TCPAddr{Port: 9999}
	client = &zdpgo_fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
		DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
		DisablePathNormalizing:        true,
		// increase DNS cache time to an hour instead of default minute
		Dial: (&zdpgo_fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
			LocalAddr:        netAddr,
		}).Dial,
	}

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
}
