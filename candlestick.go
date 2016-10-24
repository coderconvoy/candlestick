package main

import (
	"fmt"
	"math/rand"
	"mylib/asker"
	"net/http"
	"strconv"
	"time"
)

type Player struct {
	deck    []int
	hand    []int
	score   int
	isHuman bool
}

func NewPlayer(isH bool) *Player {
	res := &Player{
		makeDeck(),
		make([]int, 0),
		0,
		isH,
	}

	res.DrawCards(3)
	return res
}

func (self *Player) DrawCards(n int) bool {
	if len(self.deck) == 0 {
		return false
	}
	self.hand = append(self.hand, self.deck[:n]...)
	self.deck = self.deck[n:]
	return true
}

func (self *Player) ChooseCard() int {
	a := self.hand[0]
	if self.isHuman {
		a = AskCard(self.hand)
	}

	h2 := make([]int, 0)
	for _, v := range self.hand {
		if v != a {
			h2 = append(h2, v)
		}
	}
	self.hand = h2
	self.DrawCards(1)
	return a
}

func AskCard(cards []int) int {
	mess := "Please choose a card from "
	for _, v := range cards {
		mess += strconv.Itoa(v) + ", "
	}

	for {
		a := asker.AskInt(mess)
		if contains(cards, a) > 0 {
			return a
		}

		fmt.Println("Stick to the list please")

	}

}

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

func makeDeck() []int {
	cards := make([]int, 13)

	for i := 0; i < 13; i++ {
		cards[i] = i + 1
	}

	for i := 0; i < 13; i++ {
		newloc := rand.Intn(13)
		swapper := cards[i]
		cards[i] = cards[newloc]
		cards[newloc] = swapper
	}
	return cards

}

func round(pl []*Player, first int) int {
	pnum := first
	complete := false
	chosens := make([]int, len(pl))
	for !complete {
		chosens[pnum] = pl[pnum].ChooseCard()
		fmt.Printf("p%d : %d\n", pnum, chosens[pnum])
		pnum = (pnum + 1) % len(pl)
		if pnum == first {
			complete = true
		}
	}

	winP, winC := DecideTrick(chosens)
	pl[winP].score++
	fmt.Printf("Player %d wins with %d\n", winP, winC)

	scoresMess := "Scores : "
	for _, v := range pl {
		scoresMess += strconv.Itoa(v.score) + " "
	}
	fmt.Println(scoresMess)

	return winP

}

func playGame(nplayers int) int {
	if nplayers <= 0 {
		return -1
	}

	rand.Seed(time.Now().Unix())

	players := make([]*Player, nplayers)
	for i := 0; i < nplayers; i++ {
		players[i] = NewPlayer(i == 0)
		//		fmt.Printf("%s,\n%s\n,%d\n\n", players[i].deck, players[i].hand, players[i].score)
	}

	w := -1
	for i := 0; i < 13; i++ {
		if w == -1 {
			w = rand.Intn(nplayers)
		}
		w = round(players, w)
	}

	//Get winner.
	ws := -1
	wp := -1

	for k, p := range players {
		if p.score > ws {
			count := 0
			for i := 0; i < len(players); i++ {
				if p.score == players[i].score {
					count++
				}
			}
			if count == 1 {
				ws = p.score
				wp = k
			}
		}
	}

	if wp == -1 {
		fmt.Printf("Draw\n")
		return -1
	}

	fmt.Printf("Player %d wins\n\n", wp)
	return wp
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=\"/go\">Hello</a>, you pathed to :%s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handle)

	fmt.Printf("Server Starting\n")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
