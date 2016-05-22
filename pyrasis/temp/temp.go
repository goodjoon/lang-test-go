package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	val1 := 100

	rand.Seed(time.Now().UnixNano())
	switch swVal := val1 + rand.Intn(10); { // 조건문 안에서의 함수호출은 ; 필수!
	case swVal > 100 && swVal < 200:
		fmt.Println("100 보다 크고 200보다 작음")
	case swVal >= 200 :
		fmt.Println("200 보다 크거나 같음")
		fallthrough
	case swVal == 100:
		fmt.Println("110!!")
	default:
		fmt.Println("많이 큼")
	}
}
