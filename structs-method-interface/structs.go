package structs

import "math"

// Struct is just a named collection of fields where you can store data.
type Rectangle struct {
	Width  float64
	Height float64
}

// Method is a function with a receiver.
// Where you can just call function whenever you like, such as Area(thing) this make
// you can only call on "thing"
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Interface in go is very powerful concept they allow you to make functions that can be
// used with different types and create high-coupled code which still maintaining type-safety
// In this Interface declarion Go will implicit if type you pass in matches what the interface
// is asking for, it will compile in this place Rectangle and Circle had a method called Area
// which returns a float64 so it satisfies the Shape interface
// on other hand, if it is a string, string does not have such a method, so it doesn't satisfy
// the interface
type Shape interface {
	Area() float64
}

// Parimeter takes width and height and return Parimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area takes width and height and return the area of rectangle
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }
