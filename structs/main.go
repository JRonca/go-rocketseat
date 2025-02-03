package main

import (
	"encoding/json"
	"fmt"
)

type MyString string

type Foo struct {
	Bar string
	Baz int
}

func (Foo) FooMethod() {
	fmt.Println("Foo Method")
}

type User struct {
	Foo
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

func (User) SimpleMethod() {
	fmt.Println("Simple Method")
}

func (u User) PrintName() {
	fmt.Println(u.Name)
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func main() {
	user := User{}
	otherUser := User{Foo{}, "José Renato", 1}
	userWithFields := User{Name: "José Renato Ronca", ID: 2}
	userWithAlsoName := User{Name: "Snores"}

	fmt.Println(user)
	fmt.Println(otherUser)
	fmt.Println(userWithFields)
	fmt.Println(userWithAlsoName)

	fmt.Println(userWithFields.Name)
	fmt.Println(otherUser.ID)

	user.SimpleMethod()
	userWithAlsoName.PrintName()
	userWithAlsoName.UpdateName("John Doe")
	userWithAlsoName.PrintName()

	user.FooMethod()
	fmt.Println(user.Bar)

	res, err := json.Marshal(userWithFields)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}
