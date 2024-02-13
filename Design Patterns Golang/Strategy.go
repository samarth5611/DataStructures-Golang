package main

import "fmt"

// min of 2 numbers and max of 2 numbers using startegy

// 1) startegy interface
// 2) implement startegy interface, subclass defining startegies
// 3) context class, function like "setstrategy" and "Action" for strategy

type Strategy interface {
	returnNum(int, int) int
}

type MaxStrategy struct{}

func (_ *MaxStrategy) returnNum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type MinStrategy struct{}

func (_ *MinStrategy) returnNum(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

type Context struct {
	returnStartegy Strategy
}

func (c *Context) setstrategy(strategy Strategy) {
	c.returnStartegy = strategy
}

func (c *Context) Action(a int, b int) int {
	return c.returnStartegy.returnNum(a, b)
}

func main() {
	c := &Context{}
	c.setstrategy(&MaxStrategy{})
	a, b := 5, 10
	fmt.Println(c.Action(a, b))
	c.setstrategy(&MinStrategy{})
	fmt.Println(c.Action(a, b))
}
