package main

import "fmt"

func hello() {
	fmt.Println("Hello, Joon")
}

func main() {
	go hello() // run `hello()` function as go routine. Go routine consumes much smaller resources of O/S.
	fmt.Scanln()
}
