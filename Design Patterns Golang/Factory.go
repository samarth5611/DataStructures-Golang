// Creation of multiple objects that share a common interface or superclass.
// Define Interface first
// Define Factory Class, with function that can return different objects
package main

import "fmt"

// Shape interface represents a basic shape.
type Shape interface {
	Draw() string
}

// Circle represents a circle shape.
type Circle struct{}

// Rectangle represents a rectangle shape.
type Rectangle struct{}

// Draw draws a circle.
func (c *Circle) Draw() string {
	return "Drawing Circle"
}

// Draw draws a rectangle.
func (r *Rectangle) Draw() string {
	return "Drawing Rectangle"
}

// ShapeFactory creates shapes.
type ShapeFactory struct{}

// GetShape creates a shape based on the given type.
func (sf *ShapeFactory) GetShape(shapeType string) Shape {
	if shapeType == "Circle" {
		return &Circle{}
	} else if shapeType == "Rectangle" {
		return &Rectangle{}
	} else {
		return nil
	}
}

func main() {
	shapeFactory := &ShapeFactory{}

	// Get an instance of Circle and call its Draw method.
	circle := shapeFactory.GetShape("Circle")
	fmt.Println(circle.Draw())

	// Get an instance of Rectangle and call its Draw method.
	rectangle := shapeFactory.GetShape("Rectangle")
	fmt.Println(rectangle.Draw())
}
