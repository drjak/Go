package main

import ("fmt")

type Person struct {
	Name string
}

func (p *Person)Talk(){
	fmt.Println("Hi, my name is: ", p.Name)
}
type Android struct {
	Person Person
	Model string
}
func main(){
	//we use the type (Person) and don' give it a name
	//when defined this way the Person struct can be access using the type name
	a := new(Android)
	a.Person.Talk()

	//but we also call any Person method directly on Android
	//b := new(Android)
	//b.Talk()
}