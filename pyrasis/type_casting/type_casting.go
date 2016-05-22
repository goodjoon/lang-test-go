package main

import "fmt"

func main() {
	var val1 int64 = 1000000
	var val2 float32 = 1.234e-3

	var res = float32(val1) + val2

	fmt.Println(res)
}
