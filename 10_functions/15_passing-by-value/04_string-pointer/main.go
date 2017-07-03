package main

import "fmt"


func main() {

	name := "Robert"
	fmt.Println(&name) //mem add

	changeMe(&name)

	fmt.Println(&name) // mem add
	fmt.Println(name)//Robert
}

func changeMe(z *string) {
	fmt.Println(z) // mem add
	fmt.Println(*z)
	*z = "Todd"
	fmt.Println(z) // mem add
	fmt.Println(*z) // Todd
}

