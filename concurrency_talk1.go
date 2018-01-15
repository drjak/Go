package main

import (
	"fmt"
	"math/rand"
	"time"
)
func boring(msg string) (<- chan string){  //returns receive-only channel of strings
	c := make(chan string)
	go func(){
		for i:=1;;i++{
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)

		}
	}()
	return c //return the channel to the caller
}

func main(){
	c := boring("boring!")
	for i := 0; i<5; i++{
		fmt.Printf("you say: %q\n", <-c ) //reveive from channel
	}
	fmt.Println("leaving")
}
