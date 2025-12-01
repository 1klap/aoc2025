package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := []string{}
	file, err := os.ReadFile("day01/in.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = strings.Split(string(file), "\n")
	start := time.Now()
	part1Result := Part1(input)
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("part1 result=%d (%dμs)", part1Result, elapsed.Microseconds()))
	start = time.Now()
	part2Result := Part2(input)
	elapsed = time.Since(start)
	fmt.Println(fmt.Sprintf("part2 result=%d (%dμs)", part2Result, elapsed.Microseconds()))
}

func Part1(input []string) int {
	dial := 50
	zeroCount := 0
	for _, turn := range input {
		directionRight := turn[0:1] == "R"
		turnAmplitude, err := strconv.Atoi(turn[1:])
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(fmt.Sprintf("#%d: %t %d", index, directionRight, turnAmplitude))
		if directionRight {
			dial = (dial + turnAmplitude) % 100
		} else {
			dial = dial - turnAmplitude
			if dial < 0 {
				dial = (dial + 100) % 100 // normalize negative dial
			}
		}
		if dial == 0 {
			zeroCount++
		}
		//fmt.Println(fmt.Sprintf("#%d after: dial=%d", index, dial))
	}
	return zeroCount
}

func Part2(input []string) int {
	dial := 50
	zeroCount := 0
	for _, turn := range input {
		directionRight := turn[0:1] == "R"
		turnAmplitude, err := strconv.Atoi(turn[1:])
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(fmt.Sprintf("#%d: %t %d", index, directionRight, turnAmplitude))
		if directionRight {
			zeroCount += (dial + turnAmplitude) / 100
			dial = (dial + turnAmplitude) % 100
		} else {
			zeroCount += turnAmplitude / 100
			lastTurn := turnAmplitude % 100
			if lastTurn >= dial && dial != 0 {
				zeroCount += 1
			}
			dial = dial - lastTurn
			if dial < 0 {
				dial = (dial + 100) % 100 // normalize negative dial
			}
		}
		//fmt.Println(fmt.Sprintf("after: dial=%d", dial))
	}
	return zeroCount
}
