package main

import "fmt"



func main() {
	n := foo(1, 2)
	m := foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	p := foo(aSlice...)
	q := foo()
	fmt.Println(n, m, p, q)
}

func foo(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
		}
		return largest
}

