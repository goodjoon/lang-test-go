package main

import (
	"fmt"
)

// 함수 사용
func hello() {
	fmt.Println("Hello, World")
}

func sum(a int, b int) int {
	return a + b
}

// 함수 리턴에 변수명을 주면 리턴용 변수로 지정할 수 있다
func multiply(a int, b int) (ret int) {
	ret = a * b
	return // 리턴값이 있는데 명시적으로 return 안쓰면 에러난다
}

func main() {
	hello()

	var retVal = sum(100,200)
	fmt.Println("sum : ", retVal)

	retVal = multiply(200,300)
	fmt.Println("multiply : ", retVal)
}
