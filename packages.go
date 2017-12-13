package main

import (
"fmt"
	"mypacks"
)

func main(){
	var xs []float64
	fmt.Scanln(&xs)
	avg := math.Average(xs)
	fmt.Println(avg)
}

