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

	// map 순회하기
	for key, value := range solarSystem {
		fmt.Println("key : ", key, ", val : ", value)
	}

}
