package chess

import (
	"fmt"
	"math"
	"strconv"
)

type Coord struct {
	X, Y int
}

type CoordLink struct {
	Value Coord
	Prev  Coord
}

type Move struct {
	Value   Coord
	Prev    Coord
	Dist    float64
	Step    int
	Visited bool
}

//Generate a string key from the Coords instance
func formKey(coord Coord) string {
	return strconv.Itoa(coord.X) + "-" + strconv.Itoa(coord.Y)
}

func printPath(start Coord, finish Coord, board [][]Move) {
	point := board[finish.Y][finish.X]
	end := false

	for end == false {
		prevCoord := point.Prev

		fmt.Println(point.Step, ":", prevCoord, "=>", point.Value)

		if prevCoord.X == start.X && prevCoord.Y == start.Y {
			end = true
		} else {
			point = board[prevCoord.Y][prevCoord.X]
		}
	}
}

func CountSteps(start Coord, finish Coord, size int) int {
	TWO_MOVES_DIST := 2 * math.Sqrt(5) // 2^2 + 1^2 - one horse move
	visited := map[string]bool{}
	board := make([][]Move, size)
	steps := 0
	found := false
	xf := finish.X
	yf := finish.Y

	for i, _ := range board {
		board[i] = make([]Move, size)
	}

	board[start.Y][start.X] = Move{
		start,
		start,
		math.Sqrt(math.Pow(float64(start.X - xf), 2) + math.Pow(float64(start.Y - yf), 2)),
		0,
		true,
	}

	visitNeighbours := func(points []Coord, step int) (result []Coord) {
		ret := []Move{}
		neighbours := map[string]CoordLink{}

		sortedInsert := func(coord Move) {
			if coord.Dist > TWO_MOVES_DIST && len(ret) > 0 {
				val := ret[0]
				if val.Dist > coord.Dist {
					ret[0] = coord
				}
			} else {
				ret = append(ret, coord)
			}
		}

		for _, coord := range points {
			visited[formKey(coord)] = true

			array := []Coord{
				{coord.X - 2, coord.Y - 1},
				{coord.X - 2, coord.Y + 1},
				{coord.X - 1, coord.Y - 2},
				{coord.X - 1, coord.Y + 2},
				{coord.X + 1, coord.Y - 2},
				{coord.X + 1, coord.Y + 2},
				{coord.X + 2, coord.Y - 1},
				{coord.X + 2, coord.Y + 1},
			}

			for _, c := range array {
				neighbours[formKey(c)] = CoordLink{c, coord}
			}

		}

		nextStep := step + 1
		for _, coord := range neighbours {

			if !(coord.Value.X < 0 || coord.Value.Y < 0 || coord.Value.X >= size || coord.Value.Y >= size) {

				cell := board[coord.Value.Y][coord.Value.X]

				if !cell.Visited || nextStep < cell.Step {
					newCell := Move{
						coord.Value,
						coord.Prev,
						math.Sqrt(math.Pow(float64(coord.Value.X - xf), 2) + math.Pow(float64(coord.Value.Y - yf), 2)), nextStep,
						true,
					}

					board[coord.Value.Y][coord.Value.X] = newCell
					sortedInsert(newCell)

					if coord.Value.X == xf && coord.Value.Y == yf {
						found = true
					}
				}

			}
		}
		for _, move := range ret {
			result = append(result, move.Value)
		}
		return
	}

	next := []Coord{start}

	for !found {
		next = visitNeighbours(next, steps)
		steps++
	}

	//printPath(start, finish, board)

	return steps

}
