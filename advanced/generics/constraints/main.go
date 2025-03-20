package main

import "fmt"

type MyType string

func (m MyType) Foo() {
	fmt.Println("My type")
}

func main() {
	var mt MyType = "Teste"
	mt.Foo()
	foo(mt)
	var someString MyTypeWithStringCoreType = "Teste 123..."
	myTypeFoo(someString)
}

type MyTypeWithStringCoreType string

type MyConstraint interface {
	Foo()
}

type MyConstraintWithSomeTypes interface {
	int | ~string | []int
}

func foo[T MyConstraint](arg T) {
	fmt.Println(arg)
}

func myTypeFoo[T MyConstraintWithSomeTypes](arg T) {
	fmt.Println(arg)
}
