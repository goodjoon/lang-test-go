package main

import "fmt"

func main() {
	defer world() // call `world()` function at last of function code execution
	hello()
	hello()
	hello()

	// defer call second anonymouse function
	// last defer call would be called first within defer calls
	defer func() {
		fmt.Println("Joon")
	}()

	f := func() {
		fmt.Println("Good")
	}

	f()
	f()
}

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("World")
}