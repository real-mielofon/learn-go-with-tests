package shapes

import (
	"math"
)

// Shape is innterface of shape with Area() and Perimeter() methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle is type passed rectangle with width and height
type Rectangle struct {
	width  float64
	height float64
}

// Perimeter return perimeter of square with width and height
func (rectangle Rectangle) Perimeter() float64 {
	return (rectangle.width + rectangle.height) * 2
}

// Area return area of square with width and height
func (rectangle Rectangle) Area() float64 {
	return (rectangle.width * rectangle.height)
}

// Circle is type passed circle with radius
type Circle struct {
	radius float64
}

// Perimeter return perimeter of circle with redius
func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

// Area return area of circle with redius
func (circle Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// Triangle is type shape triangle with base and height
type Triangle struct {
	base   float64
	height float64
}

// Perimeter return perimeter of triangle with base and height
func (triangle Triangle) Perimeter() float64 {
	return triangle.base + math.Sqrt(triangle.base*triangle.base/4+triangle.height*triangle.height)
}

// Area return area of circle with triangle with base and height
func (triangle Triangle) Area() float64 {
	return triangle.base * triangle.height / 2
}
