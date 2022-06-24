//go:build !windows
// +build !windows

package zdpgo_fasthttp

import (
	"errors"
	"syscall"
)

func isConnectionReset(err error) bool {
	return errors.Is(err, syscall.ECONNRESET)
}
