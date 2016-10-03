package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	cards := make([]int, 13)

	rand.Seed(time.Now().Unix())
	for i := 0; i < 13; i++ {
		cards[i] = i + 1
	}

	for i := 0; i < 13; i++ {
		newloc := rand.Intn(13)
		swapper := cards[i]
		cards[i] = cards[newloc]
		cards[newloc] = swapper
	}

	fmt.Println(cards)
}
