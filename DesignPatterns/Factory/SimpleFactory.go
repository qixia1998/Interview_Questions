package Factory

import "fmt"

type Person interface {
	Greet()
}

type person struct {
	Name string
	Age  int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s", p.Name)
}

func NewPerson(name string, age int) Person {
	return person{
		Name: name,
		Age:  age,
	}
}
