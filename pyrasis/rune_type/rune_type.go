package main

import "fmt"

func main() {
	var val1 rune = '굿'
	var val2 = '준'
	var val3 rune = '\uAD7F'
	val4 := '\uC900'

	fmt.Println(val1, " ", val2)
	fmt.Println(val3, " ", val4)
	fmt.Printf("%c%c", val3, val4)
}
