package zdpgo_fasthttp

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"testing"
)

/*
@Time : 2022/6/24 16:10
@Author : 张大鹏
@File : fast_http_test.go
@Software: Goland2021.3.1
@Description:
*/

func TestFastHttp_GetPoolPort(t *testing.T) {
	f := New(zdpgo_log.Tmp)
	for i := 0; i < 10000; i++ {
		// 出队
		port := f.GetPoolPort(0)
		if port < 0 {
			panic("错误的端口号")
		}

		// 出队
		poolPort := f.GetPoolPort(port)
		if poolPort == port || poolPort < 0 {
			panic("错误的端口号")
		}
	}
}
