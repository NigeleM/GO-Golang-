package main

import "fmt"

func recurs(num int, number int) int {

	if num < 0 {
		return number
	} else {
		number += num
		num--
		recurs(num, number)
	}
	return recurs(num, number)
}

func main() {

	var num, number int
	fmt.Printf("Enter a number : ")
	fmt.Scan(&num)
	fmt.Println("Result ", recurs(num, number))
}
