package dove

import "testing"

type testpair struct {
	id     int
	food   int
	result int
}

var tests = []testpair{
	{
		1,
		1,
		1,
	},
	{
		2,
		3,
		2,
	},
	{
		3,
		10,
		6,
	},
	{
		4,
		18,
		8,
	},
	{
		5,
		200,
		45,
	},
}

func TestFeedCount(t *testing.T) {
	for _, pair := range tests {
		v := FeedCount(pair.food)
		if v != pair.result {
			t.Error(
				"For", pair.id,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
