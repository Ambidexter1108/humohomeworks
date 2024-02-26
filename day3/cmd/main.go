package main

import (
	"fmt"
	"log"

	"adi.das/funcs"
)

func main() {

defer funcs.Deferr()

fmt.Println("__________________________________________________")
log.Println("Программа запущена")

funcs.HelloWorld()


log.Println("Сложение чисел:", funcs.Plus(10, 20))

funcs.Minus(20, 10)

log.Println("Квадрат числа:", funcs.Square(10))

funcs.Farewell()

log.Println("Программа завершена")

fmt.Println("__________________________________________________")


}