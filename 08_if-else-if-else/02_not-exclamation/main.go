package main

import "fmt"

func main() {

	if !true {
		fmt.Println("This did not ran")
	}

	if !false {
		fmt.Println("This ran")
	}
}
