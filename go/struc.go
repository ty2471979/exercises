package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Ty", age: 45}
	fmt.Println(p)

	q := person{name: "Gwen", age: 60}
	fmt.Println(p, q)

}
