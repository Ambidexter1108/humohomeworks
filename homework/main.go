package main

import "fmt"



func plus (x, y int) int {
	return x+y
}
func minus (x, y int) int {
	return x-y
}
func mult (x, y int) int {
	return x*y
}
func div (x, y int) int {
	return x/y
}



func main (){
	fmt.Println(
	plus(10,5),
	minus(10,5),
	mult(10,5),
	div(10,5))
}