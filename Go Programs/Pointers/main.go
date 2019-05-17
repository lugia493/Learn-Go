package main

import "fmt"

func zero(x int){
	x = 0
}

func zero2(xPtr *int){
	*xPtr = 0
}

func one(xPtr *int){
	*xPtr = 1
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)

	y := 0
	zero2(&y)
	fmt.Println(y)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}
