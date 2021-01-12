package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

// method pointer struct
func (student *Student) graduate() {
	student.Name = student.Name + " S.Kom"
}

func main() {
	student := Student{3216607, "Danil Syah", 3.64}
	fmt.Println("Student Name (value) : ", student.Name)
	fmt.Println("Student Name (address) : ", &student.Name)

	student.graduate()

	fmt.Println("Student Name (value) : ", student.Name)
	fmt.Println("Student Name (address) : ", &student.Name)
}
