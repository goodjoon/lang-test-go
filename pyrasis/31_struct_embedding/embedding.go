package main

import "fmt"

// Go language does not provide Class and Inherit
// But embedding can give you similar effect

type Person struct {
	name string
	age int
}

func (p *Person) greeting() {
	fmt.Println("Hello~")
}

func (_ Person) whoami() {
	fmt.Println("whoami : I'm a human")
}

type Student struct {
	p Person		// Just include Person struct variable p inside
	school string
	grade int
}

type Teacher struct {
	Person // anonymous field. is-a relation.
	school string
}

func (teacher *Teacher) attention() {
	fmt.Println("ATTENTION!!")
}

func (_ Teacher) whoami() { // OVERRIDING EFFECT !
	fmt.Println("whoami : I'm a TEACHER")
}

func main() {
	fmt.Println("=== struct embedding / include")
	var s Student
	s.p.greeting()
	s.p.whoami()

	fmt.Println("\n=== struct embedding / is-a")
	var t Teacher
	t.greeting()
	t.attention()
	t.whoami()
}
