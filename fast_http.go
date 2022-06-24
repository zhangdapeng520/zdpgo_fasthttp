package zdpgo_fasthttp

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"net"
	"time"
)

/*
@Time : 2022/6/24 15:45
@Author : 张大鹏
@File : fast_http.go
@Software: Goland2021.3.1
@Description: 端口相关
*/

type FastHttp struct {
	Log         *zdpgo_log.Log
	Config      *Config
	ClientPorts []int // 客户端端口列表
	ClientPort  int   // 当前客户端端口
}

func New(log *zdpgo_log.Log) *FastHttp {
	return NewWithConfig(&Config{}, log)
}

func NewWithConfig(config *Config, log *zdpgo_log.Log) *FastHttp {
	f := &FastHttp{}

	// 配置
	if config.ReadTimeout == 0 {
		config.ReadTimeout = 33
	}
	if config.WriteTimeout == 0 {
		config.WriteTimeout = 33
	}
	if config.MaxIdleConnDuration == 0 {
		config.MaxIdleConnDuration = 60
	}
	f.Config = config

	// 日志
	f.Log = log

	// 返回
	return f
}

// Get 发送GET请求
func (f *FastHttp) Get(targetUrl string) {
	client := f.GetClient()

	// 设置请求对象
	req := AcquireRequest()
	defer ReleaseRequest(req)
	req.SetRequestURI(targetUrl)
	req.Header.SetMethod(MethodGet)

	// 设置响应对象
	resp := AcquireResponse()
	defer ReleaseResponse(resp)

	// 发送请求
	err := client.Do(req, resp)

	// 处理结果
	if err == nil {
		f.Log.Debug("发送请求成功", "body", string(resp.Body()))
	} else {
		f.Log.Error("发送请求失败", "error", err)
	}
}

// GetClient 获取客户端
func (f *FastHttp) GetClient() *Client {
	// 获取客户端端口
	port := f.GetPoolPort(f.ClientPort)
	f.ClientPort = port

	// 创建客户端
	netAddr := &net.TCPAddr{Port: port}
	client := &Client{
		ReadTimeout:                   time.Second * time.Duration(f.Config.ReadTimeout),
		WriteTimeout:                  time.Second * time.Duration(f.Config.WriteTimeout),
		MaxIdleConnDuration:           time.Second * time.Duration(f.Config.MaxIdleConnDuration),
		NoDefaultUserAgentHeader:      false, // 使用UserAgent
		DisableHeaderNamesNormalizing: true,  // 如果你在标题上正确设置了大小写，你可以启用它
		DisablePathNormalizing:        true,
		// 增加DNS缓存时间到一个小时，而不是默认的分钟
		Dial: (&TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
			LocalAddr:        netAddr,
		}).Dial,
	}
	return client
}

// GetHttpPort 获取可用的HTTP端口
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

// GetHttpPortSlice 获取可用端口的切片
func (f *FastHttp) GetHttpPortSlice(length int) []int {
	var data []int
	for i := 0; i < length; i++ {
		data = append(data, f.GetHttpPort())
	}
	return data
}

// GetPoolPort 从端口池中获取端口
func (f *FastHttp) GetPoolPort(oldPort int) int {
	if f.ClientPorts == nil || len(f.ClientPorts) == 0 {
		f.ClientPorts = f.GetHttpPortSlice(333)
	}
	length := len(f.ClientPorts) - 1
	if oldPort > 0 {
		f.ClientPorts = append([]int{oldPort}, f.ClientPorts[:length]...)
	} else {
		f.ClientPorts = append([]int{f.GetHttpPort()}, f.ClientPorts[:length]...)
	}
	return f.ClientPorts[length]
}
