// steps to implement prototype design
// 1) create a prototype interface with Clone() function return interface only
// 2) define clone function, returning interface, convert interface to class

package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type Product struct {
	name string
}

func NewProduct(name string) *Product {
	return &Product{name}
}

func (c *Product) Clone() Prototype {
	return &Product{name: c.name}
}

func main() {
	// Create an original object
	original := NewProduct("original")

	// Clone the original object
	clone := original.Clone().(*Product)

	// Output the details of the original and cloned objects
	fmt.Println("Original object:", original.name)
	fmt.Println("Cloned object:", clone.name)
}
