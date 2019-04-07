package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := "bool:"
	//convert string b to a byte array
	by := strconv.AppendBool([]byte(b), true)
	fmt.Println(string(b))
	//fmt.Println(byte(b))
	fmt.Println([]byte(b))

	br := []byte("rune:")
	br = strconv.AppendQuoteRune(br, '☺')
	fmt.Println(string(br))

}
