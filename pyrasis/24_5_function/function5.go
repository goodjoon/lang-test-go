package main

import "fmt"

// Assigning function to a variable
func main() {

	// variables can contain function
	var hello func(a int, b int) int = sum // declare `hello` with function definition
	world := sum // automatically assign `sum` function to `world` variable

	fmt.Println(hello(100,200))
	fmt.Println(world(500,600),"\n------ function in array")

	// also array can contain function
	funcArray := []func(int, int) int {sum, diff}
	for _, function := range funcArray {
		fmt.Println(function(500, 200))
	}
	fmt.Println("------- function in map")

	// same for map
	var funcMap = make(map[string]func(int, int) int)
	funcMap["sum"] = sum
	funcMap["diff"] = diff
	for _, function := range funcMap {
		fmt.Println(function(80,20))
	}
}

func sum(a int, b int) int {
	return a+b
}

func diff(a int, b int) int {
	return a-b
}