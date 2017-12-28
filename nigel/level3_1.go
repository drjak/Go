package main
import(
	"io"
	"net/http"
	"log"
	"math/rand"
	"encoding/json"
	"fmt"
)
type Json_builder struct{
	 Message string
}

func displayText(w http.ResponseWriter, r *http.Request){
	messages := []string{"first message", "second message", "third message", "fourth message"}
	n := rand.Intn(len(messages))
	
	object := &Json_builder{ Message:messages[n]}	
	msg, _ := json.Marshal(object)
	fmt.Println(string(msg))
	io.WriteString(w, string(msg))
}

func main(){
	http.HandleFunc("/", displayText)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", nil)
	}
}