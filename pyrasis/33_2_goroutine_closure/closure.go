package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	s := "Hello, Good Joon"

	for i := 0 ; i < 100 ; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}
