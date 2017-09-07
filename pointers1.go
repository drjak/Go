package main
import "fmt"

func one(xPtr *int){
	*xPtr = 1
}

func main(){
	//new takes a type as an argument, allocates memory to fit a value and returns pointer
	xPtr2 := new(int)
	xPtr3 := new(int)
	one(xPtr2)
	one(xPtr3)
	fmt.Println(*xPtr2)
	fmt.Println(*xPtr3)
}