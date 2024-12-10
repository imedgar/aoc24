package main

import (
	"fmt"
	"time"

	"github.com/imedgar/aoc24-imedgar/day9"
)

func main() {
	startTime := time.Now()
	day9.Day9()
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution Time: %s\n", elapsedTime)
}
