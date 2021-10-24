// intefaces

package main

import "fmt"

type bird interface {
	speak() string
	move() string
}

type parrot struct{}

func (parrot) speak() string {
	return "Hello!"
}

func (parrot) move() string {
	return "A parrot flies!"
}

type swan struct{}

func (swan) speak() string {
	return "swan!"
}

func (swan) move() string {
	return "A swan flies!"
}

func nudge(b bird) {
	fmt.Println(b.speak())
	fmt.Println(b.move())
}

func main() {

	var birds swan
	var birdy parrot
	nudge(birds)
	nudge(birdy)

}
