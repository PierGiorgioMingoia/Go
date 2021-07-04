package main

import "fmt"

const MAX int = 100

func myMessage(fName string) {
	fmt.Println("Hello", fName, "Ciao")
}

func myFunction(x int, y int) int {
	return x + y
}

func multipleReturn(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

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
		fmt.Println("20 is less then 18")
	}

	switch a[0] {
	case 1:
		x = 5
	case 2:
		x = 6
	default:
		x = 0
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println(factorial_recursion(10))

}
