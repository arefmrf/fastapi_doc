package main

import "fmt"

// Shape interface
type Shape interface {
	Draw()
}

// Circle struct
type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing a Circle")
}

// Square struct
type Square struct{}

func (s *Square) Draw() {
	fmt.Println("Drawing a Square")
}

// ShapeFactory function
func ShapeFactory(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	default:
		return nil
	}
}

func RunSimpleFactory() {
	// Creating objects using the factory
	shape1 := ShapeFactory("circle")
	shape1.Draw()

	shape2 := ShapeFactory("square")
	shape2.Draw()
}
