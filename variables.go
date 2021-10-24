package main

import (
	"fmt"
)

func main() {

	sum := 2*3 + (5 + 1)
	var total int
	fmt.Println("Default Order ", sum)
	total = 2*4*9 + 1
	fmt.Println("total is ", total)
	var totalFloat float64 = float64(total)
	fmt.Printf("total as a float %.2f \n", totalFloat)

}
