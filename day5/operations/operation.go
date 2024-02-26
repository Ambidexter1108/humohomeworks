package operations

import (
	"fmt"
	"math/rand"
	"school/student"
	"log"
)
func TakeTicket (pupil student.Student) student.Student{

	// Проверка по какому предмету экзамен
	if pupil.Age <= 10 {
		fmt.Println("Экзамен будет по родному языку")
		}else if pupil.Age > 10 && pupil.Age <= 14 {
			fmt.Println("Экзамен будет по истории и математике")	
		} else if pupil.Age > 14 && pupil.Age < 18 {
				fmt.Println("Экзамен будет по английскому и математике")	
		}else if pupil.Age == 18 {
			fmt.Println("Тест будет по родному и английскому языках, математике и истории")	
		}else {
			fmt.Println("Кто ты воин?")
		}
 //Ученик получает билет с вопросами
pupil.TicketNumber = rand.Intn(20)+5
return pupil

}
func StartExam(pupil student.Student) student.Student {
	var (
	
		students int = rand.Intn(20) + pupil.TicketNumber
	)
	
	for i := 1; i <= students; i++ {

		switch i {
	
		case pupil.TicketNumber:
			log.Println("Теперь моя очередь!", i)
			pupil.Requirements.Uniform = true
			pupil.Requirements.Attendance = "Хорошая посещаемость"
			pupil = Exam(pupil)
			return pupil

		default:
			log.Println("Очередной одноклассник сдал экзамен или возможно провалился;) ", i)
		}
	}
	return pupil
}

func Exam(pupil student.Student) student.Student {
	switch {

	case !pupil.Requirements.Uniform:
		fmt.Println("Куда собрался без юниформы?!")
		return pupil
	case pupil.Requirements.Uniform:
		fmt.Println("По поводу одежды претензий нет")
		pupil.Total++

	case pupil.Requirements.Attendance != "Хорошая посещаемость":
		fmt.Println("О хороших оценках можешь даже и не думать!")
		return pupil
	case pupil.Requirements.Attendance == "Хорошая посещаемость":
		fmt.Println("Тебя нужно наградить за посещаемость))")
		pupil.Total++

	}
	pupil.Requirements.GPA = rand.Intn(5)
		switch pupil.Requirements.GPA {
		case 0:
			fmt.Println("Твой GPA:", pupil.Requirements.GPA, "Ты провален.Экзамен тебе не поможет")
			pupil.Total--
		case 1:
			fmt.Println("Твой GPA:", pupil.Requirements.GPA, "Не повезло тебе в жизни")
			pupil.Total--
		case 2:
			fmt.Println("Твой GPA:", pupil.Requirements.GPA, "Удовлетворительно. Дадим тебе шанс")
			pupil.Total++
		case 3:
			fmt.Println("Твой GPA:", pupil.Requirements.GPA, "Хорошо. Есть шанс на пятерку")
			pupil.Total++
		case 4:
			fmt.Println("Твой GPA:", pupil.Requirements.GPA, "Отлично. Молодец!")
			pupil.Total++
		}
			pupil.Total += rand.Intn(4)
			switch {
		case pupil.Total == 1:
			fmt.Println("Не повезло! На пересдачу")
			pupil.Total--
			fmt.Println("Твой общий бал:", pupil.Total)			
			pupil.Requirements.Pass = "Не прошёл"
		case pupil.Total == 2:
			fmt.Println("Удовлетворительно")
			pupil.Total++
			fmt.Println("Твой общий бал:", pupil.Total)
			pupil.Requirements.Pass = "Прошёл"
		case pupil.Total == 3:
			fmt.Println("Хорошо")
			pupil.Total++
			fmt.Println("Твой общий бал:", pupil.Total)
			pupil.Requirements.Pass = "Прошёл"
		case pupil.Total == 4:
			fmt.Println("Отлично. Молодец!")
			pupil.Total++
			fmt.Println("Твой общий бал:", pupil.Total)
			pupil.Requirements.Pass = "Прошёл"
		case pupil.Total > 4:
			fmt.Println("Ты будешь награжден медалью и грамотой")
			pupil.Total++
			fmt.Println("Твой общий бал:", pupil.Total)
			pupil.Requirements.Pass = "Прошёл"	


			//pupil.Requirements.Pass = "Прошёл"
			//fmt.Println(pupil.Requirements.Pass)
			//fmt.Println(pupil.Total)
		
	
	}
	return pupil
}