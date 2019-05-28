package main

import (
	"fmt"
	"rstring"
	"sort"
)

// must be run like GOPATH=`pwd` go run rstring-struct.go
func main() {
	rs := rstring.NewRString("hello world")
	rs1 := rstring.NewRString("a")
	rs2 := rstring.NewRString("ab")
	rs3 := rstring.NewRString("abc")

	ary := rstring.SortableRString{rs, rs1, rs3, rs2}
	fmt.Printf("%v\n", ary)
	sort.Sort(ary)
	fmt.Printf("%v\n", ary)
}
