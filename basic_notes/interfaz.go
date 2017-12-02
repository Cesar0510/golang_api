// _Interfaces_ are named collections of method
// signatures.

package main

import "fmt"
import "math"

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
}

type triangule struct {
	width float64
}

func (t triangule) area2() float64 {
	return t.width * t.width
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())

}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	t := triangule{width: 7}

	measure(r)
	measure(c)
	measure(t)
}
