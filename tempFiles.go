/* tempFiles */
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

	tmpFile, err := ioutil.TempFile(".", "Data-*")
	check(err)
	fmt.Println(tmpFile.Name())
	fmt.Println(os.Stat(tmpFile.Name()))
}
