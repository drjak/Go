package main

import "fmt"

//functions inside the functions
func main(){
	add := func(x,y int)int{
		return x+y
	}
	fmt.Println(add(1,5))
}
//add is a local variable that has a type func(int, int)int
//when created local functions like this it has access to other local variables
