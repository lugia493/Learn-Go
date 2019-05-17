package main

import "fmt"
import "math"

type Circle struct {
	x, y, r float64
}
type Rectangle struct {
	x1, y1, x2, y2 float64
}
type Distance struct {
	x1, y1, x2, y2 float64
}
func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1) 
	return l * w 
}

type Person struct {
	Name string
}
type Android struct {
	Person
	Model string
}
func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}
func main(){
	c := Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(c.area())

	d := new(Distance)
	d.x1, d.y1, d.x2, d.y2 = 0, 0, 10, 10

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.x1, r.x2, r.y1, r.y2)
	fmt.Println(r.area())

	a := new(Android)
	a.Name = "Daniel"
	a.Talk() 
}
