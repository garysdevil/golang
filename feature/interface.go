package feature

import "fmt"

type Mammal interface {
	Say()
}

type Dog struct{}

type Cat struct{}

type Human struct{}

func (d Dog) Say() {
	fmt.Println("woof")
}

func (c Cat) Say() {
	fmt.Println("meow")
}

func (h Human) Say() {
	fmt.Println("speak")
}

func Interface() {
	var m Mammal
	m = Dog{}
	m.Say()
	m = Cat{}
	m.Say()
	m = Human{}
	m.Say()

	// m1 := Dog{}
	// m1.Say()
	// m2 := Cat{}
	// m2.Say()
	// m3 := Human{}
	// m3.Say()
}
