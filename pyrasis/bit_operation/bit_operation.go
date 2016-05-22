package main

import "fmt"

func main() {
	var val1 byte = 0x70 // 0111 0000
	var val2 byte = 0x11 // 0001 0001
	var val3, val4 = 0x70, 0x11

	fmt.Println("==== Bit Operator ====")
	fmt.Printf("%08b\n", val1^val2) // XOR 연산
	fmt.Printf("%08b\n", ^val1)
	// &^ : 좌항과 XOR 연산 값을 다시 좌항과 AND 연산 함 (좌항 값의 Bit Clear)
	fmt.Printf("%08b\n", val1&^val2)
	fmt.Printf("%08b\n", (val1^val2)&val1) // &^ 과 같음
	val3 &^= val4
	fmt.Printf("%08b\n", val3)

	// Pointer 관련 연산자
	var val5 = 10
	var ref1 = &val5

	fmt.Println("==== Pointer ====")
	fmt.Printf("%d\n", *ref1)
	fmt.Printf("val5's address : %x / ref1's value : %x", &val5, ref1)

}
