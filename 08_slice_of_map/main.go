package main

import "fmt"

func main() {
	// mirip array assosiatif
	students := []map[string]string{
		{"name": "Danil", "score": "A"},
		{"name": "Haykal", "score": "A+"},
		{"name": "Fika", "score": "B"},
	}

	fmt.Println(students[0]["score"]) //A

	for _, student := range students {
		fmt.Println(student["name"], "-", student["score"])
	}
}