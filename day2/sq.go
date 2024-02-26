package main

/*находим периметр квадрата*/
import "fmt"

func main() {
	
	var a int 
	
	fmt.Println("Enter side a")
	fmt.Scanln(&a)
	
	perimeter := 4 * a
	fmt.Println("The perimeter of the square is", perimeter)

	fmt.Scan(&a)

}