package day10

import (
	"bufio"
	"fmt"

	"github.com/imedgar/aoc24-imedgar/utils"
)

var (
	mountain   = [][]int{}
	trailheads = []Trailhead{}
	total      = 0
)

var directions = [][2]int{
	{-1, 0}, // Top
	{1, 0},  // Bottom
	{0, -1}, // Left
	{0, 1},  // Right
}

type Trailhead struct {
	x, y, score int
	visited     utils.Set[string]
	unvisited   [][]int
}

func Day10() {
	file := utils.ReadFile("./day10/input.txt")
	scanner := bufio.NewScanner(file)
	xAxis := 0

	for scanner.Scan() {
		line := scanner.Text()
		paths := utils.RunesToIntSlice([]rune(line))
		mountain = append(mountain, paths)
		trailheadIndexes := utils.FindIndexes(paths, 0)
		for _, h := range trailheadIndexes {
			trailheads = append(trailheads, Trailhead{
				x:         xAxis,
				y:         h,
				score:     0,
				visited:   utils.NewSet[string](),
				unvisited: [][]int{},
			})
		}
		xAxis++
	}

	totalScore := 0
	for _, th := range trailheads {
		trailhead := checkTrail(th)
		totalScore += trailhead.score
	}

	fmt.Printf("Total score of all trailheads: %d\n", totalScore)
}

func checkTrail(th Trailhead) Trailhead {
	visited := make(map[string]bool)
	queue := [][]int{{th.x, th.y}}
	rows, cols := len(mountain), len(mountain[0])
	score := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		x, y := current[0], current[1]
		key := fmt.Sprintf("%d-%d", x, y)
		if visited[key] {
			continue
		}
		visited[key] = true

		if mountain[x][y] == 9 {
			score++
			continue
		}

		nextLvl := mountain[x][y] + 1
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
				if mountain[newX][newY] == nextLvl {
					nKey := fmt.Sprintf("%d-%d", newX, newY)
					if !visited[nKey] {
						queue = append(queue, []int{newX, newY})
					}
				}
			}
		}
	}

	th.score = score
	return th
}

func checkSurroundings(th Trailhead, tar int) [][]int {
	rows := len(mountain)
	cols := len(mountain[0])
	next := [][]int{}
	for _, dir := range directions {
		newRow := th.x + dir[0]
		newCol := th.y + dir[1]
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			val := mountain[newRow][newCol]
			if val == tar {
				next = append(next, []int{newRow, newCol})
			}
		}
	}
	return next
}

func getKey(th Trailhead) string {
	return fmt.Sprintf("%d-%d", th.x, th.y)
}
