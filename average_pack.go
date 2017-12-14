package main

import (
	"fmt"
	"bufio"
	"mypacks"
	"os"
	"strconv"
)

func main(){
	fmt.Println("Enter numbers to calculate average")
	a := ReadInput()
	avg := math.Average(a)
	fmt.Println(len(a), "items passed in array",a, "sum of items", Sum(a))
	fmt.Println("Averages: ", avg)
}

func ReadInput()[]float64{
	scan := bufio.NewScanner(os.Stdin)
	var slice []float64
	for scan.Scan() {
		//strconv converts to&from strings from&to other types
		//64 at the end specifies float
		val, err := strconv.ParseFloat(scan.Text(), 64)
		//two enters send empty string and this finsihes loop
		if err != nil {
			break
		}	
		//appending next scanned item to slice
		slice = append(slice, []float64{val}...)
	}
	return slice
}
func Sum(xs []float64)float64{
	sum := float64(0)
	for _, x := range xs{
		sum += x
	}
	return sum
}
//mypack/math.go is in /usr/local/go/src/mypacks
/*package math

func Average(xs []float64)float64{
	total := float64(0)
	for _, x := range xs{
		total += x
	}
	return total / float64(len(xs))
	}
func Sum(xs []float64)float64{
	sum := float64(0)
	for _, x := range xs{
		sum += x
	}
	return sum
}
*/




