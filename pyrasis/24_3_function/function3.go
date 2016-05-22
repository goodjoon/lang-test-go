package main

import "fmt"

func main() {
	r := sum(1,2,3,4,5) // pass multiple arguments to function
	fmt.Println(r)

	ar := []int {10,11,12,13,14,15}
	r = sum(ar...) // pass array as multiple arguments by "...".
	fmt.Println(r)
}

func sum(n ... int) int { // function with variable arguments
	total := 0
	for _, value := range n {
		total += value
	}

	return total
}
