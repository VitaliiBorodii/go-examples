package main

import (
	"./chess"
	"./dove"
	"./forest"
	"./maze"
	Streaks "./divisible-streaks"
	"fmt"
	"math"
	"time"
	"sync"
	"strconv"
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

func runStreaksAsync(steps int) {
	start := time.Now()
	result := 0
	const num = 4
	var wg sync.WaitGroup
	var ch = make(chan int)

	for i := 1; i <= steps; i++ {
		wg.Add(1)
		go Streaks.PAsync(i, math.Pow(float64(num), float64(i)), ch)
	}

	go func() {
		for partialResult := range ch {
			result += partialResult
			wg.Done()
		}
	}()

	wg.Wait()
	fmt.Println("Async function P(i, " + strconv.Itoa(num) + "^i), where i = " + strconv.Itoa(steps) + " took :", time.Since(start), ", and result: ", result)

}

func runStreaksSync(steps int) {
	start := time.Now()
	result := 0
	const num = 4

	for i := 1; i <= steps; i++ {
		result += Streaks.P(i, math.Pow(float64(num), float64(i)))
	}

	fmt.Println("Sync function P(i, " + strconv.Itoa(num) + "^i), where i = " + strconv.Itoa(steps) + " took :", time.Since(start), ", and result: ", result)
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
	runStreaksSync(14)
	runStreaksAsync(14)
	fmt.Println(divider)
}
