// strings

package main

import (
	"fmt"
	"strings"
)

func main() {

	arr := []string{"Genesys", "The", "Producer"}
	ast := strings.Join(arr, "-")
	fmt.Println(ast)
	spl := strings.Split(ast, "-")
	fmt.Println(spl)
	fmt.Println(strings.Contains(ast, "Genesys"))
	fmt.Println(len(ast))
	fmt.Println(strings.ToLower(ast))
	fmt.Println(strings.ToUpper(ast))
	fmt.Println(strings.ToLower(ast))
	fmt.Println(strings.ToTitle(ast))

}
