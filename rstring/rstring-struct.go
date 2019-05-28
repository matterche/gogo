package main

import (
	"fmt"
	"rstring"
)

// must be run like GOPATH=`pwd` go run rstring-struct.go
func main() {
	rs := rstring.NewRString("hello world")
	fmt.Println(rs.Reverse())
}
