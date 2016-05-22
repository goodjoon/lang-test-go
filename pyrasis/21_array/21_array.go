package main

import "fmt"

func main() {
	var ar1 [10]int
	ar1[2] = 100
	var ar2 [5]int = [5]int{1,2,3,4,5}
	var ar3 = [5]int {1,2,3,4,5}
	var ar4 = [...]int {1,2,3,4,5} // ... 은 초기화 시에 자동으로 사이즈가 설정 됨
	var ar5 [3]int
	ar5 = [...]int {1,2,3}
	//ar5 = []int {1,2,3} ==> 에러
	ar6 := [...]int {1,2,3,4,5}
	var ar7 = [...]int {
		1,
		2,
		3, // ==> 여러줄일때는 반드시 마지막줄에도 콤마 , 를 붙여야 됨
	}
	var ar8 = [2][2]int{
		{1,2},
		{3,4},
	}

	// 배열 정상 확인
	fmt.Println(ar1,"\n",ar2,"\n",ar3,"\n",ar4,"\n",ar5,
				"\n",ar6,"\n",ar7,"\n",ar8)

	// 배열 출력 (for)
	fmt.Println("=== for 출력 ===")
	for i:=0 ; i<len(ar2) ; i++ {
		fmt.Print(ar2[i])
	}
	fmt.Println("\n=== range 를 사용한 출력 ===")
	for i, value := range ar2 { // 첫번째 변수(i) 에는 인덱스, 두번째 변수(value)에는 값이 들어감
		fmt.Println("[",i,"]", value)
	}
	fmt.Println("\n=== _ 로 range 의 다중리턴값 무시 ===")
	for _, value := range ar2 {
		fmt.Print(value) // i 없이 값만 받으려면 _ 를 i 대신에 넣어줌
	}

	// Go 에서 배열 변수는 배열 전체이며 포인터가 아님. 그래서 값 복사 일어남
	fmt.Println("\n=== 배열 복사 ===")
	ar9 := [5]int {1,2,3,4,5}
	ar10 := ar9
	ar10[1] = 100
	fmt.Println(ar10)
	fmt.Println(ar9)

}
