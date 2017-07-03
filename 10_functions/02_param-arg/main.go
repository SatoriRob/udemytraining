package _2_param_arg

import "fmt"

func main() {
	greet("Jane")
	greet("John")
}

func greet(name string) {
	fmt.Println(name)
}

//greet is declared with parameter
//when calling greet, pass an argument