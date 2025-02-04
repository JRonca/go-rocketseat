package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct {
	Name string
}

func (d *Dog) Sound() string {
	if d == nil {
		return "<nil>"
	}
	fmt.Println("Dog's name is", d.Name)
	return "Au! Au!"
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Sound())
}

func main() {
	dog := Dog{}
	whatDoesThisAnimalSay(&dog)

	var a Animal
	var oneDog *Dog
	a = oneDog
	fmt.Println(a.Sound())
}
