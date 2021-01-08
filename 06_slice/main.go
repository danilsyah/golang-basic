package main

import "fmt"

func main() {
	// cara penulisan slice , mirip dengan array
	var gamingConsoles []string

	gamingConsoles = append(gamingConsoles, "Playstation5")
	gamingConsoles = append(gamingConsoles, "Playstation4")
	gamingConsoles = append(gamingConsoles, "Nintendo Switch")
	gamingConsoles = append(gamingConsoles, "Xbox One")

	fmt.Println(gamingConsoles)


	// penulisan cara kedua
	persons := []string {
		"danil",
		"haykal",
		"budi",
		"emen",
	}
	
	for _, value := range persons {
		fmt.Println(value)
	}

}