package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_fasthttp"
	"log"
)

func main() {
	addr := ":8080"
	compress := false

	// 设置请求处理器
	h := requestHandler
	if compress {
		h = zdpgo_fasthttp.CompressHandler(h)
	}

	// 启动服务
	if err := zdpgo_fasthttp.ListenAndServe(addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}

func requestHandler(ctx *zdpgo_fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")
	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())
	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)
	ctx.SetContentType("text/plain; charset=utf8")

	// Set arbitrary headers
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// Set cookies
	var c zdpgo_fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)
}
