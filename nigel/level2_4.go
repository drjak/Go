package main
import "fmt"

func main() {
	letters_map := map[string]int{"a":1,"b":2,"c":3,"d":4,"e":5,"f":6,"g":7,"h":8,"i":9,"j":10,"k":11,"l":12,"m":13,"n":14,"o":15,"p":16,"q":17,"r":18,"s":19,"t":20,"u":21,"v":22,"w":23,"x":24,"y":25,"z":26 }
	sentence := "the quick brown fox"
	var value int
	var sum int
	for _, letter := range sentence{
		value=letters_map[string(letter)]
		fmt.Println(string(letter), "-", value)
		sum += value
	}	
	fmt.Println(sum)
}