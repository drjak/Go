package main

import ("fmt"; "hash/crc32")

//includes non-crytpto hash functions (adler32,crc32,crc64)
//common use for crc32 is file comparison

func main(){
	h := crc32.NewIEEE()
	
	//crc32 implements the Writer interface, so can write bytes to it like any other Writer
	h.Write([]byte("test"))
	//we call Sum32() to return uint32
	v := h.Sum32()
	fmt.Println("hashed 'test' is: ",v)
}
