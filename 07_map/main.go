package main

import (
	"fmt"
	"time"
)


func main() {

	// penulisan cara pertama
	var myMap map[string]string
	myMap = map[string]string{}

	myMap["nama"] = "danil"
	myMap["umur"] = "26"
	myMap["jenis"] = "pria"

	// fmt.Println(myMap)
	// fmt.Println(myMap["nama"])

	// mengakses map
	for key, value := range myMap{
		fmt.Println("key : ", key, ", value : ", value)
	}

	// penulisan cara kedua
	scores := map[string]int{
		"inggris" : 77,
		"pemrograman" : 80,
		"indonesia" : 90,
		"matematika" : 78,
	}

	// fmt.Println(scores)

	for key, value := range scores {
		fmt.Println("key : ", key, ", value : ", value)
	}

	// menghapus value map
	fmt.Println("======================")

	delete(scores,"indonesia")

	for key, value := range scores {
		fmt.Println("key : ", key, ", value : ", value)
	}

	// cek ada tidak nya value di map
	value, isAvailable := scores["sunda"]
	fmt.Println(value) //0
	fmt.Println(isAvailable) //false

	fmt.Println("Waktu sekarang adalah : ",time.Now())

}