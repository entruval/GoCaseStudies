package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Object interface {
	Volume() float64
}

type CircleObject interface {
	Area() float64
	Perimeter() float64
}

// embedded interfaces
type Material interface {
	Shape
	Object
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rect) Volume() float64 {
	return r.width
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var s Shape = Rect{10, 3}
	r := Rect{5.0, 4.0}
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("area: ", s.Area())
	fmt.Println("s == r ", s == r) // s & r hold the same dynamic type rect struct & dynamic value {5, 4}
	// fmt.Println(s.Volume()) // error -- because s is Shape not Object interface
	fmt.Println(r.Volume(), "\n")
	_, ok := s.(Shape) // type assertion
	fmt.Println(ok, "\n")

	s = &Circle{10} // s here is struct type
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)

	var c CircleObject = &Circle{10} // c here not an interface but circle struct after use pointer
	// circle Perimeter will automatically use receiver, because of the same circle struct
	fmt.Println("area: ", c.Area())

	var m Material = Rect{3, 2} // because Shape or Object is one of the embedded, Rect must automatically implement Material
	fmt.Printf("type of m is %T\n", m)
	fmt.Printf("value of m is %v\n", m)
}
