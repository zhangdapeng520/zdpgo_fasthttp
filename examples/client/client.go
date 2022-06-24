package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_fasthttp"
	"io"
	"net/http"
	"os"
	"reflect"
	"time"
)

var headerContentTypeJson = []byte("application/json")

var client *zdpgo_fasthttp.Client

type Entity struct {
	Id   int
	Name string
}

func main() {
	// 设置请求超时
	readTimeout, _ := time.ParseDuration("500ms")
	writeTimeout, _ := time.ParseDuration("500ms")
	maxIdleConnDuration, _ := time.ParseDuration("1h")

	// 创建客户端对象
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
		}).Dial,
	}

	// 发送GET请求
	sendGetRequest()

	// 发送POST请求
	sendPostRequest()
}

func sendGetRequest() {
	req := zdpgo_fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/")
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

func sendPostRequest() {
	// per-request timeout
	reqTimeout := time.Duration(100) * time.Millisecond

	reqEntity := &Entity{
		Name: "New entity",
	}
	reqEntityBytes, _ := json.Marshal(reqEntity)

	req := zdpgo_fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/")
	req.Header.SetMethod(zdpgo_fasthttp.MethodPost)
	req.Header.SetContentTypeBytes(headerContentTypeJson)
	req.SetBodyRaw(reqEntityBytes)
	resp := zdpgo_fasthttp.AcquireResponse()
	err := client.DoTimeout(req, resp, reqTimeout)
	zdpgo_fasthttp.ReleaseRequest(req)
	if err == nil {
		statusCode := resp.StatusCode()
		respBody := resp.Body()
		fmt.Printf("DEBUG Response: %s\n", respBody)
		if statusCode == http.StatusOK {
			respEntity := &Entity{}
			err = json.Unmarshal(respBody, respEntity)
			if err == io.EOF || err == nil {
				fmt.Printf("DEBUG Parsed Response: %v\n", respEntity)
			} else {
				fmt.Fprintf(os.Stderr, "ERR failed to parse reponse: %v\n", err)
			}
		} else {
			fmt.Fprintf(os.Stderr, "ERR invalid HTTP response code: %d\n", statusCode)
		}
	} else {
		errName, known := httpConnError(err)
		if known {
			fmt.Fprintf(os.Stderr, "WARN conn error: %v\n", errName)
		} else {
			fmt.Fprintf(os.Stderr, "ERR conn failure: %v %v\n", errName, err)
		}
	}
	zdpgo_fasthttp.ReleaseResponse(resp)
}

func httpConnError(err error) (string, bool) {
	errName := ""
	known := false
	if err == zdpgo_fasthttp.ErrTimeout {
		errName = "timeout"
		known = true
	} else if err == zdpgo_fasthttp.ErrNoFreeConns {
		errName = "conn_limit"
		known = true
	} else if err == zdpgo_fasthttp.ErrConnectionClosed {
		errName = "conn_close"
		known = true
	} else {
		errName = reflect.TypeOf(err).String()
		if errName == "*net.OpError" {
			// Write and Read errors are not so often and in fact they just mean timeout problems
			errName = "timeout"
			known = true
		}
	}
	return errName, known
}
