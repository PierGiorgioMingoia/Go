package main

import "fmt"

const MAX int = 100

//This is a comment
func main() {
	fmt.Println("Hello world")

	var student1 string = "John"
	var student2 = "Jane"

	x := 2

	fmt.Println(student1, student2, x)

	var i int
	fmt.Scan(&i)
	fmt.Println("Your number is: ", i)

	var a = [3]int{1, 2, 3}

	fmt.Println(a)

	if 20 > 18 {
		fmt.Println("20 is greater then 18")
	} else {
		fmt.Println("20 is greater then 18")
	}

}
