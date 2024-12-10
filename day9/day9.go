package day9

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/imedgar/aoc24-imedgar/utils"
)

var (
	disk    = []int{}
	output  = []string{}
	output2 = []string{}
)

func Day9() {
	file := utils.ReadFile("./day9/input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		disk = append(disk, utils.StrSliceToInt(strings.Split(line, ""))...)
	}
	deblock()
	output2 = make([]string, len(output))
	copy(output2, output)
	// part1
	//	moveFileBlocks()
	//	fmt.Println(getChecksum(output))
	// part2
	moveWholeFileBlocks()
	fmt.Println(getChecksum(output2))
}

func deblock() {
	id := 0
	for i, v := range disk {
		add := ""
		if i%2 != 0 {
			add = "."
		} else {
			add = strconv.Itoa(id)
			id++
		}
		addTimesN(v, add)
	}
}

func addTimesN(t int, v string) {
	for j := 0; j < t; j++ {
		output = append(output, v)
	}
}

func moveFileBlocks() {
	for slices.Contains(output, ".") {
		for i := len(output) - 1; i > 0; i-- {
			block := output[i]
			if block != "." {
				emptyBlockIdx := slices.Index(output, ".")
				output[emptyBlockIdx] = block
				output = utils.RemoveElementFrom(output, i)
				break
			}
		}
	}
}

func moveWholeFileBlocks() {
	for i := len(output2) - 1; i > 0; i-- {
		block := output2[i]
		if block != "." {
			firstIdIdx := slices.Index(output2, block)
			slotNeeded := i - firstIdIdx + 1
			mvA, mvB := checkFreeSpace(slotNeeded)
			if (mvA == -1 && mvB == -1) || mvB > i {
				i = firstIdIdx
				continue
			}
			for m := 0; m < slotNeeded; m++ {
				output2[mvA+m] = block
			}
			for r := firstIdIdx; r <= i; r++ {
				output2[r] = "."
			}
		}
	}
}

func checkFreeSpace(n int) (int, int) {
	size := 0
	for i, v := range output2 {
		if size >= n {
			return i - n, i - 1
		}
		if v == "." {
			size++
		} else {
			size = 0
		}
	}
	return -1, -1
}

func getChecksum(fs []string) int {
	checksum := 0
	for i, v := range fs {
		if v != "." {
			checksum += i * utils.StrToInt(v)
		}
	}
	return checksum
}
