package main

import "fmt"

func main() {
	// Constructor-like instantiation
	rect := NewRectangle(100, 200)
	fmt.Println(rect)

	rect1 := &Rectangle{300, 400}
	fmt.Println(rect1)
}

type Rectangle struct {
	width, height int
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}
}
