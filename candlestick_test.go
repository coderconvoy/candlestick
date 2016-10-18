package main

import (
	"fmt"
	"testing"
)

type testtable struct {
	question []int
	expP     int
	expC     int
}

func TestUniques(t *testing.T) {
	tab := []testtable{
		testtable{[]int{3, 6, 5, 6, 4, 5}, 4, 4},
		testtable{[]int{3, 6, 5, 7, 4, 5}, 0, 3},
		testtable{[]int{13, 6, 5, 1, 4, 5}, 3, 1},
		testtable{[]int{3, 13, 13, 6, 5, 1, 4, 5}, 3, 6},
		testtable{[]int{3, 4, 3, 4, 3, 4}, -1, -1},
	}

	for k, v := range tab {
		rP, rC := DecideTrick(v.question)
		if rP != v.expP {
			t.Fail()
			fmt.Printf("line %d:Test PLayer : %s , expected %d, got %d \n", k, v.question, v.expP, rP)

		}
		if rC != v.expC {
			t.Fail()
			fmt.Printf("line %d:Test Card : %s , expected %d, got %d \n", k, v.question, v.expC, rC)

		}
	}

}

type containsTest struct {
	qArr []int
	qCom int
	exp  int
}

func TestContains(t *testing.T) {
	ar1 := []int{4, 7, 3, 4, 6, 6, 4}
	tab := []containsTest{
		containsTest{ar1, 4, 3},
		containsTest{ar1, 7, 1},
		containsTest{ar1, 3, 1},
		containsTest{ar1, 6, 2},
		containsTest{ar1, 15, 0},
	}

	for _, v := range tab {
		res := contains(v.qArr, v.qCom)
		if res != v.exp {
			t.Fail()
			fmt.Printf("%s,%d: expected %d, got %d\n", v.qArr, v.qCom, v.exp, res)
		}
	}
}

func TestNewPlayers(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := NewPlayer()
		if len(res.deck) != 10 {
			t.Fail()
			fmt.Printf("deck not 10 long")
		}
		if len(res.hand) != 3 {
			t.Fail()
			fmt.Printf("hand not 3 long")
		}

		for _, v := range res.hand {
			if contains(res.deck, v) > 0 {
				t.Fail()
				fmt.Printf("overlap %d", v)
			}
		}

	}
}
