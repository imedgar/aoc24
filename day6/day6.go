package day6

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/imedgar/aoc24-imedgar/utils"
)

var (
	lab          = [][]rune{}
	labInit      = [][]rune{}
	route        = [][]int{}
	guardMov     = []string{"^", ">", "v", "<"}
	guardPos     = &[]int{}
	guardPosInit = []int{}
	isLooped     = []string{}
)

func Day6() {
	file := utils.ReadFile("./day6/input.txt")

	fl := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		located := locateGuard(line)
		if located != "" {
			*guardPos = []int{fl, strings.Index(line, located)}
			guardPosInit = utils.CopySlice(*guardPos) // Use CopySlice here
		}
		lab = append(lab, []rune(line))
		fl++
	}
	labInit = utils.DeepCopy(lab) // Use DeepCopy for a proper deep copy
	for move(guardPos) != 0 {
		//debug(10)
	}
	steps := 1
	for _, f := range lab {
		for _, t := range f {
			if t == 'X' {
				steps++
			}
		}
	}
	fmt.Println("finished patrolling", steps)
	fmt.Println("route len", len(route))
	anomalies := 0
	route = utils.RemoveElementFrom(route, 0)
	//test := []int{40, 51}
	//route = [][]int{test}
	for _, xy := range route {
		x, y := xy[0], xy[1]
		lab = utils.DeepCopy(labInit)             // Reset lab to its initial state
		*guardPos = utils.CopySlice(guardPosInit) // Reset guard position properly
		isLooped = []string{}
		fmt.Println(isLooped)
		if lab[x][y] != '.' {
			continue
		}

		fmt.Println("checking anomaly", x, y)
		lab[x][y] = 'O'
		stuck := move(guardPos)
		for stuck == 1 {
			if x == 39 && y == 51 {
				//debug(10)
			}
			stuck = move(guardPos)
		}
		if stuck == -1 {
			fmt.Println("guard stuck bc of anomaly at", x, y)
			anomalies++
		}
		lab = utils.DeepCopy(labInit) // Reset lab to its initial state
	}
	fmt.Println("anomalies:", anomalies)
	fmt.Println("route len", len(route))
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
	clearScreen()
	printLab()
	time.Sleep(milli * time.Millisecond)
}

func printLab() {
	fmt.Println("\nLab:")
	for _, f := range lab {
		fmt.Println(string(f))
	}
}

func move(currPos *[]int) int {
	pos := *currPos
	x, y := pos[0], pos[1]
	guard := lab[x][y]
	if len(isLooped) > 200 {
		isLooped = isLooped[len(isLooped)-200:]
	}
	stuck := slices.Index(isLooped, fmt.Sprintf("%d%d%c", x, y, guard))
	switch guard {
	case '^': // Move up
		if x-1 < 0 {
			return 0
		}
		if stuck != -1 {
			return -1
		}
		if lab[x-1][y] == '#' || lab[x-1][y] == 'O' {
			lab[x][y] = 'X'
			route = append(route, []int{x, y})
			lab[x][y+1] = '>'
			(*currPos)[1] = y + 1
		} else if lab[x-1][y] == '.' || lab[x-1][y] == 'X' {
			lab[x][y] = 'X'
			route = append(route, []int{x, y})
			lab[x-1][y] = '^'
			(*currPos)[0] = x - 1
		}
	case '>': // Move right
		if y+1 >= len(lab[x]) {
			return 0
		}
		if stuck != -1 {
			return -1
		}
		if lab[x][y+1] == '#' || lab[x][y+1] == 'O' {
			lab[x][y] = 'X'
			route = append(route, []int{x, y})
			lab[x+1][y] = 'v'
			(*currPos)[0] = x + 1
		} else if lab[x][y+1] == '.' || lab[x][y+1] == 'X' {
			lab[x][y] = 'X'
			route = append(route, []int{x, y})
			lab[x][y+1] = '>'
			(*currPos)[1] = y + 1
		}
	case 'v': // Move down
		if x+1 >= len(lab) {
			return 0
		}
		if stuck != -1 {
			return -1
		}
		if lab[x+1][y] == '#' || lab[x+1][y] == 'O' {
			lab[x][y] = 'X'
			route = append(route, []int{x, y})
			lab[x][y-1] = '<'
			(*currPos)[1] = y - 1
		} else if lab[x+1][y] == '.' || lab[x+1][y] == 'X' {
			lab[x][y] = 'X'
			route = append(route, []int{x, y})
			lab[x+1][y] = 'v'
			(*currPos)[0] = x + 1
		}
	case '<': // Move left
		if y-1 < 0 {
			return 0
		}
		if stuck != -1 {
			return -1
		}
		if lab[x][y-1] == '#' || lab[x][y-1] == 'O' {
			lab[x][y] = 'X'
			route = append(route, []int{x, y})
			lab[x-1][y] = '^'
			(*currPos)[0] = x - 1
		} else if lab[x][y-1] == '.' || lab[x][y-1] == 'X' {
			lab[x][y] = 'X'
			route = append(route, []int{x, y})
			lab[x][y-1] = '<'
			(*currPos)[1] = y - 1
		}
	}
	isLooped = append(isLooped, fmt.Sprintf("%d%d%c", x, y, guard))
	return 1
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
