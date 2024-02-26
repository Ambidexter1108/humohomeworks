package funcs


import (
	"fmt"
	"log"
	"math"
)

func Kvadrat () {
	
	var a int 
	
	fmt.Println("Enter side a")
	fmt.Scanln(&a)
	
	perimeter := 4 * a
	fmt.Println("The perimeter of the square is", perimeter)

	fmt.Scan(&a)

}

func RectangleArea() {

	var a, b int 
	
			fmt.Println("Enter side a")
			fmt.Scanln(&a)
	
			fmt.Println("Enter side b")
			fmt.Scanln(&b)

			area := a * b 
			fmt.Println("The area of the rectangle is", area)

			fmt.Scan(&a)

}

func RectanglePerimeter() {

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

func Anniversary () {

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