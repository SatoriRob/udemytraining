package main

import "fmt"

func main() {
	var snum int64
	var lnum int64
	fmt.Print("Enter a small number:")
	fmt.Scan(&snum)
	fmt.Print("Enter a larger number:")
	fmt.Scan(&lnum)
	x := lnum / snum
	y := lnum % snum
	fmt.Println(x, "R=", y)
}
