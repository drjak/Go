package main

import (
	"os"
	"fmt"
	"log"
	"io/ioutil"
)

func error_dump(ermsg string){
	//convert error type to []byte
	ermsg1 := []byte(ermsg)
	err := ioutil.WriteFile("/Users/jka07/Desktop/go/error.txt", ermsg1, 0666)
		if err != nil{
		fmt.Println("Can't create error dump file")
		return
	}
}

func main(){
	file, err := os.Open("/Users/jka07/Desktop/go/text.txt")
	if err != nil {
		log.Println(err)
		
		//error has an interface
		//type error interface{
		//	Error() string}
		//so can call err.Error() to get string
		error_dump(err.Error())
		//panic cause break and print error
		//panic(err)
		os.Exit(2)
	}
	defer file.Close()
	
	//check file size
	stat, _ := file.Stat()
	if stat.Size() == 0{
		fmt.Println("Size 0B ?")
		error_dump("File size 0B")
		os.Exit(2)
	}
	
	//read the file
	//creates map
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil{
		fmt.Println("Can't read file")
		error_dump(err.Error())
		
	}
	fmt.Println("Before conversion to string", bs)
	str := string(bs)
	fmt.Println(str)
	
	
}
