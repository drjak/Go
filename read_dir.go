package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check_err(e error){
		if e != nil{
		return
	}
}

func main(){
	dir, err := os.Open("/Users/jka07/Desktop/go/")
	check_err(err)
	defer dir.Close()
	
	fileInfos, err := dir.Readdir(-1)
	check_err(err)
	
	for _, fi := range fileInfos{
		fmt.Println(fi.Name(), fi.Size(),"B")
	}
	fmt.Println("****PRINT PATH******")
	filepath.Walk("/Users/jka07/Desktop/go/", func(path string, info os.FileInfo, err error)error{
			fmt.Println(path)
			return nil
		})
}

