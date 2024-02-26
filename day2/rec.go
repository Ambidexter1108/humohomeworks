package main

//находим площадь прямоугольника

import "fmt"

func main() {

	var a, b int 
	
			fmt.Println("Enter side a")
			fmt.Scanln(&a)
	
			fmt.Println("Enter side b")
			fmt.Scanln(&b)

			area := a * b 
			fmt.Println("The area of the rectangle is", area)

			fmt.Scan(&a)

}