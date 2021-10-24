package main

import "fmt"

func main() {

	var num int
	fmt.Printf("enter a number : ")
	fmt.Scan(&num)
	if 5 > num {
		fmt.Printf("number  %d is less than 5 \n", num)
	} else {
		fmt.Printf("number  %d is greater than 5 \n", num)
	}
}
