package divisible_streaks

import (
	"testing"
)

type testpair struct {
	id     int
	s      int
	N      float64
	result int
}

var tests = []testpair{
	{
		1,
		3,
		13,
		1,
	},
	{
		2,
		6,
		1000000,
		14286,
	},
}

func TestP(t *testing.T) {

	for _, pair := range tests {
		v := P(pair.s, pair.N)
		if v != pair.result {
			t.Error(
				"For", pair.id,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func TestPAsync(t *testing.T) {
	for _, pair := range tests {
		var c = make(chan int)
		var v int

		go PAsync(pair.s, pair.N, c)

		v = <-c

		if v != pair.result {
			t.Error(
				"For", pair.id,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
