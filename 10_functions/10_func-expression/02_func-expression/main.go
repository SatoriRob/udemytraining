package main

import "fmt"

func greeting() {
	greeting := func() {
		fmt.Println("Hello world!")
	}

	greeting()
}



