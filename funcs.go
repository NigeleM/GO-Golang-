// functions

package main

import "fmt"

func hello() {
	fmt.Println("Hello world")
}

func sqr(num int) int {

	sqr := num * num
	return sqr

}

func main() {

	var num, number int
	fmt.Printf("Enter a number : ")
	fmt.Scan(&num)
	hello()
	number = sqr(num)
	fmt.Println("Square of : ", num, "is ", number)

}
