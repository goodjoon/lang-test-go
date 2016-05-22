package main

import "fmt"

func main() {
	func() {
		fmt.Println("anonymous function")
	}()

	func (s string) {
		fmt.Println(s)
	}("immediate call")

	r := func(a int, b int) int {
		return a+b
	}(100,200)

	fmt.Println(r)

	// function variable
	sum := func (a int, b int) int {
		return a+b
	}

	v := sum(1,2)
	fmt.Println(v)

}
