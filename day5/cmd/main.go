package main

import (
	"log"
	"math/rand"
	"school/operations"
	"school/student"
)

func main() {
	var Student student.Student = student.Student{Name: "Albert", Surname: "Einstein", Age: rand.Intn(12) + 7, TicketNumber: 0, Requirements: student.Requirements{Uniform: true, Attendance: "хорошая посещаемость", GPA: 0, Pass: "Нужно пройти экзамен"}}

	log.Println("Привет! Я пришел сдать экзамен.")

	Student = operations.TakeTicket(Student)
	log.Println("Я забрал билет. \nНомер моего билета:", Student.TicketNumber)
	log.Println("Экзамен начался")
	Student = operations.StartExam(Student)
	log.Println("Экзамен закончился. Вот результаты:")
	if Student.Requirements.Pass != "Прошёл" {
		log.Printf("Увы ты остался на один год:( %+v", Student)
	} else {
		log.Printf("Поздравляю и новых успехов в будущем! %+v", Student)
	}

}
