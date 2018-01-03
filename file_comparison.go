package main

import ("fmt";"hash/crc32";"io/ioutil")

func getHash(filename string)(uint32, error){
	bs, err := ioutil.ReadFile(filename)
	if err != nil{
		return 0, err
	}
	
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

func main(){
	h1, err := getHash("/Users/jka07/Desktop/go/writing_IOUTUL.txt")
	if err != nil {
		return
	}
	h2, err := getHash("/Users/jka07/Desktop/go/writing_OS.txt")
	if err != nil {
		return
	}
	fmt.Println("file1:", h1," file2: ",h2, h1==h1)

}

