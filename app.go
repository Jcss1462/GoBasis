package main

import (
	"fmt"
)

// Definimos una interfaz
type Animal interface {
	Speak() string
}

// Estructura que implementa la interfaz
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Otra estructura que implementa la interfaz
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {

	var animal Animal

	dog := Dog{}
	cat := Cat{}

	animal = dog
	fmt.Println("Dog says:", animal.Speak())

	animal = cat
	fmt.Println("Cat says:", animal.Speak())

}
