package main

import (
	"./maze"
	"./chess"
	"fmt"
)

func runMaze() {
	board := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1, 1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
	}
	start := maze.Coord{0, 0}
	finish := maze.Coord{9, 5}
	isCanPass := maze.CanPass(start, finish, board)

	fmt.Println("Maze:")

	for i, row := range board {
		//This is done only to visualize the maze
		row1 := make([]int, len(row))
		copy(row1, row)
		if i == start.Y {
			row1[start.X] = 8
		}
		if i == finish.Y {
			row1[finish.X] = 8
		}
		fmt.Println(row1)

	}

	fmt.Println("start:", start)
	fmt.Println("finish:", finish)
	fmt.Println("Can find a way:", isCanPass)
}

func runChess() {
	boardSize := 100
	start := chess.Coord{3, 1}
	finish := chess.Coord{98, 72}

	steps := chess.CountSteps(start, finish, boardSize)

	fmt.Println("boardSize:", boardSize)
	fmt.Println("start:", start)
	fmt.Println("finish:", finish)
	fmt.Println("steps:", steps)
}

func main() {
	runMaze()
	fmt.Println("====================================================")
	runChess()
}
