package main

import "fmt"

func main() {

	counter := 0
	limit := 0
	fmt.Printf("Enter a limit ")
	fmt.Scan(&limit)

	for counter <= limit {
		fmt.Println(counter, "==__==", limit)
		counter++
	}

}
