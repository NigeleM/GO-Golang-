package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) talk() string {
	return "Hello"
}

func main() {

	man := person{name: "Nigele", age: 25}
	fmt.Println("My name is ", man.name, "and I'm ", man.age, man.talk(), "Everyone")
}
