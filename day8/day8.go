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
	ant    = make(map[string][]Antenna) // Correctly initialize the map
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
	// pair antennas
	for k := range ant {
		for _, curr := range ant[k] {
			for _, a := range ant[k] {
				if curr.x == a.x && curr.y == a.y {
					continue
				}
				antiX, antiY := a.x-(curr.x-a.x), a.y-(curr.y-a.y)
				if antiX < 0 || antiX > len(area)-1 || antiY < 0 || antiY > len(area[antiX])-1 {
					// fmt.Println("antinode out of borders")
					continue
				}
				// fmt.Println("trying antinode for", curr.x, curr.y, "to", a.x, a.y, "at", antiX, antiY)
				potential := area[antiX][antiY]
				if potential == "." {
					//					fmt.Println("current", curr.x, curr.y, "pairing to", a.x, a.y, "at", antiX, antiY)
					area[antiX][antiY] = "#"
					key := fmt.Sprintf("%s-%s", antiX, antiY)
					if !unique.Contains(key) {
						unique.Add(fmt.Sprintf("%s-%s", antiX, antiY))
						total++
					}
				} else if potential != "#" {
					//				fmt.Println("current", curr.x, curr.y, "pairing to", a.x, a.y, "diff", antiX, antiY, "overlaps antenna")
					key := fmt.Sprintf("%s-%s", antiX, antiY)
					if !unique.Contains(key) {
						unique.Add(fmt.Sprintf("%s-%s", antiX, antiY))
						total++
					}
				} else if potential == "#" {
					//			fmt.Println("current", curr.x, curr.y, "pairing to", a.x, a.y, "diff", antiX, antiY, "overlaps antinode")
					// total++
				}
			}
		}
	}
	fmt.Println("total antinodes", total)
	//	for _, f := range area {
	//		fmt.Println(f)
	//	}
}
