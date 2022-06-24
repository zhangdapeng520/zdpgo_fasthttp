package pprofhandler

import (
	"github.com/zhangdapeng520/zdpgo_fasthttp"
	zdpgo_fasthttpadaptor "github.com/zhangdapeng520/zdpgo_fasthttp/fasthttpadaptor"
	"net/http/pprof"
	rtp "runtime/pprof"
	"strings"
)

var (
	cmdline = zdpgo_fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline)
	profile = zdpgo_fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile)
	symbol  = zdpgo_fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol)
	trace   = zdpgo_fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace)
	index   = zdpgo_fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index)
)

// PprofHandler serves server runtime profiling data in the format expected by the pprof visualization tool.
//
// See https://golang.org/pkg/net/http/pprof/ for details.
func PprofHandler(ctx *zdpgo_fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "text/html")
	if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/cmdline") {
		cmdline(ctx)
	} else if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/profile") {
		profile(ctx)
	} else if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/symbol") {
		symbol(ctx)
	} else if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/trace") {
		trace(ctx)
	} else {
		for _, v := range rtp.Profiles() {
			ppName := v.Name()
			if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/"+ppName) {
				namedHandler := zdpgo_fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler(ppName).ServeHTTP)
				namedHandler(ctx)
				return
			}
		}
		index(ctx)
	}
}
