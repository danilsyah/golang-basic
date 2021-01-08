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
