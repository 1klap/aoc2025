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
	file, err := os.ReadFile("day05/in.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = strings.Split(string(file), "\n\n")
	start := time.Now()
	//part1Result := Part1(input)
	//fmt.Println(fmt.Sprintf("part1 result=%d (%dms)", part1Result, time.Since(start).Milliseconds()))
	//start = time.Now()
	part2Result := Part2(input)
	fmt.Println(fmt.Sprintf("part2 result=%d (%dms)", part2Result, time.Since(start).Milliseconds()))
}

func Part1(input []string) int {
	fmt.Println("part 1 ranges=", input[0])
	fmt.Println("part 1 ingredients=", input[1])
	//rangeStrings := strings.Split(input[0], "\n")
	//for _, rangeString := range rangeStrings {
	//	interval := mapSlice(strings.Split(rangeString, "-"), func(s string) int {
	//		strconv.Atoi(s)
	//	})
	//}
	ranges := mapSlice(strings.Split(input[0], "\n"), func(rangeString string) []int {
		return mapSlice(strings.Split(rangeString, "-"), func(s string) int {
			converted, _ := strconv.Atoi(s)
			return converted
		})
	})
	//fmt.Println("part 1 intRanges=", ranges)
	ingredients := mapSlice(strings.Split(input[1], "\n"), func(ingredientString string) int {
		converted, _ := strconv.Atoi(ingredientString)
		return converted
	})
	fmt.Println("part 1 ingredients=", ingredients)
	freshCount := 0
	for _, ingredient := range ingredients {
		if isFresh(ingredient, &ranges) {
			freshCount++
		}
	}
	return freshCount
}

func isFresh(ingredient int, ranges *[][]int) bool {
	for _, rnge := range *ranges {
		if (ingredient >= rnge[0]) && (ingredient <= rnge[1]) {
			//fmt.Printf("! fresh i=%v r=%v\n", ingredient, rnge)
			return true
		}
	}
	return false
}

func Part2(input []string) int {
	ranges := mapSlice(strings.Split(input[0], "\n"), func(rangeString string) []int {
		return mapSlice(strings.Split(rangeString, "-"), func(s string) int {
			converted, _ := strconv.Atoi(s)
			return converted
		})
	})

	for i := len(ranges) - 1; i > 1; i-- {
		rngeA := ranges[i]
		//for j := 0; j < i; j++ {
		for j := i - 1; j >= 0; j-- {
			rngeB := ranges[j]
			//if (rngeA[0] <= rngeB[0] && rngeA[1] > rngeB[0] && rngeA[1] <= rngeB[1]) || (rngeA[0] >= rngeB[0] && rngeA[0] < rngeB[1] && rngeA[1] >= rngeB[1]) {
			//if (rngeA[0] == rngeB[1]) || (rngeA[1] == rngeB[0]) || (rngeA[0] >= rngeB[0] && rngeA[0] <= rngeB[1]) || (rngeA[1] >= rngeB[0] && rngeA[1] <= rngeB[1]) || (rngeA[0] >= rngeB[0] && rngeA[1] <= rngeB[1]) || (rngeA[0] <= rngeB[0] && rngeA[1] >= rngeB[1]) {
			if (rngeA[0] >= rngeB[0] && rngeA[0] <= rngeB[1]) || (rngeA[1] >= rngeB[0] && rngeA[1] <= rngeB[1]) || (rngeA[0] >= rngeB[0] && rngeA[1] <= rngeB[1]) || (rngeA[0] <= rngeB[0] && rngeA[1] >= rngeB[1]) {
				fmt.Printf("merging A=%v B=%v ", rngeA, rngeB)
				rngeB[0], rngeB[1] = min(rngeA[0], rngeB[0]), max(rngeA[1], rngeB[1])
				fmt.Printf("B'=%v\n", rngeB)
				//ranges[j] = rngeB // not needed
				ranges = ranges[:len(ranges)-1]
				break
			}
		}
	}
	fmt.Printf("ranges=%v\n", ranges)

	freshCount := 0
	for _, rnge := range ranges {
		freshCount = freshCount + (rnge[1] + 1 - rnge[0])
	}

	return freshCount
}

func mapSlice[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

// 315957683089678 too low
// 317728688359509 too low
// 422408896659291 not right
// 6655443634006
// 344025747660214 not right

// bound: 407477686692412

//241645066313
//9785310865148
//342866708370012
//346965358558799
