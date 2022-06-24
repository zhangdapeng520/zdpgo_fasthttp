package zdpgo_fasthttp

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"net"
)

/*
@Time : 2022/6/24 15:45
@Author : 张大鹏
@File : fast_http.go
@Software: Goland2021.3.1
@Description: 端口相关
*/

type FastHttp struct {
	Log *zdpgo_log.Log
}

func (f *FastHttp) GetHttpPort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		f.Log.Error("解析TCP地址失败", "error", err)
		return 0
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		f.Log.Error("创建tcp监听失败", "error", err)
		return 0
	}
	defer l.Close()

	// 获取端口号
	p := l.Addr().(*net.TCPAddr).Port
	return p
}
