package main

import "fmt"

type Person struct {
	name string
	surname string
	age int
}

type Student struct {
	name string
	surname string
	age int
}

func (p Person) GetName() string {
	return p.name
}

func (p *Person) SetName(newName string) {
	p.name = newName
}



func main() {
	p := Person{
		name: "Person",
		surname: "Mr",
		age: 10,
	}
	fmt.Println(Student(p))
}