package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs{
		total += v
	}
	return total/float64(len(xs))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func EvenNumberGen() func() uint {
	i := uint(100)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fact(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * fact(x - 1)
}

func main(){
	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))
	fmt.Println(add(1,2,3,4,5))

	ys := []int{1,2,3,4,5,6}
	fmt.Println(add(ys...))

	subtract := func (x, y int) int{
		return x - y
	}
	fmt.Println(subtract(10,7))

	a := 0
	increment := func() int {
		a++
		return a
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := EvenNumberGen()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println(fact(10))
}

