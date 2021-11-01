// listeningReq
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting Web Server on Port 8000")
	http.ListenAndServe(":8000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Handler")
	io.WriteString(w, "Response handler \n")
	io.WriteString(w, "Path : "+r.URL.Path)

}
