package main

import "fmt"

/*Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given a array of numbers representing the stock prices of a company in
chronological order, write a function that calculates the maximum profit you
could have made from buying and selling that stock once. You must buy before you can sell it.
For example, given [9, 11, 8, 5, 7, 10], you should return 5,
since you could buy the stock at 5 dollars and sell it at 10 dollars.

*/

func main() {

	arr := []int{9, 11, 8, 5, 7, 10}

	sums := 0
	var num = 0
	for v := range arr {
		for b := range arr {
			if v > b {
				num = arr[v] - arr[b]
			}
			if v < b {
				num = (-arr[v]) + arr[b]
			}
			if v == b {
				num = 0
			}

			fmt.Println(arr[v], "<- on this number ->", arr[b], num)
			if num > sums {
				sums = num
			}

		}
	}

	fmt.Println("Answer : ", sums)
}
