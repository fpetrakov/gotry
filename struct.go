package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	p := person{"Bob", 20}
	p = person{name: "Alice", age: 30}

	fmt.Println(p)
	fmt.Println(&p)

	j := newPerson("John")
	fmt.Println(j.age)

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	s.age = 51
	fmt.Println(s)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
