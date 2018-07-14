package designpattern

// Factory for Factory Method pattern
type Factory interface {
	Produce() Product
}

// ConcreteFactory strcut for Factory Method pattern
type ConcreteFactory struct{}

// Produce ConcreteProduct
func (f ConcreteFactory) Produce() Product {
	product := ConcreteProduct{}
	return product
}

// ConcreteFactory1 strcut for Factory Method pattern
type ConcreteFactory1 struct{}

// Produce ConcreteProduct
func (f ConcreteFactory1) Produce() Product {
	product := ConcreteProduct1{}
	return product
}
