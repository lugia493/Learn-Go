package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func third(){
	fmt.Println("3rd")
}

func main() {
	defer second()
	first() 

	//defer used when resources need to be freed in some way
	//Ex. When we open a file we need to make sure to close it later 
	/*
	f, _ := os.Open(filename)
	defer f.Close()
	*/
	
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic ("PANIC")
}