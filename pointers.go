package main

import "fmt"

func zero(xPtr *int){
	*xPtr = 0
}
func main(){
	x:=5
	zero(&x)
	fmt.Println(x)  //x is 0
}
//pointers reference a location in memory where value is stored rater
//than the value itlsef
//by using a pointer (*int) the zero function is able to modify original variable

