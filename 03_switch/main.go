package main

import "fmt"

func main() {
	menu := 5

	switch menu {
	case 1:
		fmt.Println("Program Luas Segitiga")
	case 2:
		fmt.Println("Program Luas Kotak")
	case 3:
		fmt.Println("Program Kalkulator")
	default :
		fmt.Println("Pilihan tidak ada")
	}

	// switch short statement
	name := "danil"
	switch length := len(name); length > 5 {
	case true :
		fmt.Println("nama terlalu panjang", len(name))
	case false :
		fmt.Println("nama sudah benar", len(name))
	}

	// buat program validasi panjang password
	// password minimal 8 karakter
	password := "12345678"
	switch length := len(password); length >= 8 {
	case true:
		fmt.Println("password success")
	case false:
		fmt.Println("password minimal 8 karakter")
	}

}