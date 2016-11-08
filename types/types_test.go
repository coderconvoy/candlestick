package types

import (
	"testing"
)

func Test_playcardpos(t *testing.T) {
	p := NewPlayer(true)
	h := p.Hand
	for i := 0; i < 3; i++ {
		c := p.PlayCardPos(0)
		if len(p.Hand) != 3 {
			t.Logf("Hand len = %d", len(p.Hand))
			t.Fail()
		}
		if h[i] != c {
			t.Logf("i = %d, c = %d, h[i] = %d\n", i, c, h[i])
			t.Fail()
		}
	}

}
