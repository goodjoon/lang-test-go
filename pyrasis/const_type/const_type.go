package main

import "fmt"

const CONST_NUM = 100
const CONST_STR string = "굿준"
const (
	RED   int = 0
	GREEN int = 1
	BLUE  int = 2
)
const (
	ZERO = iota // 0 부터 시작하며 하나씩 증가
	ONE
	TWO
	THREE
	FOUR
)
const (
	HANA = 0
	DUL
	SET
)

func main() {
	fmt.Println(CONST_STR)
	fmt.Println(GREEN)
	fmt.Println(THREE)
	fmt.Println(DUL)
}
