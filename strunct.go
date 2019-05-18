package main

import "fmt"

type Student struct {
	RollNumber int
	Name       string
}

func main() {
	s := Student{11, "Jack"}
	ps := &s
	psnew := new(Student)
	fmt.Println(ps)
	fmt.Println((*ps).Name)
	fmt.Println(psnew)
	a := "test"
	p := &a
	fmt.Println(p)
	fmt.Printf("Whatever.GetNameByReference() memory address: %p", ps)

}
