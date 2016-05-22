package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str1 string = "goodjoon.tistory.com "
	var str2 string = `굿준 블로그`
	var str3 string = `
어서 오세요
환영합니다` // 여러줄을 사용

	fmt.Println(str1, " - ", str2, str3)
	fmt.Println("len str1 : ", len(str1)) // len() 을 사용하면 문자열 전체 바이트 수 가져옴
	fmt.Println("len str2 : ", len(str2))
	// utf8 패키지 사용하여 utf8 문자열의 문자 수 카운트
	fmt.Println("RuneCountInString(str1) : ", utf8.RuneCountInString(str1))
	fmt.Println("RuneCountInString(str2) : ", utf8.RuneCountInString(str2))

	fmt.Println(" + : ", str1 + " - " + str2) // + 연산자
	fmt.Println(" == ", str2 == "굿준 블로그", " / ", str1 == str2) // Go 는 == 로 문자열 비교함
	fmt.Println("str1[1]", str1[1]) // 배열로 문자열 []byte 를 접근함. 수정은 불가함 read-only
	fmt.Println("str2[1]", str2[1])
	for i:=0 ; i<len(str2) ; i++ {
		fmt.Printf("%x|",str2[i])
	}
}
