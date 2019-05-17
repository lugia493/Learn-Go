package main

import "fmt"

func main(){
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
	fmt.Println(x["key"])
	delete(x, "key")
	x["change"] = 7
	fmt.Println(x)
	fmt.Println(x["change"])
}