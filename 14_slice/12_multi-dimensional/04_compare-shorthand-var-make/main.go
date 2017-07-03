package main

import"fmt"

func main() {
	student := []string{}
	students := [][]string{}
	student[0] = "Todd" //this results in an error
	//student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
	}
