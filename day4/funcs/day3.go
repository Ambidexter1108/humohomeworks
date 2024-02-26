package funcs

import (
	"fmt"
	"log"
)

func HelloWorld() {
	fmt.Println("Hello World!")
}

func Farewell() {
	fmt.Println("The end")
}

func Deferr() {

	defer fmt.Println("Вот теперь точно всё!\n_____________________________________")

	fmt.Println("Это отложенная функция")

	log.Println("Последння функция")

}

func Minus (x, y int) {
	log.Println("Результат вычитания:", x-y)
}	



func Plus (x, y int64) int64 {
		return x + y

}

func Square(num int8) int8 {
	return num * num
}