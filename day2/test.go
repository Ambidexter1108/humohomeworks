package main

/*сколько нужно лет до юбилея
решение 2*/

import (
	"fmt"
	"log"
	"math"
)


func main() {

	log.Println("\n___________ Программа запущена ___________")
			var name string
			fmt.Println("Как вас зовут?")
			fmt.Scanln(&name)
			fmt.Println("Добро пожаловать", name, "!")

	var a int 
	
	fmt.Println("Введите год вашего рождения: ")
	fmt.Scan(    &a     )
	
age := (2024 - a)
	fmt.Println(name, "Вам",age,"лет/год(а)")

	c := float64(age) / 10
  d := math.Ceil(c) * 10

cycle := int(d) - age

	fmt.Println("Ваш Юбилей через", cycle, "лет/год(а)")

	log.Println("\n___________ Программа завершена ___________")

	fmt.Scan(&a)

}