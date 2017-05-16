package divisible_streaks

func streak(n int) int {
	var k int;

	for i := 1; i < n; i++ {
		k = i;
		var ifStatement = (n + k) % (i + 1)
		if (ifStatement != 0) {
			break;
		}
	}

	return k;
}

func P(s int, N float64) int {
	var acc int = 0
	var i float64;

	for i = 1; i < N; i++ {
		if (streak(int(i)) == s) {
			acc++
		}
	}

	return acc
}