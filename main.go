package main

import (
	"fmt"
	"time"

	"github.com/imedgar/aoc24-imedgar/day10"
)

func main() {
	startTime := time.Now()
	day10.Day10()
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution Time: %s\n", elapsedTime)
}
