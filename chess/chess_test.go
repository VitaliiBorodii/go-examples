package chess

import "testing"

type testpair struct {
	id     int
	size   int
	start  Coord
	finish Coord
	result int
}

var tests = []testpair{
	{
		1,
		10000,
		Coord{0, 0},
		Coord{9999, 9999},
		6666,
	},
	{
		2,
		1000,
		Coord{0, 0},
		Coord{999, 999},
		666,
	},
	{
		3,
		100,
		Coord{0, 0},
		Coord{99, 99},
		66,
	},
	{
		4,
		10000,
		Coord{0, 0},
		Coord{9, 9},
		6,
	},
	{
		5,
		1000,
		Coord{34, 44},
		Coord{456, 654},
		344,
	},
}

func TestCountSteps(t *testing.T) {
	for _, pair := range tests {
		v := CountSteps(pair.start, pair.finish, pair.size)
		if v != pair.result {
			t.Error(
				"For", pair.id,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
