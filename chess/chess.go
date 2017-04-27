package chess

import (
	"math"
	"strconv"
	"fmt"
)

type Coord struct {
	X, Y int
}

type Move struct {
	Value Coord
	Prev  Coord
	Dist  float64
	Step int
	Visited bool
}

//Generate a string key from the Coords instance
func formKey(coord Coord) string {
	return strconv.Itoa(coord.X) + "-" + strconv.Itoa(coord.Y)
}


func CountSteps(start Coord, finish Coord, size int) int {
	TWO_MOVES_DIST := 2 * math.Sqrt(5)// 2^2 + 1^2 - one horse move
	visited := map[string]bool{}
	board := make([][]Move, size)

	for i, _ := range board {
		board[i] = make([]Move, size)
	}

	xf := finish.X
	yf := finish.Y
	steps := 0
	found := false

	createCoord := func(point Coord, prev Coord) Move {
		return Move{
			point,
			prev,
			math.Sqrt(math.Pow(float64(point.X - xf), 2) + math.Pow(float64(point.Y - yf), 2)),
			steps,
			true,
		}
	}

	visitNeighbours := func(points []Move, step int) (ret []Move) {
		ret = []Move{}
		neighbours := map[string]Move{}

		sortedInsert := func(coord Move) {
			if (coord.Dist > TWO_MOVES_DIST && len(ret) > 0) {
				val := ret[0]
				if (val.Dist > coord.Dist) {
					ret[0] = coord
				} else {
					//ret[0] = val
				}
			} else {
				ret = append(ret, coord);
			}
		}

		for _, coord := range points {
			visited[formKey(coord.Value)] = true

			array := []Coord{
				{coord.Value.X - 2, coord.Value.Y - 1},
				{coord.Value.X - 2, coord.Value.Y + 1},
				{coord.Value.X - 1, coord.Value.Y - 2},
				{coord.Value.X - 1, coord.Value.Y + 2},
				{coord.Value.X + 1, coord.Value.Y - 2},
				{coord.Value.X + 1, coord.Value.Y + 2},
				{coord.Value.X + 2, coord.Value.Y - 1},
				{coord.Value.X + 2, coord.Value.Y + 1},
			}

			for _, c := range array {
				neighbours[formKey(c)] = createCoord(c, coord.Value)
			}

		}

		nextStep := step + 1;
		for _, coord := range neighbours {

			if !(coord.Value.X < 0 || coord.Value.Y < 0 || coord.Value.X >= size || coord.Value.Y >= size) {

				cell := board[coord.Value.X][coord.Value.Y]

				if (!cell.Visited || nextStep < cell.Step) {
					board[coord.Value.X][coord.Value.Y] = Move{
						coord.Value,
						coord.Prev,
						coord.Dist,
						nextStep,
						true,
					}

					if (coord.Value.X == xf && coord.Value.Y == yf) {
						found = true
					}
					sortedInsert(coord)
				}

			}
		}
		return
	}

	next := []Move{{start, Coord{}, 0, 0, true,}}

	for !found {
		//fmt.Println(steps, next[0].Prev, "=>", next[0].Value)
		next = visitNeighbours(next, steps)
		steps++
	}

	return steps

}

func main() {
	boardSize := 10000
	start := Coord{0, 0}
	finish := Coord{9999, 9999}

	stepCount := CountSteps(start, finish, boardSize)

	fmt.Println("boardSize:", boardSize)
	fmt.Println("start:", start)
	fmt.Println("finish:", finish)
	fmt.Println("stepCount:", stepCount)

}