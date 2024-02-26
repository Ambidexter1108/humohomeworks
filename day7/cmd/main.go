package main
import ("fmt"; "sept/funcs")

func main() {


funcs.Reflect()


	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chet, nechet := funcs.Divide(numbers)
	fmt.Println("Четные числа:", chet)
	fmt.Println("Нечетные числа:", nechet)
}