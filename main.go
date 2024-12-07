package main

import (
	"fmt"
	"time"

	"github.com/imedgar/aoc24-imedgar/day1"
	"github.com/imedgar/aoc24-imedgar/day2"
	"github.com/imedgar/aoc24-imedgar/day3"
	"github.com/imedgar/aoc24-imedgar/day4"
	"github.com/imedgar/aoc24-imedgar/day5"
	"github.com/imedgar/aoc24-imedgar/day6"
)

func main() {
	startTime := time.Now()
	fmt.Println("\nDay 1")
	day1.Day1()
	fmt.Println("\nDay 2")
	day2.Day2()
	fmt.Println("\nDay 3")
	day3.Day03()
	fmt.Println("\nDay 4")
	day4.Day4()
	fmt.Println("\nDay 5")
	day5.Day5()
	fmt.Println("\nDay 6")
	day6.Day6()
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution Time: %s\n", elapsedTime)
}
