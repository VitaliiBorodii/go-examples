package maze

import "strconv"

type Coord struct {
	X, Y int
}

//Generate a string key from the Coords instance
func formKey(coord Coord) string {
	return strconv.Itoa(coord.X) + "-" + strconv.Itoa(coord.Y)
}

func CanPass(start Coord, finish Coord, maze [][]int) bool {
	visited := map[string]bool{}
	y := len(maze)
	x := len(maze[0])

	visitNeighbours := func(points []Coord) (ret []Coord) {
		ret = []Coord{}
		neighbours := []Coord{}

		/*
			Check all possible moves from the array of entry points
			and mark entry points as `visited`
		*/
		for _, coord := range points {
			visited[formKey(coord)] = true

			neighbours = append(neighbours, Coord{coord.X - 1, coord.Y})
			neighbours = append(neighbours, Coord{coord.X + 1, coord.Y})
			neighbours = append(neighbours, Coord{coord.X, coord.Y - 1})
			neighbours = append(neighbours, Coord{coord.X, coord.Y + 1})
		}

		/*
			Filter out points that get out from the maze range
			or were visited at previous iterations
		*/
		for _, coord := range neighbours {

			key := formKey(coord)
			_, ok := visited[key]

			if !(ok || coord.X < 0 || coord.Y < 0 || coord.X >= x || coord.Y >= y) {
				/*
					if we can go to this point then add it to the result array
					if not - mark it as unreachable
				*/
				if maze[coord.Y][coord.X] == 0 {
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

	_, ok := visited[formKey(finish)]

	return ok

}
