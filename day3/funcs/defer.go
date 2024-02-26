package funcs

import (
	"fmt"
	"log"
)

func Deferr() {

	defer fmt.Println("Вот теперь точно всё!\n_____________________________________")

	fmt.Println("Это отложенная функция")

	log.Println("Последння функция")

	

}