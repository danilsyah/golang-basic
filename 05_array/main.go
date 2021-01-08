package main

import "fmt"

func main() {
	// penulisan cara pertama dengan nilai element di tentukan
	var languages [5]string
	languages[0] = "Go"
	languages[1] = "Python"
	languages[2] = "Ruby"
	languages[3] = "Kotlin"
	languages[4] = "PHP"

	for index, lang := range languages {
		fmt.Println("ke-",index," : ", lang)
	}

	fmt.Println("jumlah element :", len(languages))

	// penulisan cara kedua
	names := [5]string{"danil","haykal","icih","udin","emen"}
	fmt.Println(names)


	// penulisan array dengan nilai elemen dinamis
	matakuliah := [...]string{
		"b.indonesia",
		"inggris",
		"matematika",
		"pemrograman",
		"analisis basis data",
	}

	fmt.Println(matakuliah , len(matakuliah))
}