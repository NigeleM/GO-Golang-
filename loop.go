package main
import "fmt"

func main(){

 var number int

    fmt.Print("Enter a number: ")
    fmt.Scanf("%d", &number)
    fmt.Println("number is", number)

    for count := 0; count <= number; count++ {
      fmt.Println(count)
      }
}
