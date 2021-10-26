// buffer

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Enter Text : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	} else {
		fmt.Println(scanner.Text())
	}

}
