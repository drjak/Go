package main

import (
	"fmt"
	"container/list"
)
func main(){
	/*linked list
	each node of list contain value (1,22,31) and a pointer to next node
	*/
	var x list.List
	x.PushBack(1)
	x.PushBack(22)
	x.PushBack(31)
	/* zero value for a List is an empty list 
	Values are appended to the list using PushBack. We loop back over each item 
	*/
	for e := x.Front(); e != nil; e=e.Next(){
		fmt.Println(e.Value.(int))
	}
}

