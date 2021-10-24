package main

import "fmt"

func main() {

	// Nested loops
	for counter := 1; counter <= 10; counter++ {
		for counts := 1; counts <= 10; counts++ {
			fmt.Println(counter, "==__==", counts)
		}
	}

}
