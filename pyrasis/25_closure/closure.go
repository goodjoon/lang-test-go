package main

import "fmt"

func main(){
	a, b := 3, 5

	f := func(x int) int {
		return a*x+b // use a, b which declared outside of this functino
	}

	y := f(5)

	fmt.Println(y)

	//---

	f = calc()

	fmt.Println("==========")
	fmt.Println(f(10))
	fmt.Println(f(20))
}

func calc() func(int) int { // define `calc()` function returns function
	a, b := 3, 5
	return func(x int) int{
		return a*x+b
	}
}