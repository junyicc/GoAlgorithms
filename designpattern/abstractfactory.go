package designpattern

import "fmt"

// AbstractProductA interface
type AbstractProductA interface {
	DoSomethingA()
}

// AbstractProductB interface
type AbstractProductB interface {
	DoSomethingB()
}

// ProductA1 for Abstract Factory pattern
type ProductA1 struct{}

// DoSomethingA for ProductA1
func (p ProductA1) DoSomethingA() {
	fmt.Println("ProductA func")
}

// ProductA2 for Abstract Factory pattern
type ProductA2 struct{}

// DoSomethingA for ProductA2
func (p ProductA2) DoSomethingA() {
	fmt.Println("ProductA func")
}

// ProductB1 for Abstract Factory pattern
type ProductB1 struct{}

// DoSomethingB for ProductB1
func (p ProductB1) DoSomethingB() {
	fmt.Println("ProductA func")
}

// ProductB2 for Abstract Factory pattern
type ProductB2 struct {
	AbstractProductB
}

// DoSomethingB for ProductB2
func (p ProductB2) DoSomethingB() {
	fmt.Println("ProductA func")
}

// AbstractFactory for Abstract Factory pattern
type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

// ConcreteFactoryAB1 for Abstract Factory pattern
type ConcreteFactoryAB1 struct{}

// CreateProductA for ConcreteFactoryAB1
func (cf ConcreteFactoryAB1) CreateProductA() AbstractProductA {
	return ProductA1{}
}

// CreateProductB for ConcreteFactoryAB1
func (cf ConcreteFactoryAB1) CreateProductB() AbstractProductB {
	return ProductB1{}
}

// ConcreteFactoryAB2 for Abstract Factory pattern
type ConcreteFactoryAB2 struct{}

// CreateProductA for ConcreteFactoryAB2
func (cf ConcreteFactoryAB2) CreateProductA() AbstractProductA {
	return ProductA2{}
}

// CreateProductB for ConcreteFactoryAB2
func (cf ConcreteFactoryAB2) CreateProductB() AbstractProductB {
	return ProductB2{}
}
