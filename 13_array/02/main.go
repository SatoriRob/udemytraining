package main

import "fmt"

func main() {
	var x [58]string

	for i :=65; i<= 122; i++ {
		x[i-65] = string(i) //turns i into letter(ASCII), stores it in array
	}
	fmt.Println(x)
	fmt.Println(x[42])
}
