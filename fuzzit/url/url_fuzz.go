//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"bytes"
)

func Fuzz(data []byte) int {
	u := zdpgo_fasthttp.AcquireURI()
	defer zdpgo_fasthttp.ReleaseURI(u)

	u.UpdateBytes(data)

	w := bytes.Buffer{}
	if _, err := u.WriteTo(&w); err != nil {
		return 0
	}

	return 1
}
