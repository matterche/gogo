// Package rstring only works when calling GOPATH=`pwd` go run example.go or move reverse into workspace
package rstring

import (
	"bytes"
	"log"
)

// RString represent reversable strings
type RString struct {
	S string
}

// NewRString constructs RString
func NewRString(s string) RString {
	return RString{S: s}
}

// Reverse reverses a RString
func (s RString) Reverse() RString {
	s.S = reverse(s.S)
	return s
}

func reverse(s string) string {
	buf := bytes.Buffer{}
	for i := len(s) - 1; i >= 0; i-- {
		err := buf.WriteByte(s[i])
		if err != nil {
			log.Fatalf("Out of memory!")
		}
	}
	return buf.String()
}

// String implements the stringer interface
func (s RString) String() string {
	return s.S
}

// SortableRString enables sort on slices of RString
type SortableRString []RString

// implement the sort interface
func (s SortableRString) Len() int {
	return len(s)
}
func (s SortableRString) Less(i, j int) bool {
	return s[i].S < s[j].S
}
func (s SortableRString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
