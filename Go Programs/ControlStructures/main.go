package main

import "fmt"

func main(){
	fmt.Println(`1
	2
	3
	4
	5
	6
	7
	8
	9
	10`)

	i := 1
	for i <= 10{
		fmt.Println(i)
		i += 1
	}
	for i := 1; i <= 10; i++ {
		fmt.Print(i)
		fmt.Print(" ")
	}

	var x [5]int
	for i := 0; i < len(x); i++{
		x[i] = i
		fmt.Print(x[i])
	}
	z := [5]int{98, 34, 13, 65, 75}
	fmt.Println(len(z))

	var total int = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total)
}