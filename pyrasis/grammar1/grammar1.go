package main

import "fmt"

func main() {
	some := 100
	if some == 100 {
		fmt.Println("백 ")
		fmt.Println("입니다")
	} else {
		fmt.Println("뭥미")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("==> ", i)
	}
}
