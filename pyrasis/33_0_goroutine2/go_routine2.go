package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r := rand.Intn(100) // generate random number
	time.Sleep(time.Duration(r)) // sleep for number of nanoseconds , int64 type
	fmt.Println(n)
}

func main() {
	for i := 0 ; i < 100 ; i++ {
		go hello(i) // execute `hello()` as go routine which each sleeps for random nano seconds
	}

	fmt.Scanln()
}