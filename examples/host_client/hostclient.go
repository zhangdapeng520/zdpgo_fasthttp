package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_fasthttp"
	"os"
)

func main() {
	// Get URI from a pool
	url := zdpgo_fasthttp.AcquireURI()
	url.Parse(nil, []byte("http://localhost:8080"))
	url.SetUsername("Aladdin")
	url.SetPassword("Open Sesame")

	hc := &zdpgo_fasthttp.HostClient{
		Addr: "localhost:8080", // The host address and port must be set explicitly
	}

	req := zdpgo_fasthttp.AcquireRequest()
	req.SetURI(url)                // copy url into request
	zdpgo_fasthttp.ReleaseURI(url) // now you may release the URI
	req.Header.SetMethod(zdpgo_fasthttp.MethodGet)
	resp := zdpgo_fasthttp.AcquireResponse()

	// 发送100万次请求
	var err error
	err = hc.Do(req, resp)
	zdpgo_fasthttp.ReleaseRequest(req)
	if err == nil {
		fmt.Printf("Response: %s\n", resp.Body())
	} else {
		fmt.Fprintf(os.Stderr, "Connection error: %v\n", err)
	}
	zdpgo_fasthttp.ReleaseResponse(resp)
}
