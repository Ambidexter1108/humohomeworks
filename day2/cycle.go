package main

//сколько нужно лет до юбилея
//решение 1

import (
	"fmt"
	"log"
)


func main() {

	
	log.Println("Программа запущена")

	var age int 
	
	fmt.Println("Введите ваш возраст")
	fmt.Scanln(&age)
	




cycle := (10 - age % 10)%10

	fmt.Println("Ваш юбилей через", cycle, "лет/год(а)")

	log.Println("Программа завершена")

	fmt.Scan(&age)

}