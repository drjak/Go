package main

import "fmt"

//When we write *xPtr = 0 we are saying “store the int 0 in the memory location xPtr refers to
//by using a pointer (*int) the zero function is able to modify original variable
func zero(xPtr *int){
	*xPtr = 0
}

func square(x *float64){
	*x = *x * *x
}

func main(){
	x:=5
	
	fmt.Println(x)  //x is 5
	
//Finally we use the & operator to find the address of a variable. &x returns a *int (pointer to an int) because x is an int. 
//This is what allows us to modify the original variable. &x in main and xPtr in zero refer to the same memory location
	zero(&x)
	fmt.Println(x)  //x is 0

	y := 1.5
	square(&y)
	fmt.Println(y)


}
/*
Pointers reference a location in memory where value is stored rather than the value itlsef.
*/
