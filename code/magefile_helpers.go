//+build mage

package main

import (
	"runtime"
	"strings"
)

func requireGoVersion() error {
	if !strings.HasPrefix(runtime.Version(), "go1.16") {
		return error.New("requires go version 1.6")
	}
	return nil
}
