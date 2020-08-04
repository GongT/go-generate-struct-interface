//go:generate go run ../cmd/go-generate-struct-interface
// +build !windows

package main

import (
	fff "fmt"

	net "net"
	"runtime"
)

func init() {
	fff.Print(runtime.GOOS)
}

type someType struct {
	Sub string
}

type someInterface interface {
	Sub() string
}

type someStruct struct {
	U, U1, U2, u3  uint
	StrPtr         *string       `description:"bbbb"`
	OtherInterface someInterface `json:"-"`
	OtherStruct    someType
	OtherStructPtr *someType
	Channel        chan bool
	Array          []byte
	ArrayFixed     [4]byte
	Map            map[string]bool
	ComplexMap1    map[*someType]struct{ i int }
	ComplexMap2    map[struct{ i int }]*someType

	RemoteField net.IP

	priField      string
	priFieldNoGet string `getter:"-"`
}
