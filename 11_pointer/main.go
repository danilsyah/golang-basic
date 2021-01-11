package main

import "fmt"

func main() {
	// pointer adalah sebuah variabel yang menyimpan alamat memory di sebuah nilai
	numberA := 5
	// gunakan simbol & untuk referencing alamat memory variabel (referen)
	numberB := &numberA // numberB *int = &numberA

	fmt.Println(numberA)
	fmt.Println(numberB)
	// gunakan simbol * untuk menampilkan nilai pada alamat memory (direferen)
	fmt.Println(*numberB)
	// ketika nilai variable numberB di ubah, maka variabel nilaiA pun ikut terubah,
	// karena yang di ubah alamat memory bukan nilai variabel nya
	*numberB = 15
	fmt.Println(*numberB)
	fmt.Println(numberA)

	var nameA string = "danil"
	var nameB *string = &nameA
	fmt.Println(nameA)
	fmt.Println(nameB)
	fmt.Println(*nameB)

	nameA = "udin"
	fmt.Println(nameA)
	fmt.Println(nameB)
	fmt.Println(*nameB)
	fmt.Println("=======================================")
	// contoh pointer mengubah nilai
	angka := 5
	fmt.Println("alamat memory : ", &angka)
	fmt.Println("nilai awal : ", angka)

	change(&angka, 100)

	fmt.Println("Nilai akhir : ", angka)
	fmt.Println("Alamat memory : ", &angka)

}

// contoh pointer
func change(old *int, new int) {
	// fmt.Println("alamat memory di dalam function : ", &old)
	// fmt.Println("nilai : ", *old)
	*old = new
	fmt.Println("Alamat memory dlm function : ", old)
	fmt.Println("DI dalam function : ", *old)

}
