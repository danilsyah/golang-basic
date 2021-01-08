package hitung

import "fmt"

// lakukan penjumlahan nilai array dan lakukan pembagian untuk dapatkan nilai rata-rata
func Avg() {
	scores := [8]int{100,80,75,92,70,93,88,67} //total 665

	var sum int
	for _, score := range scores{
		sum += score
	}

	var sumKonversi float32 = float32(sum)
	var length float32 = float32(len(scores))

	// hitungRata = 665 / 8
	resultAvg := sumKonversi / length

	fmt.Println("total nilai score : ", sum, "Rata-rata :", resultAvg)
}
