package main

import "fmt"

/*Good morning! Here's your coding interview problem for today.

This problem was asked by Glassdoor.

An imminent hurricane threatens the coastal town of Codeville.
If at most two people can fit in a rescue boat, and the maximum weight
limit for a given boat is k, determine how many boats will be needed to save everyone.

For example, given a population with weights [100, 200, 150, 80] and a
boat limit of 200, the smallest number of boats required will be three.*/

func main() {

	// Size of list
	var size int = 0
	fmt.Print("Enter size of list : ")
	fmt.Scan(&size)
	slc := make([]int, size)

	// Enter value into a slice
	for i := 0; i < size; i++ {
		num := 0
		fmt.Print("Enter a number : ")
		fmt.Scan(&num)
		slc[i] = num
	}
	// Weight
	var weight int = 0
	fmt.Print("Enter weight : ")
	fmt.Scan(&weight)
	fmt.Println(getBoats(slc, weight))
}

func getBoats(num []int, weight int) int {

	sum := 0
	notEven := false
	for i := range num {
		sum += num[i]
		//fmt.Println(sum)
	}

	boats := sum / weight
	if sum%weight > 0 {
		notEven = true
	} else {
		notEven = false
	}

	if notEven == true {
		return boats + 1
	}

	return boats

}
