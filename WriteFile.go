package main

import (
	"bufio"
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

	fmt.Print("Enter data : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	str += "\n"
	err := ioutil.WriteFile("Root2.txt", []byte(str), 0777)
	check(err)

	file, err := os.OpenFile("Root2.txt", os.O_APPEND|os.O_WRONLY, 0777)
	check(err)
	defer file.Close()

	fmt.Print("Enter data : ")
	scanner.Scan()
	//fmt.Println(scanner.Text())
	content := scanner.Text()
	fmt.Println(content)
	nb, err := file.Write([]byte(content))
	check(err)
	fmt.Println("Appended - > ", nb)

}
