package pure

//contains returns the number of elements in the slice that match the comparator (c)
func contains(ar []int, c int) int {
	res := 0
	for i := 0; i < len(ar); i++ {
		if c == ar[i] {
			res++
		}
	}
	return res
}

//beats returns true if c1 beats c2
func beats(c1, c2 int, aceHigh, reverse bool) bool {
	if aceHigh {
		if c1 == 1 {
			c1 = 14
		}
		if c2 == 1 {
			c2 = 14
		}
	}
	if reverse {
		return c2 > c1
	}
	return c1 > c2
}

//DecideTrick Takes a list of cards, one for each player, and decides who won according to CandleStick Rules

//returns winningPlayer, winningCard
//returns -1,-1 for a draw
func DecideTrick(ar []int) (int, int) {

	uniques := make([]int, 0)

	for i := 0; i < len(ar); i++ {
		if contains(ar, ar[i]) == 1 {
			uniques = append(uniques, ar[i])
		}
	}

	if len(uniques) == 0 {
		return -1, -1
	}

	aceHigh := contains(uniques, 13) > 0

	reverse := false
	if contains(uniques, 7) > 0 {
		reverse = !reverse
	}

	if contains(uniques, 8) > 0 {
		reverse = !reverse
	}

	winningCard := uniques[0]
	for i := 1; i < len(uniques); i++ {
		if beats(uniques[i], winningCard, aceHigh, reverse) {
			winningCard = uniques[i]
		}
	}

	winningPlayer := -1
	for k, v := range ar {
		if v == winningCard {
			winningPlayer = k
		}
	}
	return winningPlayer, winningCard

}
