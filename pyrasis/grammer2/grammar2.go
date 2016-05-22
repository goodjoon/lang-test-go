/**
변수 선언 예제
*/
package main

import _ "fmt"

func main() {
	// 변수 선언 예제
	var val1 = "1"
	_ = val1
	var val2 int
	var val3 int = 3
	val4 := 0x04

	// 여러 변수 동시 선언
	var val5, val6 = "5", 6
	var val7, val8 int
	var val9, val10 int = 9, 10
	val11, val12 := 11, "12"

	var (
		val13, val14 int = 13, 14
		val15, val16 = "15", 16
	)
}
