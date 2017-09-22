package main
import (
        "fmt"
        "os"
        "strconv"
)

func main(){
		if len(os.Args) < 2{
		fmt.Println("Please provide one argument")
		os.Exit(2)
	}
	
	arg := os.Args[1]
	if  arg, err := strconv.Atoi(arg); err == nil {
		if arg > 0 {
			fmt.Println("Positive int:",arg)
			os.Exit(0)
		}else if arg < 0{
				fmt.Println("Negative int:",arg)
				os.Exit(1)
		}else {
			fmt.Print("Zero is neither positive or negative")
		}
		
	}else {
		fmt.Print("It should be integer!\n")
		os.Exit(2)
	}
	
}
