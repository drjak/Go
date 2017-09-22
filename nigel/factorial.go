package main
import (
"fmt"
)

var input uint64

func read_from_user() {
        fmt.Print("Enter number to be calculated: ")
        fmt.Scanln(&input)
}

func main(){
        read_from_user()
        
        var i uint64
        var fact uint64 = 1
        for i = 1; i<=input; i++{
                fact *= i
        }
      fmt.Printf("Factorial from %d id %d \n", input,fact)

}
