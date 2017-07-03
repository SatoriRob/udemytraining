package main

import "fmt"

func main () {
	age := 44
	fmt.Println(&age) // memory address of age

	changeMe(&age)

	fmt.Println(&age)// mem. add. of age
	fmt.Println(age) // 24
}

func changeMe(z int) {
	fmt.Println(z) // mem. add. of age/z
	fmt.Println(*z) //44
	*z = 24
	fmt.Println(z) //mem. add. of age/z
	fmt.Println(*z) //24
}

