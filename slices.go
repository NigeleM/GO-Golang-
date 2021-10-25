// Slices

package main

import "fmt"

func main() {

	array := [3]string{"Dad", "Mom", "Grandma"}
	slice := array[:]
	slice = append(slice, "Brother", "sister")
	array2 := [3]string{"Brat", "baddy", "jock"}
	slice2 := array2[:]
	slice2 = append(slice2, "Bro", "sis")
	list(slice...)
	list(slice2...)
}

func list(family ...string) {
	for i, v := range family {
		fmt.Printf("family member %v : %v \n", i, v)
	}
}
