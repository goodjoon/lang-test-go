package main

import (
	"fmt"
)

func main() {
	a := map[string]int{"hello": 10, "world": 20}

	//map 삭제하기
	fmt.Println(a)
	delete(a, "world")
	fmt.Println(a)
}
