package main

import (
        "fmt"
)
func main(){
	var fizz, buzz, fizzbuzz int
	for i :=1; i<=100; i++ {
		if i%3==0 && i%5!=0{
			fmt.Println("Fizz", i)
			fizz++
		}
		if i%5==0 && i%3!=0{
			fmt.Println("Buzz", i)
			buzz++
		} 
		if i%5==00 && i%3==0 {
			fmt.Println("FizzBuzz", i)
			fizzbuzz++
		}else if i%5!=0 && i%3!=0 {
			fmt.Println("Neither", i)
		}
}
			fmt.Printf("Fizz %s, Buzz %s, FizzBuff %s", fizz, buzz, fizzbuzz)

}