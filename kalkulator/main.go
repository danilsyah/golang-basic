package main

import (
	"fmt"
	"kalkulator/operator"
)

func main(){
	fmt.Println("--Program Kalkulator--")

	// memanggil function dari package operator
	result := operator.Multi(3,4)

	fmt.Println("hasil dari perkalian : ",result)

}