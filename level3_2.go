package main
import(
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
)

type Json_builder struct{
	Message string
}
func json_value(msg string){
	var v Json_builder
	json.Unmarshal([]byte(msg), &v)
	fmt.Println("Value only:",v.Message)
}

func main(){
	url := "http://localhost:9000/"
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println("Can't connect to server at ", url)
	}	
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Whole body: ",string(body))
	
	json_value(string(body))
}