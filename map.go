package main

import "fmt"

func length(s string) int {
	last0ccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := last0ccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		last0ccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(
		length("我爱慕课网abc ababababaaaaaa"))

	// x := make(map[string]int)
	// x["yang"] = 1
	// fmt.Println(x["yang"])
	// for k, v := range x {
	// 	fmt.Println(k, v)
	// }
	// delete(x, "yang")
	// if a, ok := x["yang"]; ok {
	// 	fmt.Println(a)
	// }
}
