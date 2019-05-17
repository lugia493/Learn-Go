package main

import "fmt"

var z int = 8

func main(){
	var (   //Can use const too
		 a = 10
		 b = 15
		 c = 40
	)

	fmt.Println("Hello World!")
	var x string = "Hello World"
	fmt.Println(x)
	y := "Hello World"
	fmt.Println(y)
	fmt.Println("I am printing out", y)
	fmt.Println(z)
	const foo int = 10
	fmt.Println(foo)
	fmt.Println(a + b + c)
}
