// GoRoutines
package main

import (
	"fmt"
	"time"
)

func count(item string) {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, "  ", item)
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}

func main() {
	fmt.Println("Code ---> ==== <--- ")
	go count("Man")
	go count("Ant")
	count("Data")

}
