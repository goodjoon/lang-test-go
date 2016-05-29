package main

import "fmt"

type Duck struct {
}

func (d Duck) quack() {
	fmt.Println("Quack!~")
}

func (d Duck) feathers() {
	fmt.Println("Duck has white and grey feathers")
}

type Person struct {
}

func (p Person) quack() {
	fmt.Println("Person sounds QUUAK")
}

func (p Person) feathers() {
	fmt.Println("Person picks up a feather and shows it up")
}

type Quacker interface {
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}

func main() {
	var donald Duck
	var john Person

	inTheForest(donald)
	inTheForest(john)

	// check if an instance implements specific interface
	// by `interface{}(<instance>).(<interface>)
	if v, ok := interface{}(donald).(Quacker); ok {
		fmt.Println(v, ok)
	}
}
