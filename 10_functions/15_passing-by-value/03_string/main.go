package main

import "fmt"


func main() {

	name := "Robert"
	fmt.Println(name) //Robert

	changeMe(name)

	fmt.Println(name) //Robert
}

func changeMe(z string) {
	fmt.Println(z)// Robert
	z = "Todd"
	fmt.Println(z) // Todd
}
// strings are immutable, but a new string
//can be assigned