package main

import ("fmt";"log";"super/funcs"; "sept/funcs")




func main() {
	
fmt.Println("-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|")	
log.Println("Программа запущена>>>>>")

fmt.Println("Здравствуйте! Пожалуйста, выберите день обучения:\n",
"- 1 День. Знакомство\n",
"- 2 День. Первая программа\n",
"- 3 День. Типы, фукнции и модули\n",
"- 4 День. Условия и циклы\n",
"........",
)

var day int
fmt.Scanln(&day)

switch {
case day == 0:
	fmt.Println("Такого дня не бывает...\nТак что переходим на следующий день!")
	fallthrough

// first day
case day == 1:
	fmt.Println("Опаньки!Домашнего задания не было;)\nПерезапустите программу")
	
//second day
case day == 2:
			fmt.Println("Выберите задание:",
			"Задание 1: Программа, которая находит периметр квадрата",
			"Задание 2: Программа, которая находит периметр прямоугольника",
			"Задание 3: Программа, которая находит площадь прямоугольника",
			"Задание 4: Программа, которая будет вычислять, сколько пользователю ещё нужно лет до его юбилея")

		
			var task2 int
			fmt.Scanln(&task2)
			if task2 == 1 {
				funcs.Kvadrat()
			} else if task2 == 2 {
				funcs.RectanglePerimeter()
			} else if task2 == 3 {
				funcs.RectangleArea()
			}else if task2 == 4 {
				funcs.Anniversary()
			}else {
				fmt.Println("Такого задания нет!!!")
			}	

//day 3	
case day == 3:
	fmt.Println("Выберите задание:",
	"\nЗадание 1: Блоки кода разбить на функции",
	"\nЗадание 2: Распределить код из main.go на пакеты и отдельные Go файлы. А в main только вызвать существующие функции. Создать модуль.")
			var task3 int
			fmt.Scanln(&task3)
			if task3 == 1 {
				funcs.HelloWorld()				
				funcs.Farewell()
				funcs.Deferr()
		
	
			} else if task3 == 2 {
							plus := funcs.Plus(15, 20)
							fmt.Println(plus)
							
							funcs.Minus(10, 20)
			}else {
				fmt.Println("Такого задания нет!!!")
			}	

//day 4
case day == 4:
	fmt.Println("Выберите задание:",
	"\nЗадание 1: Объеденить предыдущие задания",
	"\nЗадание 2(бонусное): При помощи циклов вывести в терминал таблицу умножения",
	"\nЗадание 3(бонусное): Используя рекурсию написать функцию, вычисляющую числа Фибоначчи")
	var task4 int
	fmt.Scanln(&task4)	
	if task4 == 1 {
		fmt.Println("Всё сделано!!!")
	} else if task4 == 2 {
		funcs.Pythagoras()
	} else if task4 == 3 {
						var input int
						fmt.Println("Введите номер числа Фибоначчи")
						fmt.Scanln(&input)
						result := funcs.Fibonacci(input)
						fmt.Printf("Число Фибоначчи для %d: %d\n", input, result)
						
		
	}	else {
		fmt.Println("Такого задания нет!!!")
	}

//day 7	
case day == 7:
	fmt.Println("Выберите задание:", "\nЗадание 1: Напишите функцию, которая отзеркалит массив наоборот. К примеру, на полученный массив строк Hello, how, are, you, ? вернётся ответ ?, you, are, how, Hello. Функция должна работать с любыми типами одинаково.",
	"\nЗадание 2: Напишите функцию, которая будет разбивать один массив с числами на два слайса: В первом слайсе будут чётные числа из главного массива, а во втором нечётные числа из главного массива")
		
	var task7 int
	if task7 == 1 {
		funcs.Reflect()	
	} else{
		
			numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			chet, nechet := funcs.Divide(numbers)
			fmt.Println("Четные числа:", chet)
			fmt.Println("Нечетные числа:", nechet)
}
	}
		
		

		





	








log.Println("<<<<<Программа завершена")
fmt.Println("-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|")	

}