package main

//находим периметр прямоугольника

import "fmt"

func main() {

	fmt.Println("Enter side a")
	var a int
	fmt.Scanln(&a)

	fmt.Println("Enter side b")
	var b int
	fmt.Scanln(&b)

	perimeter := 2 * (a+b)
	fmt.Println("The perimeter of the rectangle is",  perimeter)

	fmt.Scanln(&a)
	
}