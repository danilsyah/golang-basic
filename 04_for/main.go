package main

import "fmt"

func main() {
	// cara pertama
	for i := 1; i <= 10; i++ {
		fmt.Println(i, "saya belajar golang")
	}

	// cara kedua
	i := 1
	for i <= 100{
		fmt.Println("Keep Spririt")
		i++
	}

	// foreach
	// gunankan _ jika variabel tidak akan di pakai
	title := "Golang the best language"
	for index, letter := range title {
		fmt.Println("index : ", index, "Letter : ", string(letter) , " = ", letter)
	}

	nama := "danil"
	for _, value := range nama{
		if value == 105 {
			continue
		} 
		fmt.Println("ejaan : ", string(value), "=", value)
	}

	// quiz 
	// 1. pengulangan dari sebuah string, jika index nya genap maka di cetak,
	// 2. cetak hanya huruf vokal , a,i,u,e,o
	quiz := "GolAng the bEst langUage"

	// mencetak karakter index genap
	for index, letter := range quiz{
		if index % 2 == 0 {
			fmt.Println("index : ", index,"letter : ", string(letter))
		}
	}

	// mencetak hanya karakter vocal
	for _, letter := range quiz {
		letterString := string(letter)
		if letterString == "a" || letterString == "A" || letterString == "i" || letterString == "I" || letterString == "u" || letterString == "U" || letterString == "e" || letterString == "E" || letterString == "o" || letterString == "O"{
			fmt.Println(letterString)
		}
	}

	for index, letter := range quiz {
		letterString := string(letter)

		switch letterString {
		case "a","A","i","I","u","U","e","E","o","O":
			fmt.Println("index : ", index, " letter : ", letterString)
		}
	}
}