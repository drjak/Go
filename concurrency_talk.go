//compare with concurrency_talk1.go which makes more sense

package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	//declare and initialize  channel
	c := make(chan string)
	go boring("boring!", c)
	
	for i := 0; i<5; i++{
		fmt.Printf("you say: %q\n", <-c ) //reveive from channel
	}
	fmt.Println("leaving")
}
func boring(msg string, c chan string){
	for i:=0;;i++{
		c <- fmt.Sprintf("%s %d", msg, i) //expression sent to channel can be any value
		time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
	}
}

