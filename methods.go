package main

import (
	"fmt";
	"math"
)
type Rectangle struct {
	x1,y1,x2,y2 float64
}
func distance(x1,y1,x2,y2 float64)float64{
	a := x2-x1
	b := y2-y1
	return math.Sqrt(a*a+b*b)
}

func (r *Rectangle)area()float64{
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l*w
}
func main(){
	g := Rectangle{0,0,10, 10}
	fmt.Println("distance",distance(1,2,3,4))
	fmt.Println(g.area())
}