//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"bytes"
)

func Fuzz(data []byte) int {
	c := zdpgo_fasthttp.AcquireCookie()
	defer zdpgo_fasthttp.ReleaseCookie(c)

	if err := c.ParseBytes(data); err != nil {
		return 0
	}

	w := bytes.Buffer{}
	if _, err := c.WriteTo(&w); err != nil {
		return 0
	}

	return 1
}
