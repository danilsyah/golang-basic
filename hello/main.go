package main

import (
	"fmt"
	"hello/calculation"
)

func main() {
	fmt.Println("Hello belajar golang")
	// sentence := TestAja()
	// fmt.Println(sentence)

	// memanggil function dari package lain
	resultAdd := calculation.Add(8,3)
	resultSubstraction := calculation.Sub(10,3)
	fmt.Println("jumlah",resultAdd)
	fmt.Println("kurang : ",resultSubstraction)
}