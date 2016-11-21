package types

import (
	"fmt"
	"github.com/coderconvoy/candlestick/pure"
	"math/rand"
)

const TURN_BOT = 0
const TURN_HUMAN = 1
const TURN_ROUND_DONE = 2

type Player struct {
	Deck    []int
	Hand    []int
	Score   int
	IsHuman bool
}

type GameRound struct {
	Played []int
	First  int
	Curr   int
	Winner int
}

func NewGameRound(nplayers, first int) *GameRound {
	return &GameRound{make([]int, nplayers), first, first, -1}
}

type Game struct {
	Players []*Player
	Rounds  []*GameRound
	Message string
}

func NewGame(nplayers int) *Game {

	players := make([]*Player, nplayers)
	for i := 0; i < nplayers; i++ {
		players[i] = NewPlayer(i == 0)
	}
	return &Game{players, []*GameRound{NewGameRound(nplayers, 0)}, ""}
}

//HumanTurn takes a card and moves the game on by playing that card
//params: c card name
//returns bool true if OK
func (g *Game) HumanTurn(c int) bool {

	cr := g.Rounds[len(g.Rounds)-1]

	//Check card in deck
	pos := -1
	for k, v := range g.Players[cr.Curr].Hand {
		if v == c {
			pos = k
		}
	}
	if pos == -1 {
		return false
	}

	cr.Played[cr.Curr] = g.Players[cr.Curr].PlayCardPos(pos)
	cr.Curr = (cr.Curr + 1) % len(g.Players)

	if cr.Curr == cr.First {
		cr.Winner, _ = pure.DecideTrick(cr.Played)
		g.Rounds = append(g.Rounds, NewGameRound(len(g.Players), cr.Winner))
	}

	return true

}

func (g *Game) HelloWorld(s string) string {
	return fmt.Sprintf("hello world1111, %s\n", s)
}

func (g *Game) TryTurn() (int, int) {
	cr := g.Rounds[len(g.Rounds)-1]
	if g.Players[cr.Curr].IsHuman {
		return TURN_HUMAN, cr.Curr
	}
	cr.Played[cr.Curr] = g.Players[cr.Curr].PlayCardPos(0)
	cr.Curr = (cr.Curr + 1) % len(g.Players)

	if cr.Curr == cr.First {
		cr.Winner, _ = pure.DecideTrick(cr.Played)
		g.Rounds = append(g.Rounds, NewGameRound(len(g.Players), cr.Winner))
		return TURN_ROUND_DONE, 0
	}

	//g.Players[cr.Curr]
	return TURN_BOT, 0
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

func (p *Player) PlayCardPos(n int) int {
	res := p.Hand[n]

	nHand := make([]int, 0)
	for i := 0; i < len(p.Hand); i++ {
		if i != n {
			nHand = append(nHand, p.Hand[i])
		}
	}
	p.Hand = nHand
	p.DrawCards(1)
	return res

}

func (p *Player) DrawCards(n int) bool {
	if len(p.Deck) == 0 {
		return false
	}
	p.Hand = append(p.Hand, p.Deck[:n]...)
	p.Deck = p.Deck[n:]
	return true
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
