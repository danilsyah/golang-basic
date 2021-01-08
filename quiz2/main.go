package main

import (
	"fmt"
	"quiz2/hitung"
)

func main() {
	hitung.Avg()

	scores := [8]int{100,80,75,92,70,93,88,67}
	var goodScores []int

	for _,score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}

	fmt.Println(goodScores)
}