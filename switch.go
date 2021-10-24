package main

import "fmt"

func main() {

	var num int
	fmt.Printf("Enter a number : ")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println("number is equal to 1")
	case 2:
		fmt.Println("number is equal to 2")
	case 3:
		fmt.Println("number is equal to 2")
	default:
		fmt.Println("number is not in range")

	}
}
