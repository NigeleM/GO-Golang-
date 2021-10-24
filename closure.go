package main

import "fmt"

func area(len int, wid int) int {
	//var area = len * wid
	return len * wid
}

func volume() func(num, n int) int {
	area := func(length, width int) int {
		return length * width
	}
	return area

}

func main() {

	vol := volume()
	fmt.Printf("Area Type : %T \n", area)
	fmt.Println("Area : ", area(5, 5))
	fmt.Println(" : ", area(5, 5))
	fmt.Println(" : ", area(5, 5)*10)
	fmt.Println(" : ", vol(10, 40)*8)

}
