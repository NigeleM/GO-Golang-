//arrays
package main

import "fmt"

func main() {
	var count int
	var datum string
	fmt.Printf("Enter a number : ")
	fmt.Scan(&count)
	var data = make([]string, count)

	for counter := 0; counter < count; counter++ {
		fmt.Printf("Enter datum : ")
		fmt.Scan(&datum)
		data[counter] = datum
	}

	for counter := 0; counter < count; counter++ {
		fmt.Println(data[counter])
	}

	for i, v := range data {
		fmt.Printf("index %v Value : %v \n", i, v)
	}

}
