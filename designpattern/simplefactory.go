package designpattern

import "fmt"

// Product interface
type Product interface {
	DoSomething()
}

// ConcreteProduct of Product
type ConcreteProduct struct {
	Product
}

// DoSomething produce product1
func (p ConcreteProduct) DoSomething() {
	fmt.Println("produce product")
}

// ConcreteProduct1 of Product
type ConcreteProduct1 struct{}

// DoSomething produce product1
func (p1 ConcreteProduct1) DoSomething() {
	fmt.Println("produce product1")
}

// ConcreteProduct2 of Product
type ConcreteProduct2 struct{}

// DoSomething produce product2
func (p2 ConcreteProduct2) DoSomething() {
	fmt.Println("produce product2")
}

// SimpleFactory struct for Simple Factory pattern
type SimpleFactory struct{}

// Produce products
func (f SimpleFactory) Produce(t int) Product {
	switch t {
	case 1:
		return ConcreteProduct1{}
	case 2:
		return ConcreteProduct2{}
	default:
		return ConcreteProduct{}
	}
}
