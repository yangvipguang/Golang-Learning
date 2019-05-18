package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes 今天周四"
	fmt.Println(len(s))

	for _, v := range []byte(s) {
		fmt.Printf("%X ", v)
	}

	fmt.Println()
	for i, ch := range s { //ch is a rune  and is also a int32
		fmt.Printf("(%d %x)", i, ch)
	}
	fmt.Println()
	fmt.Printf("%s\n", []byte(s))
	fmt.Println("Rune count is: ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
}
