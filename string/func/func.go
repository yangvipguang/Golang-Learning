package main

import (
	"fmt"
	"math/rand"
)

func returnMulti() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func main() {
	a, b := returnMulti()
	fmt.Println(a, b)
}
