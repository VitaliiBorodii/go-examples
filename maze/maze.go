package maze

import (
	"fmt"
	"strconv"
)

type Coord struct {
	x, y int
}

//Generate a string key from the Coords instance
func formKey(c Coord) string {
	return strconv.Itoa(c.x) + "-" + strconv.Itoa(c.y)
}

func CanPass(start Coord, finish Coord, maze [][]int) bool {
	visited := map[string]bool{}
	y := len(maze)
	x := len(maze[0])

	visitNeighbours := func(coords []Coord) (ret []Coord) {
		ret = []Coord{}
		neighbours := []Coord{}

		/*
			Check all possible moves from the array of entry points
			and mark entry points as `visited`
		*/
		for _, c := range coords {
			visited[formKey(c)] = true

			neighbours = append(neighbours, Coord{c.x - 1, c.y})
			neighbours = append(neighbours, Coord{c.x + 1, c.y})
			neighbours = append(neighbours, Coord{c.x, c.y - 1})
			neighbours = append(neighbours, Coord{c.x, c.y + 1})
		}

		/*
			Filter out points that get out from the maze range
			or were visited at previous iterations
		*/
		for _, coord := range neighbours {

			key := formKey(coord)
			_, ok := visited[key]

			if !(ok || coord.x < 0 || coord.y < 0 || coord.x >= x || coord.y >= y) {
				/*
					if we can go this point the add it to the result array
					if not - mark it as unreachable
				*/
				if maze[coord.y][coord.x] == 0 {
					ret = append(ret, coord)
				} else {
					visited[key] = false
				}
			}
		}
		return
	}

	next := []Coord{start}

	for len(next) != 0 {
		next = visitNeighbours(next)
	}

	key, ok := visited[formKey(finish)]

	if ok {
		return key
	} else {
		return false
	}

}

func main() {

	maze := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1, 1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
	}
	start := Coord{0, 0}
	finish := Coord{9, 5}
	isCanPass := CanPass(start, finish, maze)

	fmt.Println("Maze:")

	for _, row := range maze {
		fmt.Println(row)

	}

	fmt.Println("start:", start)
	fmt.Println("finish:", finish)
	fmt.Println("Can find a way:", isCanPass)

}
