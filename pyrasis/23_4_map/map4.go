package main

import (
	"fmt"
)

func main() {
	// map 안에 map 만들기

	terrestrialPlanet := map[string]map[string]float32{
		"Mercury": map[string]float32{
			"meanRadius":    2439.7,
			"mass":          3.302E+23,
			"orbitalPeriod": 87.969,
		},
		"Venus": map[string]float32{
			"meanRadius":    6051.8,
			"mass":          4.444E+23,
			"orbitalPeriod": 99.123,
		},
	}

	fmt.Println(terrestrialPlanet["Venus"]["mass"])
}
