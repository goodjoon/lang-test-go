package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Num of CPU Cores : ", runtime.NumCPU())
	fmt.Println("Num of Goroutine : ", runtime.NumGoroutine())

	fmt.Println("Initial CPU Use : ", runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Current CPU Use : ", runtime.GOMAXPROCS(0))

	s := "Hello, Good Joon"

	for i := 0 ; i < 100 ; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}
