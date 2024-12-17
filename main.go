package main

import (
	"fmt"
	"time"

	"github.com/imedgar/aoc24-imedgar/day11"
)

func main() {
	startTime := time.Now()
	day11.Day11()
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution Time: %s\n", elapsedTime)
}
