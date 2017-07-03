package main

import (
	"fmt"
)

func main() {
	n := greatest(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func greatest(numbers ...float64) float64 {
	var largest float64
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}
