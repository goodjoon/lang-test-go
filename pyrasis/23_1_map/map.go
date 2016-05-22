package main

import (
	"fmt"
)

func main() {
	solarSystem := make(map[string]float32)
	solarSystem["Mercury"] = 87.969
	solarSystem["Venus"] = 224.7069
	solarSystem["Earth"] = 365.25641
	solarSystem["Mars"] = 686.9600

	fmt.Println("Earth :", solarSystem["Earth"])
	fmt.Println("잘못된값 :", solarSystem["xxx"]) // 잘못된 키로 조회하면 빈 값이 리턴된다

	// 값이 있는지 검사하려면 리턴값 사용한다
	value, ok := solarSystem["xxx"]
	fmt.Println("value : ", value, " ok : ", ok)
	value, ok = solarSystem["Mars"]
	fmt.Println("value : ", value, " ok : ", ok)

	// if 문으로 결합하면
	if value, ok := solarSystem["Mercury"]; ok {
		fmt.Println("value : ", value, " ok : ", ok)
	}
}
