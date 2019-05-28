package main

import "fmt"

func main() {
	s := "中华人民共和国"
	for _, v := range s {
		fmt.Printf("%[1]c %[1]d\n", v)
	}
	fmt.Println(len(s))
	a := []rune(s)
	fmt.Println(len(a))

	str := "a;b;c"

}
