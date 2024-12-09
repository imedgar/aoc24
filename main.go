package main

import (
	"fmt"
	"time"

	"github.com/imedgar/aoc24-imedgar/day8"
)

func main() {
	startTime := time.Now()
	day8.Day8()
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution Time: %s\n", elapsedTime)
}
