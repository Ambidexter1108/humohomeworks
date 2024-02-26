package funcs

import (
	"fmt"
)


func Pythagoras(){

fmt.Println("Таблица умножения от 1 до 10")
			for a :=1; a <=10; a++{
				for b :=1; b <=10; b++{
					fmt.Printf("%dx%d=%d\n", a, b, a*b)
				}
			fmt.Println()
			}

}




	


func Fibonacci (num int) int {
	if num <= 1 {
		return num
	}
	return Fibonacci(num-1)+Fibonacci(num-2)
}

