package main

import (
	"./chess"
	"./dove"
	"./forest"
	"./maze"
	Streaks "./divisible-streaks"
	"fmt"
	"math"
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
	boardSize := 10000
	start := chess.Coord{3, 1}
	finish := chess.Coord{908, 72}
	logSteps := true

	steps := chess.CountSteps(start, finish, boardSize, logSteps)

	fmt.Println("boardSize:", boardSize)
	fmt.Println("start:", start)
	fmt.Println("finish:", finish)
	fmt.Println("steps:", steps)
}

func runDove() {
	food := 55
	result := dove.FeedCount(food)
	fmt.Println("Can feed '", result, "' doves with `", food, "` ammount of food")

}

func runFell() {
	initialTrees := 26
	treesToLeft := 12
	result := forest.CountWays(initialTrees, treesToLeft)

	fmt.Println("Have trees:", initialTrees)
	fmt.Println(result, "ways, to left", treesToLeft, " trees with same distance between them")

}

func runStreaks() {
	var result int = 0
	for i := 1; i < 14; i++ {
		fmt.Println("calculating i =  ", i);
		localResult := Streaks.P(i, math.Pow(4, float64(i)))
		fmt.Println("local result:", localResult)

		result += localResult;

	}
	fmt.Println("Streaks.P(3, 13)", Streaks.P(3, 13));
	fmt.Println("Streaks.P(6, 10^6)", Streaks.P(6, math.Pow(10, 6)))
	fmt.Println("Result: ", result);
}

func main() {
	divider := "<================================================================================>"
	runMaze()
	fmt.Println(divider)
	runChess()
	fmt.Println(divider)
	runDove()
	fmt.Println(divider)
	runFell()
	fmt.Println(divider)
	runStreaks()
	fmt.Println(divider)
}
