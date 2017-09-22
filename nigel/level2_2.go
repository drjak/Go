package main

import (
        "fmt"
        "net/http"
        "os"
)
func check_args(){
		if len(os.Args) < 2{
			fmt.Println("Please provide URL")
			os.Exit(2)
	}
}

func check_res(url string){
	
	        req, err := http.Get(url)
		    if err !=nil {
	        	fmt.Print("Some error")
	        	os.Exit(2)
	        }

		code := req.StatusCode
        fmt.Println("Reponse", code)
}

func main() {

url := []string{"http://wp.pl", "http://google.com", "http://bbc.com"}
for i := 0; i<len(url); i++ {
	fmt.Println("site ",url[i])
	check_res(url[i])
}
        

        
}