package main

import (
"fmt"
)

func zeroval(ival int){
	ival = 0
}

func zeroptr(iptr *int){
	*iptr = 0
}

func main(){
	 i := 1
	 fmt.Println("initilal value i is: ",i)
	 
	 zeroval(i)
 	 fmt.Println("passing i value to non pointer: ",i)

	//The &i syntax gives the memory address of i, i.e. a pointer to i.
	zeroptr(&i)
	fmt.Println("passing to pointer zeroptr: ",i)
	//Pointers can be printed too.
	fmt.Println("pointer can be printed too &i: ",&i)

}