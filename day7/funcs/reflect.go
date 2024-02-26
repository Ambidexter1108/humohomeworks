package funcs
import "fmt"

func Reflect() {
	
	massiw:= [...]int{1, 2, 3, 4, 5}
	fmt.Println("Массив с типом чисел:", massiw)
	
fmt.Println(massiw[4],massiw[3],massiw[2],massiw[1],massiw[0])



massive :=[...]string{"Hello", "how", "are", "you","!"}

fmt.Println("Массив с типом строк:", massive)	
fmt.Println(massive[4],massive[3],massive[2],massive[1],massive[0])

}