package day6

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/imedgar/aoc24-imedgar/utils"
)

var (
	guardMov = []string{"^", ">", "v", "<"}
	startX   = -1
	startY   = -1
	lab      = [][]rune{}
)

func Day6() {
	file := utils.ReadFile("./day6/input.txt")
	fl := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		located := locateGuard(line)
		if located != "" {
			startX, startY = fl, strings.Index(line, located)
		}
		lab = append(lab, []rune(line))
		fl++
	}
	fmt.Println(startX, startY)
	guard := NewGuard(lab[startX][startY], startX, startY)

	guard.patrol()
	fmt.Println("finished patrolling", len(guard.visited))
	anomalies := 0
	visited := guard.visited
	for _, pos := range visited {
		x, y := parsePosition(pos)
		savePos := lab[startX][startY]
		guard2 := NewGuard(savePos, startX, startY)
		lab[x][y] = 'O'
		if guard2.checkAnomalies() {
			anomalies++
		}
		lab[x][y] = '.'
		lab[startX][startY] = savePos
	}
	fmt.Println("anomalies:", anomalies)
}

func locateGuard(floor string) string {
	for _, t := range guardMov {
		if strings.Contains(floor, t) {
			return t
		}
	}
	return ""
}

func debug(milli time.Duration) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("\nLab:")
	for _, row := range lab {
		fmt.Println(string(row))
	}
	time.Sleep(milli * time.Millisecond)
}

type Guard struct {
	direction rune
	position  []int
	visited   []string
	stuck     utils.Set[string]
}

func NewGuard(dir rune, x, y int) *Guard {
	return &Guard{direction: dir, position: []int{x, y}, stuck: utils.NewSet[string]()}
}

func (g *Guard) patrol() {
	patrolling := 1
	for patrolling == 1 {
		g.memorize()
		patrolling = g.move()
	}
}

func (g *Guard) checkAnomalies() bool {
	patrolling := 1
	for patrolling == 1 {
		if g.checkStuck() {
			return true
		}
		patrolling = g.move()
	}
	return false
}

func (g *Guard) rotate() {
	switch g.direction {
	case '^':
		g.direction = '>'
	case '>':
		g.direction = 'v'
	case 'v':
		g.direction = '<'
	case '<':
		g.direction = '^'
	}
}

func (g *Guard) move() int {
	switch g.direction {
	case '^': // Moving up
		if g.position[0] == 0 { // Edge of the lab
			return 0
		}
		facing := lab[g.position[0]-1][g.position[1]]
		if facing == '#' || facing == 'O' { // Wall or obstacle
			g.rotate()
		} else {
			g.position[0]--
		}

	case '>': // Moving right
		if g.position[1] == len(lab[0])-1 { // Edge of the lab
			return 0
		}
		facing := lab[g.position[0]][g.position[1]+1]
		if facing == '#' || facing == 'O' { // Wall or obstacle
			g.rotate()
		} else {
			g.position[1]++
		}

	case 'v': // Moving down
		if g.position[0] == len(lab)-1 { // Edge of the lab
			return 0
		}
		facing := lab[g.position[0]+1][g.position[1]]
		if facing == '#' || facing == 'O' { // Wall or obstacle
			g.rotate()
		} else {
			g.position[0]++
		}

	case '<': // Moving left
		if g.position[1] == 0 { // Edge of the lab
			return 0
		}
		facing := lab[g.position[0]][g.position[1]-1]
		if facing == '#' || facing == 'O' { // Wall or obstacle
			g.rotate()
		} else {
			g.position[1]--
		}
	}
	return 1 // Guard has not escaped
}

func (g *Guard) memorize() {
	location := strconv.Itoa(g.position[0]) + "-" + strconv.Itoa(g.position[1])
	if slices.Index(g.visited, location) == -1 {
		g.visited = append(g.visited, location)
	}
}

func (g *Guard) checkStuck() bool {
	location := strconv.Itoa(g.position[0]) + "-" + strconv.Itoa(g.position[1]) + string(g.direction)
	if !g.stuck.Contains(location) {
		g.stuck.Add(location)
		return false
	}
	return true
}

func parsePosition(pos string) (int, int) {
	var x, y int
	fmt.Sscanf(pos, "%d-%d", &x, &y)
	return x, y
}
