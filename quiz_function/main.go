package main

import (
	"errors"
	"fmt"
)


func main() {
	// soal 1
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println("Jumlah dari ", scores , " = ", total)

	// soal 2
	result, err := calculate(10,3,"=")
	if err != nil{
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

// menjumlahkan sebuah slice
func sum(scores[]int) int {
	var total int
	for _, score := range scores {
		total += score
	}

	return total
}

func calculate(number, numbertwo float64, operator string)(float64, error){
	var result float64
	var errorResult error
	// if operator == string("+") {
	// 	result = number + numbertwo
	// }else if operator == string("-") {
	// 	result = number - numbertwo
	// }else if operator == string("*"){
	// 	result = number * numbertwo
	// }else if operator == string("/"){
	// 	result = number / numbertwo
	// }else {
	// 	return 0 ,errors.New("operator salah")
	// }

	switch operator {
	case "+" :
		result = number + numbertwo
	case "-":
		result = number - numbertwo
	case "*":
		result = number * numbertwo
	case "/":
		result = number / numbertwo
	default :
		errorResult= errors.New("Unknown Operation")
	}
	return result, errorResult
}