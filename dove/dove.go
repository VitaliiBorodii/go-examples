package dove

type Dove struct {
	fed bool
}

func (d *Dove) eat(food *int) {

	if *food > 0 {
		*food--
		d.fed = true
	}
}

func FeedCount(food int) int {

	count := 0

	if food == 0 {
		return 0
	}

	doves := []Dove{}

	for food > 0 {
		count++

		for i := 0; i < count; i++ {
			doves = append(doves, Dove{})
		}

		for i, _ := range doves {
			doves[i].eat(&food)
		}

	}

	result := 0

	for _, dove := range doves {
		if dove.fed {
			result++
		}
	}

	return result
}
