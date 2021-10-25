// map Key and values

package main

import "fmt"

func main() {

	passwordMap := make(map[string]string)
	passwordMap["Nigele"] = "password123"
	passwordMap["Khris"] = "123password"
	passwordMap["Data"] = "123pass123"

	for users, pass := range passwordMap {
		fmt.Println("Username : ", users, " Password : ", pass)
	}

}
