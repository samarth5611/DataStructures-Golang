// steps to build a product class
// 1) Create Builder interface for Building functions return builder, 
// this should be step by step process, with build function too return product pointer
// 2) Create Concreate-Class with exact same variable of Product class
// 3) Define BuildingSteps under concreate class with returning builder interface
// 4) Last define Build() method under concreate class, and assign 
//    all the values in concreate class object to return product class


package main
import "fmt"

type Product struct{
	sa string 
	sb string
}
type Builder interface{
	buildstep1(string)Builder
	buildstep2(string)Builder
	Build()*Product
}
// buildstep1
// buildstep2
// concretebuilder class

type concretebuilder struct{
	sa string 
	sb string
}

func (cb *concretebuilder) buildstep1(keyString string) Builder{
	cb.sa = keyString
	return cb
}
func (cb *concretebuilder) buildstep2(keyString string) Builder{
	cb.sb = keyString
	return cb
}
func (cb *concretebuilder)Build()*Product{
	return &Product{
		sa: cb.sa,
		sb: cb.sb,
	}
}

func main(){
	builder := &concretebuilder{}
	product := builder.
			buildstep1("Samartha").
			buildstep2("Jadhao").
			Build()
	fmt.Println(product)
}