package main

import "fmt"

type Rectangle struct {
	width, height int
}

// get pointer of Rectangle and return it's area
// call by Reference!
func rectangleArea(rect *Rectangle) int {
	return rect.width * rect.height
}

func rectangleScaleA(rect *Rectangle, factor int) {
	rect.width *= factor
	rect.height *= factor
}

func rectangleScaleB(rect Rectangle, factor int) {
	rect.width *= factor
	rect.height *= factor
}

func main() {
	rect := Rectangle{100,200}
	area := rectangleArea(&rect)
	fmt.Println(area)

	// call by reference and value test
	fmt.Println("Original ", rect)
	rectangleScaleA(&rect, 100)
	fmt.Println("Call By Reference ", rect)
	rectangleScaleB(rect, 500)
	fmt.Println("Call By Value ", rect, "Not Changed!")
}


