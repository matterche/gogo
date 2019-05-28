package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	fmt.Println(reverseByteBuf("abc"))
	fmt.Println(reverseAppend("abc"))
}

func reverseByteBuf(s string) string {
	buf := bytes.Buffer{}
	for i := len(s) - 1; i >= 0; i-- {
		err := buf.WriteByte(s[i])
		if err != nil {
			log.Fatalf("Out of memory!")
		}
	}
	return buf.String()
}

func reverseAppend(s string) string {
	var buf []byte
	for i := len(s) - 1; i >= 0; i-- {
		buf = append(buf, s[i])
	}
	return string(buf)
}
