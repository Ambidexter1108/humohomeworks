package main 

import ("fmt";
				"plus.go"; 
				"minus.go"; 
				"div.go"; 
				"mult.go") 
func main() {
	fmt.Println(
		plus(10,5),
		minus(10,5),
		div(10,5),
		mult(10,5))
}