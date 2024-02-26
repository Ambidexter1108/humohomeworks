package main

//находим периметр треугольника

import "fmt"

func main() {

	var a, b, c int 

	fmt.Println("Enter side a")
	fmt.Scanln(&a)

	fmt.Println("Enter side b")
	fmt.Scanln(&b)

	fmt.Println("Enter side c")
	fmt.Scanln(&c)

	perimeter := a + b + c
	fmt.Println("The perimeter of the triangle is", perimeter)

fmt.Scan(&a)

}