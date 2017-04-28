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
	Value Coord
	Prev  Coord
	Dist  float64
	Step  int
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

func CountSteps(start Coord, finish Coord, size int, log bool) int {
	TWO_MOVES_DIST := 2 * math.Sqrt(5) // 2^2 + 1^2 - one horse move
	board := make([][]Move, size)
	steps := 0
	found := false

	for i, _ := range board {
		board[i] = make([]Move, size)
	}

	board[start.Y][start.X] = Move{
		start,
		start,
		math.Sqrt(math.Pow(float64(start.X - finish.X), 2) + math.Pow(float64(start.Y - finish.Y), 2)),
		0,
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

				//Check if we weren't on this board place before or we can find more optimized way to it
				if (cell == Move{} || nextStep < cell.Step) {
					newCell := Move{
						coord.Value,
						coord.Prev,
						math.Sqrt(math.Pow(float64(coord.Value.X - finish.X), 2) + math.Pow(float64(coord.Value.Y - finish.Y), 2)), nextStep,
					}

					board[coord.Value.Y][coord.Value.X] = newCell
					sortedInsert(newCell)

					if newCell.Dist == 0 {
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

	if log {
		printPath(start, finish, board)
	}

	return steps

}
