package main

import (
"fmt"
	"time"
)

func main(){
	c1 := make(chan string)
	c2 := make(chan string)
	
	go func(){
		for{
			c1 <- "from 1 - one sec"
			time.Sleep(time.Second *1)
		}
	}()
	
	go func(){
		for{
			c2 <- "from 2 - two sec"
			time.Sleep(time.Second *2)
		}
	}()
	//pick the channel that is ready and receives from it(or send to it)
	//if more than one is ready it randomly picks which one to receive from
	go func(){
		for{
			select{
				case msg1 := <- c1:
				fmt.Println(msg1)
				case msg2 := <- c2:
				fmt.Println(msg2)
				//if none channel is ready
				case <- time.After(time.Second):
				fmt.Println("timeout")
				//default:
				//fmt.Println("nothing ready")
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}