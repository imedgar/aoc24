package day8

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/imedgar/aoc24-imedgar/utils"
)

type Antenna struct {
	x, y int
}

var (
	ant    = make(map[string][]Antenna)
	area   = [][]string{}
	total  = 0
	unique = utils.NewSet[string]()
)

func Day8() {
	file := utils.ReadFile("./day8/input.txt")
	scanner := bufio.NewScanner(file)
	l := 0
	for scanner.Scan() {
		line := scanner.Text()
		floor := []string{}
		for i, c := range strings.Split(line, "") {
			floor = append(floor, c)
			re := regexp.MustCompile(`[a-zA-Z0-9]`)
			matches := re.FindAllString(c, -1)
			if matches != nil {
				antenna := Antenna{
					x: l,
					y: i,
				}

				key := matches[0]
				ant[key] = append(ant[key], antenna)
			}
		}
		area = append(area, floor)
		l++
	}

	part1()
	fmt.Println("total =", total)

	unique = utils.NewSet[string]()
	total = 0

	part2()
	fmt.Println("total =", total)
}

func part1() {
	for k := range ant {
		for _, curr := range ant[k] {
			for _, a := range ant[k] {
				if curr.x == a.x && curr.y == a.y {
					continue
				}
				antiX, antiY := a.x-(curr.x-a.x), a.y-(curr.y-a.y)
				if antiX < 0 || antiX > len(area)-1 || antiY < 0 || antiY > len(area[antiX])-1 {
					continue
				}
				addAntinode(antiX, antiY)
			}
		}
	}
}

func part2() {
	for k := range ant {
		for _, curr := range ant[k] {
			for _, a := range ant[k] {
				if curr.x == a.x && curr.y == a.y {
					continue
				}
				addAntinodes(curr.x, curr.y, a.x, a.y)
			}
		}
	}
}

func addAntinode(antiX, antiY int) {
	key := fmt.Sprintf("%d-%d", antiX, antiY)
	if !unique.Contains(key) {
		unique.Add(key)
		total++
	}
}

func addAntinodes(prevX, prevY, newX, newY int) {
	subsX, subsY := prevX-newX, prevY-newY
	for prevX >= 0 && prevX < len(area) && prevY >= 0 && prevY < len(area[prevX]) {
		addAntinode(prevX, prevY)
		prevX, prevY = prevX-subsX, prevY-subsY
	}
}
