package main

import "fmt"

func main() {

	score := 52
	var grade string
	age := 17
	isMarried := true


	if age > 17 {
		fmt.Println("Sudah dewasa")
	}else{
		fmt.Println("Masih Bocah")
	}

	if score >= 80 && score <= 100{
		grade = "A+"
	}else if score >= 70 && score <= 79{
		grade = "A"
	}else if score >= 60 && score <= 69{
		grade = "A-"
	}else if score >= 50 && score <= 59{
		grade = "B"
	}else if score >= 40 && score <= 49 {
		grade = "C"
	}else if score >= 33 && score <= 39 {
		grade = "D"
	}else if score >= 0 && score <= 32 {
		grade = "F"
	}else {
		grade = "Score tidak boleh lebih dari 100 dan tidak boleh kurang dari 0"
	}

	if isMarried{
		fmt.Println("sudah menikah")
	}else{
		fmt.Println("masih single")
	}

	fmt.Println("Grade ", grade)




}

// if
// if else
// else if
// if bersarang