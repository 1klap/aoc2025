package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	input := []string{}
	file, err := os.ReadFile("day04/in.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = strings.Split(string(file), "\n")
	start := time.Now()
	part1Result := Part1(input)
	fmt.Println(fmt.Sprintf("part1 result=%d (%dms)", part1Result, time.Since(start).Milliseconds()))
	start = time.Now()
	part2Result := Part2(input)
	fmt.Println(fmt.Sprintf("part2 result=%d (%dms)", part2Result, time.Since(start).Milliseconds()))
}

var SurroundingVectors = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func Part1(input []string) int {
	height := len(input)
	width := len(input[0])
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
	}
	for y, row := range input {
		rowRunes := []rune(row)
		for x, cell := range rowRunes {
			grid[y][x] = cell
		}
		//fmt.Printf("%v: %c\n", y, grid[y])
	}

	accessibleCount := 0
	for y, row := range grid {
		//rowRunes := []rune(row)
		for x, cell := range row {
			if cell != '@' {
				continue
			}

			surroundingCells := []rune{}
			paperCount := 0
			for _, v := range SurroundingVectors {
				if x+v[0] >= 0 && y+v[1] >= 0 && x+v[0] < width && y+v[1] < height {
					//fmt.Printf("-- v=%v\n", v)
					surroundingCells = append(surroundingCells, grid[y+v[1]][x+v[0]])
				}
			}
			for _, cell := range surroundingCells {
				if cell == '@' {
					paperCount++
				}
			}
			if paperCount < 4 {
				accessibleCount++
			}
			fmt.Printf("%v,%v: c=%c p=%v s=%c\n", x, y, cell, paperCount, surroundingCells)
		}
	}

	return accessibleCount
}

func Part2(input []string) int {
	height := len(input)
	width := len(input[0])
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
	}
	for y, row := range input {
		rowRunes := []rune(row)
		for x, cell := range rowRunes {
			grid[y][x] = cell
		}
		//fmt.Printf("%v: %c\n", y, grid[y])
	}

	removedCount := 0
	//continueIter := true
	for true {
		positions := accessiblePositions(grid)
		if len(positions) == 0 {
			//continueIter = false
			return removedCount
		} else {
			removedCount = removedCount + len(positions)
			for _, position := range positions {
				grid[position[1]][position[0]] = '#'
			}
		}
		fmt.Printf("r=%v new grid\n", removedCount)
		printGrid(grid)
		fmt.Printf("++++++++\n")
	}

	return 0
}

func accessiblePositions(grid [][]rune) [][]int {
	height := len(grid)
	width := len(grid[0])
	accessiblePos := [][]int{}
	for y, row := range grid {
		//rowRunes := []rune(row)
		for x, cell := range row {
			if cell != '@' {
				continue
			}

			surroundingCells := []rune{}
			paperCount := 0
			for _, v := range SurroundingVectors {
				if x+v[0] >= 0 && y+v[1] >= 0 && x+v[0] < width && y+v[1] < height {
					//fmt.Printf("-- v=%v\n", v)
					surroundingCells = append(surroundingCells, grid[y+v[1]][x+v[0]])
				}
			}
			for _, cell := range surroundingCells {
				if cell == '@' {
					paperCount++
				}
			}
			if paperCount < 4 {
				accessiblePos = append(accessiblePos, []int{x, y})
			}
			//fmt.Printf("%v,%v: c=%c p=%v s=%c\n", x, y, cell, paperCount, surroundingCells)
		}
	}
	return accessiblePos
}

func printGrid(grid [][]rune) {
	for y, row := range grid {
		fmt.Printf("%v: %c\n", y, row)
	}

}
