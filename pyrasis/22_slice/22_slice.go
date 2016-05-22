package main

import "fmt"

/**
	slice 는 배열과 같지만 길이가 정해지지 않고 동적으로 늘어남
	배열과 다르게 Reference 타입임
	make 함수로 공간 할당해야 값을 넣을 수 있음

	make([]<타입>, 길이)
 */
func main() {
	var ar1 []int // Slice 선언. 현재 길이는 0
	ar1 = make([]int, 3)
	var ar2 []int = make([]int, 3)
	var ar3 = make([]int, 3)
	ar4 := make([]int, 3)

	var ar5 = []int{1,2,3}
	var ar6 = []int{
		1,
		2,
		3,
	}

	fmt.Println(ar1, ar2, ar3, ar4, ar5, ar6);

	// make 에 Buffer Capacity 적용. 크면 성능 좋으나 메모리 초기부터 많이먹음
	var ar7 = make([]int, 3, 5) // 마지막 파라미터는 Capacity 이며 부족한 경우 Buffer 사이즈가 늘어남
	// ar[4] ==> Capacity 가 length 보다 크더라도 length 가 현재 3 으로 잡혀있고 0으로 3개만 초기화 되어있음. Index out of range 에러 발생함
	fmt.Println("len(ar8) : ", len(ar7), "/ cap(ar8)", cap(ar7))
	ar7 = []int {1,2,3,4,5,6,7,8,9}
	fmt.Println("len(ar8) : ", len(ar7), "/ cap(ar8)", cap(ar7))
	fmt.Println(ar7)

	ar8 := []int {1,2,3}
	fmt.Println(ar8)

	// APPEND
	ar8 = append(ar8, 4,5,6)
	fmt.Println(ar8)
}
