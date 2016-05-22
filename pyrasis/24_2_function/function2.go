package main

import (
	"fmt"
)

func main() {
	sum, diff := SumAndDiff(200, 100)
	fmt.Printf("Sum : %d, Diff : %d \n", sum, diff)

	_, div := MultAndMod(203, 8)
	fmt.Printf("Mod : %f \n", div)
}

// Returning multiple values
func SumAndDiff(a int, b int) (int, int) {
	return a+b, a-b
}

func MultAndMod(a int, b int) (mult int, div float32) {
	mult = a*b
	div = float32(a)/float32(b)
	return
}