package main

import (
	"fmt"
	"unicode/utf8"
)

func lenth(s string) int {
	for i, ch := range s {

	}
}

func main() {
	var str = "hello 你好"
	fmt.Println("len(str):", len(str))
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	fmt.Println("rune:", len([]rune(str)))
}
