package main

import (
	"fmt"
)
func average(xs []float64)float64{
	total := 0.0
	for _, a := range xs{
		total += a
	}
	return total / float64(len(xs))
}
func main(){
	xs := []float64{54,56,56,65,78}
	fmt.Println(average(xs))

	some_other := []float64{45,657,34,65,7}
	fmt.Println(average(some_other))
}
