package main
import (
        "fmt"
        "io/ioutil"

)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func vowel_count(myStr string) {
        t := 0
        for _, value := range myStr {
                switch value {
                case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
                        t++
                }
        }
                fmt.Printf("Vowels: %d \n", t)
}
func main(){
    content,err := ioutil.ReadFile("/tmp/Jarek")
    check(err)
    fmt.Print(string(content))

    myStr := string(content)
    fmt.Println("Length: ", len(myStr))
    
    vowel_count(myStr)
}

