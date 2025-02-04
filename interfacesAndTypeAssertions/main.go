package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct {
	Name string
}

type Cat struct{}

func (d *Dog) Sound() string {
	return "Au! Au!"
}

func (d *Cat) Sound() string {
	return "Meaw!"
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Sound())
}

func takeAnimal(t Animal) {
	switch t.(type) {
	case *Dog:
		fmt.Println("This is a dog")
		t.Sound()
	case *Cat:
		fmt.Println("This is a cat")
		t.Sound()
	default:
		fmt.Println("This is an unknown animal")
	}
}

type Snores struct{}

func (Snores) String() string {
	return "Snores stringer method"
}

func main() {
	dog := Dog{Name: "Rex"}
	whatDoesThisAnimalSay(&dog)
	takeAnimal(&dog)
	s := Snores{}
	fmt.Println(s)
}
