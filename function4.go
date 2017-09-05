package main

import "fmt"

//function that return other function
func makeEvenGeneraor()func()uint{
	i := uint(0)
	return func() (ret uint) {
		ret = 1
		i += 2
		return
	}
}
func main(){
	nextEven := makeEvenGeneraor()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}