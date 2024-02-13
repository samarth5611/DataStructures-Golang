package main

import "fmt"

type Person struct {
	// Personal details
	name, address, pin string
	// Job details
	company, position string
	salary            int
}

type Builder interface {
	setName(string) Builder
	setAddress(string) Builder
	setPin(string) Builder
	setCompany(string) Builder
	setPosition(string) Builder
	setSalary(int) Builder
	Build() *Person
}

type ConcreateBuilder struct {
	// Personal details
	name, address, pin string
	// Job details
	company, position string
	salary            int
}

func (cb *ConcreateBuilder) setName(s string) Builder {
	cb.name = s
	return cb
}

func (cb *ConcreateBuilder) setAddress(s string) Builder {
	cb.address = s
	return cb
}

func (cb *ConcreateBuilder) setPin(s string) Builder {
	cb.pin = s
	return cb
}

func (cb *ConcreateBuilder) setCompany(s string) Builder {
	cb.company = s
	return cb
}
func (cb *ConcreateBuilder) setPosition(s string) Builder {
	cb.position = s
	return cb
}
func (cb *ConcreateBuilder) setSalary(s int) Builder {
	cb.salary = s
	return cb
}

func (cb *ConcreateBuilder) Build() *Person {
	return &Person{
		name:     cb.name,
		address:  cb.address,
		pin:      cb.pin,
		company:  cb.company,
		position: cb.position,
		salary:   cb.salary,
	}
}

func main() {
	builder := &ConcreateBuilder{}
	person := builder.setName("Samartha").
		setAddress("India").
		setPin("123456").
		setCompany("Samsung").
		setPosition("Software Engineer").
		setSalary(10).
		Build()

	fmt.Println(person)
}
