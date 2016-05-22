package main

import (
	"fmt"
)

// 부분 슬라이스 만들기
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3] // 0으로 시작하는 1번 index 부터, 1로 시작하는 3번 index 까지
	// Slice 는 복사 하지 않고 값을 참조만 한다~
	b[1] = 999

	var c []int
	c = a[0:len(a)]
	d := a[:3]
	e := a[0:3:len(a)] // 0 ~ 2 까지 가져와서 capacity 가 a 의 크기 만큼 인 Slice 를 만듦

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("e : ", e, "cap : ", cap(e))
}
