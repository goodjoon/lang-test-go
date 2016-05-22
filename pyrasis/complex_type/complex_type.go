package main

func main() {

	var val0 complex128 = 1.123e-4 + 2.234e+3i // 실수와 허수 모두 입력
	var val1 complex128 = 2.234e+3i
	var val2 complex128 = 2.234e+3i                    // 허수부만 입력
	var val3 complex128 = complex(1.123e-4, 2.234e+3i) // complex() 로 입력

	var rVal1 float64 = real(val1) // 실수부
	var iVal1 float64 = imag(val1) // 허수부

}
