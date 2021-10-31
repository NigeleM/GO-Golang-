package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	txt, err := ioutil.ReadFile("root.txt")
	check(err)
	fmt.Println(string(txt))

	file, err := os.Open("root.txt")
	check(err)
	defer file.Close()

	pos, err := file.Seek(5, 0)
	check(err)
	slice := make([]byte, 5)
	nb, err := file.Read(slice)
	check(err)
	fmt.Println(nb, pos)

}
