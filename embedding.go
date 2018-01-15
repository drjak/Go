package main

import ( "fmt"
)

type Cat struct{
	Name string
}
func (c Cat)Legs()int{
	return 4
}

func(c Cat)PrintLegs(){
	fmt.Printf("Have %d legs", c.Legs())
}

type OctoCat struct{
	Cat
}

func (o OctoCat)Legs()int{
	return 8
}

func main(){
	var octo OctoCat
	fmt.Println(octo.Legs())
	octo.PrintLegs()
}

