package main

import ( "fmt"; "sort")

type Person struct{
	Name string
	Age int
}

type ByName []Person
type ByAge []Person
/* Sort function (sorting arbitraty data) has several predefined functions
Sort function takes sort.Interface and sorts it. The sort.Interface requires 3 methods: Len, Less, Swap
To define our own sort we create a new type ByName and make eqivalent to a slice of what we want to sort
We then define 3 methods
*/
func (this ByName) Len() int{
	return len(this)
}
func (this ByName) Less(i, j int) bool{
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int){
	this[i], this[j] = this[j], this[i]
}

func (this ByAge) Len() int{
	return len(this)
}
func (this ByAge) Less(i, j int) bool{
	return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int){
	this[i], this[j] = this[j], this[i]
}


func main(){
	kids := []Person{
		{"Jill",9},
		{"Jack",10},
		{"Hans",12},
	}
	
	fmt.Println("By Name\n")
	sort.Sort(ByName(kids))
	fmt.Println(kids)
	
	fmt.Println("\n by age")
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}