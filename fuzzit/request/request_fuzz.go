//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"bufio"
	"bytes"
)

func Fuzz(data []byte) int {
	req := zdpgo_fasthttp.AcquireRequest()
	defer zdpgo_fasthttp.ReleaseRequest(req)

	if err := req.ReadLimitBody(bufio.NewReader(bytes.NewReader(data)), 1024*1024); err != nil {
		return 0
	}

	w := bytes.Buffer{}
	if _, err := req.WriteTo(&w); err != nil {
		return 0
	}

	return 1
}
