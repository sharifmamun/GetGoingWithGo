package main

import (
	"fmt"
)

type animal interface {
	speak()
}

type person struct {
	name  string
	email string
}

type cat struct {
	name string
}

func (p person) speak() {
	fmt.Println(p.name)
}

func (c cat) speak() {
	fmt.Println("meow")
}

func saysomething(a animal) {
	a.speak()
}

func main() {
	p := person{name: "Sharif", email: "sharif.mamun@xxxxxx.com"}
	saysomething(p)

	c := cat{name: "Billu"}
	saysomething(c)
}
