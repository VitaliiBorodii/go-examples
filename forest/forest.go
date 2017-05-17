package forest

import "math"

func CountWays(trees int, left int) int {

	if trees == left || trees < 0 || left < 0 {
		return 0
	}

	if trees == left {
		return 1
	}

	result := 0

	diff := int(math.Ceil(float64(trees) / float64(left)))

	for i := 0; i < diff; i++ {
		acc := trees - (left-1)*(i+1)
		if acc < 0 {
			break
		}
		result += acc
	}

	return result
}
