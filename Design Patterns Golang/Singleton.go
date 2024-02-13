// Steps to create Singleton Design Pattern
// 1) Create a Product Class
// 2) Create a global object of that class, and "once sync.One"
// 3) write a method getinstance for declaring prev golang object
// 4) use once.Do, and rest becomes similiar

package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

func getInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{"Initialized"}
	})
	return instance
}

func (s *Singleton) getData() string {
	return s.data
}

func (s *Singleton) setData(str string) {
	s.data = str
}

func main() {
	instance1 := getInstance()
	fmt.Println(instance1)
	instance1.setData("golang")

	instance2 := getInstance()
	fmt.Println(instance2)
}
