package main

import "fmt"

func add(args ...int)int{
	total := 0
	for _, v := range args  {
		total+=v
	}
	return total
}
func average(xs []float64)float64{
        total := 0.0
        for _, a := range xs{
                total += a
        }
        return total / float64(len(xs))
}

//function that return other function
func makeEvenGenerator()func()uint{
	i := uint(0)
	return func() (ret uint){
		ret = 1
		i += 2
		return
		}
	}

//deffering functions
func first(){
	fmt.Println("1st")
	}
func second(){
	fmt.Println("2nd")
	}

func main(){
	fmt.Println("simple adding",add(1,2,4))
	//you can also pass slice
	slice := []int{1,2,3}
	fmt.Println("slice: ", add(slice...))

        //using function average
        xs := []float64{54,56,56,65,78}
        fmt.Println(average(xs))

        some_other := []float64{45,657,34,65,7}
        fmt.Println(average(some_other))
	
	//function inside the function
	sum := func(x,y int)int{
		return x+y
		}
		//add is a local variable that has a type func(int, int)int
		//when created local functions like this it has access to other local variables
	fmt.Println("function inside function, sum",sum(1,6))

 	x := 0
	increment := func()int{
		x++
		return x
		}
	fmt.Println("increment", increment())
	fmt.Println("increment", increment())
        
	//return other function test
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	
	//deffering
	fmt.Println("testing deffering")
	defer second()
	first()

}
