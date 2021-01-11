package main

import (
	"fmt"
	"strconv"
)

func main() {
	printMyResult()

	printMyResult2("danil")
	printMyResult2("haykal")
	printMyResult2("razil")

	hasilTambah := tambah(5, 6)
	fmt.Println(hasilTambah)

	fmt.Println(fullName("danil", "syah", 26))

	// gunakan _ ketika tidak membutuhkan nilai kembalian luas atau keliling
	luas, keliling := calculate(2, 5)
	fmt.Println("luas : ",luas)
	fmt.Println("keliling : ",keliling)

	fullName, umur := biodata("danil","syah", 26)
	fmt.Println(fullName)
	fmt.Println(umur)
}

// function tanpa parameter
func printMyResult() {
	fmt.Println("hello danil, kamu sedang belajar go")
}

// function dengan parameter
func printMyResult2(name string) {
	fmt.Println(name, "akan menjadi programmer")
}

// function dengan nilai kembalian
func tambah(number int, numbertwo int) int {
	return number + numbertwo
}

func fullName(firstName, lastName string, old int) string {
	fullName := firstName + " " + lastName + " my Old : " + strconv.Itoa(old)
	return fullName
}

// function multiple return
func calculate(panjang int, lebar int)(int, int){
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

// penggunaan nama nilai kembalian
func biodata(firstName,lastName string, umur int)(fullName string,old int){
	fullName = firstName + lastName
	old = umur

	return
}
