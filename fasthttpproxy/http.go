package zdpgo_fasthttpproxy

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_fasthttp"
	"net"
	"strings"
	"time"
)

// FasthttpHTTPDialer returns azdpgo_fasthttp.DialFunc that dials using
// the provided HTTP proxy.
//
// Example usage:
//	c := &zdpgo_fasthttp.Client{
//		Dial: fasthttpproxy.FasthttpHTTPDialer("username:password@localhost:9050"),
//	}
func FasthttpHTTPDialer(proxy string) zdpgo_fasthttp.DialFunc {
	return FasthttpHTTPDialerTimeout(proxy, 0)
}

// FasthttpHTTPDialerTimeout returns azdpgo_fasthttp.DialFunc that dials using
// the provided HTTP proxy using the given timeout.
//
// Example usage:
//	c := &fasthttp.Client{
//		Dial: fasthttpproxy.FasthttpHTTPDialerTimeout("username:password@localhost:9050", time.Second * 2),
//	}
func FasthttpHTTPDialerTimeout(proxy string, timeout time.Duration) zdpgo_fasthttp.DialFunc {
	var auth string
	if strings.Contains(proxy, "@") {
		split := strings.Split(proxy, "@")
		auth = base64.StdEncoding.EncodeToString([]byte(split[0]))
		proxy = split[1]
	}

	return func(addr string) (net.Conn, error) {
		var conn net.Conn
		var err error
		if timeout == 0 {
			conn, err = zdpgo_fasthttp.Dial(proxy)
		} else {
			conn, err = zdpgo_fasthttp.DialTimeout(proxy, timeout)
		}
		if err != nil {
			return nil, err
		}

		req := fmt.Sprintf("CONNECT %s HTTP/1.1\r\nHost: %s\r\n", addr, addr)
		if auth != "" {
			req += "Proxy-Authorization: Basic " + auth + "\r\n"
		}
		req += "\r\n"

		if _, err := conn.Write([]byte(req)); err != nil {
			return nil, err
		}

		res := zdpgo_fasthttp.AcquireResponse()
		defer zdpgo_fasthttp.ReleaseResponse(res)

		res.SkipBody = true

		if err := res.Read(bufio.NewReader(conn)); err != nil {
			conn.Close()
			return nil, err
		}
		if res.Header.StatusCode() != 200 {
			conn.Close()
			return nil, fmt.Errorf("could not connect to proxy: %s status code: %d", proxy, res.Header.StatusCode())
		}
		return conn, nil
	}
}
