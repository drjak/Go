package main

import (
	"fmt"
	"bufio"
	"io/ioutil"
	"os"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func main(){
	//create with ioutil
	//ioutil requires []byte type
	d1 := []byte("heloo\ngo\n")
	err := ioutil.WriteFile("/Users/jka07/Desktop/go/writing_IOUTUL.txt", d1, 0644)
	check(err)
	
	//create with OS
	f, err := os.Create("/Users/jka07/Desktop/go/writing_OS.txt")
	check(err)
	defer f.Close()
	
	d2 := []byte{115, 111, 109, 101, 10}
	n1, err := f.Write(d2)
	check(err)
	fmt.Printf("Wrote %d bytes\n", n1)
	
	n2, err := f.WriteString("writes\n")
	fmt.Printf("Wrote %d bytes\n", n2)
	
	f.Sync()
	
	//using buffio
	w := bufio.NewWriter(f)
	n3, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n",n3)
	w.Flush()
}

