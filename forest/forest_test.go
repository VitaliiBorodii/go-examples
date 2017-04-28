package forest

import "testing"

type testpair struct {
	id     int
	trees  int
	left   int
	result int
}

var tests = []testpair{
	{
		1,
		5,
		3,
		4,
	},
	{
		2,
		6,
		3,
		6,
	},
	{
		3,
		10,
		8,
		3,
	},
	{
		4,
		10,
		6,
		5,
	},
	{
		5,
		10,
		4,
		12,
	},
}

func TestCountWays(t *testing.T) {
	for _, pair := range tests {
		v := CountWays(pair.trees, pair.left)
		if v != pair.result {
			t.Error(
				"For", pair.id,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
