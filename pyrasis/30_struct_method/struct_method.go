package main

import "fmt"

type Rectangle struct {
	width, height int
}

// add `area()` method to Rectangle struct
// `rect` is called as `Receiver`
func (rect *Rectangle) area() int { //Receiver represent struct itself (this)
	return rect.width * rect.height
}

// Receiver as a pointer (reference)
func (rect *Rectangle) scaleA(factor int) {
	rect.width *= factor
	rect.height *= factor
}

// Receiver as a value - no modification occurs
func (rect Rectangle) scaleB(factor int) {
	rect.width *= factor
	rect.height *= factor
}

// dummy method can omit receiver by noting _
func (_ Rectangle) dummy() {
	fmt.Println("I'm a dummy method")
}

func main() {
	rect := Rectangle{100, 200}
	area := rect.area()

	fmt.Println(area)

	//
	fmt.Println("Original ", rect)
	rect.scaleA(100)
	fmt.Println("Pointer receiver ", rect)
	rect.scaleB(200)
	fmt.Println("Value receiver ", rect, "Nothing changed!")

	rect.dummy()
}
