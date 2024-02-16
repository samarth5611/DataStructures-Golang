// Steps to detect and solve using adapter design patter
// 1) recognize Target interface and Adaptee interface
// 2) Create class adapter containing a field that represents the Adaptee object
// 3) Implement methods for the Target interface. These methods will delegate calls to the appropriate methods of the Adaptee object.
// 4) Create Adaptee, the instance of adapter, then use target interface

package main

import "fmt"

// Target interface represents the expected interface that the client uses
type Target interface {
	Request() string
}

// Adaptee represents the incompatible interface
type Adaptee struct{}

// SpecificRequest represents the method in the incompatible interface
func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// Adapter adapts the Adaptee interface to the Target interface
type Adapter struct {
	adaptee *Adaptee
}

// NewAdapter creates a new Adapter
func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

// Request implements the Request method of the Target interface by calling the Adaptee's SpecificRequest method
func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	// Create an instance of the Adaptee
	adaptee := &Adaptee{}

	// Create an adapter with the Adaptee instance
	adapter := NewAdapter(adaptee)

	// Call the Request method of the Target interface through the adapter
	result := adapter.Request()
	fmt.Println(result) // Output: Specific request
}
