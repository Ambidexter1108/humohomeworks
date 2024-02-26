package student


type Requirements struct {
	Uniform bool
	Attendance string
	GPA int
	Pass string
}





type Student struct {
	Name    string
	Surname string
	Age     int
	TicketNumber   int
	Requirements Requirements
	Total int
}
