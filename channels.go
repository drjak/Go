package main
import (
	"fmt"
	"time"
)
func pinger(c chan string){
	for i:=0;;i++{
		c <- "ping"
	}
}
func counter(d chan int){
	for i:=0;;i++{
		d <- i
	}
}
func printer(c chan string, d chan int){
	for{
		msg := <- c
		fmt.Println(msg)
		fmt.Println(<- d)
		time.Sleep(time.Second*1)
	}
}
func main(){
	var c chan string = make(chan string)
	var d chan int = make(chan int)
	go pinger(c)
	go counter(d)
	go printer(c,d)

	var input string
	fmt.Scanln(&input)
}