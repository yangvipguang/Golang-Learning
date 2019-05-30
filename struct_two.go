package main

import "fmt"

type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e Employee) String() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)

}

func main() {
	e := Employee{"ji", "yang", 20}
	fmt.Println(e)
}
