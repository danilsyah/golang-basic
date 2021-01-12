package main

import "fmt"

type BangunDatar interface {
	HitungLuas() float32
}

type Persegi struct {
	Sisi float32
}

func (persegi Persegi) HitungLuas() float32 {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang float32
	Lebar   float32
}

func (persegiPanjang PersegiPanjang) HitungLuas() float32 {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

type SegitigaSikuSiku struct {
	Alas   float32
	Tinggi float32
}

func (segitiga SegitigaSikuSiku) HitungLuas() float32 {
	return 0.5 * segitiga.Alas * segitiga.Tinggi
}

func main() {
	persegiPanjang := PersegiPanjang{Panjang: 6, Lebar: 5}
	luas := SeberapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang : ", luas)

	persegi := Persegi{Sisi: 4}
	luas = SeberapaLuas(persegi)
	fmt.Println("Luas Persegi : ", luas)

	segitiga := SegitigaSikuSiku{Alas: 10, Tinggi: 8}
	luas = SeberapaLuas(segitiga)
	fmt.Println("Luas Segitiga :", luas)
}

func SeberapaLuas(bangunDatar BangunDatar) float32 {
	return bangunDatar.HitungLuas()
}
