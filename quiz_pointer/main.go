package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(namaGame string) {
	gamer.Games = append(gamer.Games, namaGame)
}

func main() {
	gamer := Gamer{Name: "Danil"}

	gamer.AddGame("Fifa 2020")
	gamer.AddGame("God Of War")
	gamer.AddGame("PES 2020")
	gamer.AddGame("Tekken 2021")
	gamer.AddGame("Need For Speed")

	fmt.Println("Gamer : ", gamer.Name)
	for _, game := range gamer.Games {
		fmt.Println("-> ", game, "| address = ", &game)
	}
}

// 1. buat sebuah method struct gamer dengan nama AddGame
