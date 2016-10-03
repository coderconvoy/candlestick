package main

import (
	"fmt"
	"testing"
)

type testtable struct {
	question []int
	expected int
}

func TestUniques(t *testing.T) {
	tab := []testtable{
		testtable{[]int{3, 6, 5, 6, 4, 5}, 4},
		testtable{[]int{3, 6, 5, 7, 4, 5}, 3},
		testtable{[]int{13, 6, 5, 1, 4, 5}, 1},
		testtable{[]int{3, 13, 13, 6, 5, 1, 4, 5}, 6},
	}

	for i := 0; i < len(tab); i++ {
		r := DecideTrick(tab[i].question)
		if r != tab[i].expected {
			t.Fail()
			fmt.Printf("Test : %s , expected %d, got %d \n", tab[i].question, tab[i].expected, r)

		}
	}

}
