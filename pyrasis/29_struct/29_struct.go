package main

import (
	"fmt"
)

type Rectangle struct {
	width, height int
}

func main() {
	var rect Rectangle
	var rect1 *Rectangle   // using pointer
	rect1 = new(Rectangle) // use new() function to allocate struct instance
	rect2 := new(Rectangle)

	// struct field initialization
	var rect3 Rectangle = Rectangle{10, 20}
	rect4 := new(Rectangle)
	rect4.width = 30
	rect4.height = 40
	rect5 := Rectangle{width: 50, height: 60}

	fmt.Println(rect, rect1, rect2, rect3, *rect4, rect5)
}
