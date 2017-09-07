package main
import ("fmt"; "math")

type Circle struct {
	x, y, r float64
}
func circleArea(c Circle)float64{
	return math.Pi * c.r * c.r
}
func circleArea1(c *Circle)float64{
	return math.Pi * c.r * c.r
}
//method
func (c *Circle)area()float64{
	return math.Pi * c.r * c.r
}
func main(){
	//we can create instance of Circle in many ways

	//local Circle variable that is by default set o zero
	//for strut zero means each of the fields is set to zero
	var a Circle
	//we can acess fields by using . (dot)
	fmt.Println(a.x, a.y, a.r)

	//allocate memory for all the fields, sets each to zero and return pointer (*Circle)
	b := new(Circle)
	fmt.Println(b.x, b.y, b.r)
	b.x = 10
	b.y = 5
	fmt.Println(b.x,b.y, b.r)

	//when want to give each field value
	c := Circle{x:0, y:0, r:5}
	//or
	d := Circle{0, 0, 6}

	fmt.Println(circleArea(c))
	//if we attempt to modify one of the fields inside of the circleArea function
	//it would not modify the original variable, because of this we would write function like circleArea1
	fmt.Println(circleArea1(&d))

	e := Circle{1,2,3}
	fmt.Println(e.area())

}