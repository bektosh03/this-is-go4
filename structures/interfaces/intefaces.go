package main

import (
	"fmt"
)

type Figure interface {
	Surface() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Surface() float64 {
	return 3.14 * c.radius * c.radius
}

type Square struct {
	side float64
}

func (s Square) Surface() float64 {
	return s.side * s.side
}

func main() {
	var c = Circle{radius: 5}
	var s = Square{side: 10}
	surfaceC := CalculateSurface(c)
	surfaceS := CalculateSurface(s)
	fmt.Println(surfaceC, surfaceS)

	a := 6
	b := "asdfaf"
	abc(a)
	abc(b)
	abc(c)
}

func CalculateSurface(f Figure) float64 {
	return f.Surface()
}

func abc(a interface{}) bool {
	var ok bool
	switch a.(type) {
	case int:
		fmt.Println("This is an integer:", a.(int))
		ok = true
	case string:
		fmt.Println("this is a string:", a.(string))
		ok = true
	default:
		fmt.Println("Unrecognized type:")
	}
	return  ok
}
