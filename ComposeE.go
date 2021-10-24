//ComposeElements
package main

import "fmt"

type member struct {
	firstname string
	lastname  string
}

func (m member) fullname() string {
	return m.firstname + " " + m.lastname
}

type article struct {
	title string
	body  string
	member
}

func (a article) content() {
	fmt.Println("Title: ", a.title)
	fmt.Println("Content: ", a.body)
	fmt.Println("Author: ", a.fullname())
}

func main() {
	mems := member{"John", "Edwards"}
	arts := article{"Mango Book", " Health Benefits of Mangos", mems}

	arts.content()

}
