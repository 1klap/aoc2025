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
	file, err := os.ReadFile("day03/in.txt")
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

func Part1(input []string) int {
	maxJoltage := make([]int, len(input))
	for i, bank := range input {
		bankRunes := []rune(bank)
		bankInts := make([]int, len(bankRunes))
		for index, bankRune := range bankRunes {
			bankInts[index] = int(bankRune - '0')
		}
		fmt.Printf("bankInts=%v\n", bankInts)
		maxJoltage[i] = findMax(bankInts)
		fmt.Printf("max=%v\n", maxJoltage[i])
	}
	return sumArray(maxJoltage)
}

func sumArray(numbers []int) int {
	var result int = 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func sumArray2(numbers []int64) int64 {
	var result int64 = 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func findMax(numbers []int) int {
	arrayLength := len(numbers)
	maxVal := (numbers[0] * 10) + numbers[1]
	for i := 3; i <= arrayLength; i++ {
		fmt.Printf("-- next i=%v len=%v\n", i, arrayLength)
		firstMaxIndex, firstMax := findLocalMax(numbers[0 : i-1])
		fmt.Printf("debug firstMaxIndex=%v firstMax=%v\n", firstMaxIndex, firstMax)
		_, secondMax := findLocalMax(numbers[firstMaxIndex+1 : i])
		fmt.Printf("debug secondMax=%v\n", secondMax)
		candidate := (firstMax * 10) + secondMax
		if candidate > maxVal {
			maxVal = candidate
			fmt.Printf("* new local max=%v len=%v\n", maxVal, i)
		}
	}
	return maxVal
}

func findLocalMax(numbers []int) (int, int) {
	maxVal := 0 //numbers[0]
	maxIndex := 0
	for i, val := range numbers {
		if val > maxVal {
			maxVal = val
			maxIndex = i
			//fmt.Printf("new local max=%v maxIndex=%v\n", maxVal, maxIndex)
		}
	}
	return maxIndex, maxVal
}

func Part2(input []string) int64 {
	maxJoltage := make([]int64, len(input))
	for i, bank := range input {
		bankRunes := []rune(bank)
		bankInts := make([]int, len(bankRunes))
		for index, bankRune := range bankRunes {
			bankInts[index] = int(bankRune - '0')
		}
		fmt.Printf("bankInts=%v\n", bankInts)
		maxJoltage[i] = findMaxV2(bankInts, 12)
		fmt.Printf("max=%v\n", maxJoltage[i])
	}
	return sumArray2(maxJoltage)
}

func findMaxP2(numbers []int) int64 {
	arrayLength := len(numbers)
	maxVal := arrayToNumber(numbers[0:12])
	batteryCount := 12
	for i := 13; i <= arrayLength; i++ {
		currOffset := 0
		indexArray := make([]int, batteryCount)
		valueArray := make([]int, batteryCount)
		fmt.Printf("-- next i=%v len=%v\n", i, arrayLength)
		// TODO: store n max and index in array, do this 12 times
		for b := 0; b < batteryCount; b++ {
			fmt.Printf("_ b=%v currOffset=%v\n", b, currOffset)
			// [0:]
			currMaxIndex, currMax := findLocalMax(numbers[currOffset:(batteryCount - (b + 1))])
			indexArray[b] = currMaxIndex
			valueArray[b] = currMax
			currOffset = currMaxIndex + 1
		}
		//firstMaxIndex, firstMax := findLocalMax(numbers[0 : i-1])
		//fmt.Printf("debug firstMaxIndex=%v firstMax=%v\n", firstMaxIndex, firstMax)
		//_, secondMax := findLocalMax(numbers[firstMaxIndex+1 : i])
		//fmt.Printf("debug secondMax=%v\n", secondMax)
		// TODO compute candidate from array
		candidate := arrayToNumber(valueArray)
		//candidate := (firstMax * 10) + secondMax
		if candidate > maxVal {
			maxVal = candidate
			fmt.Printf("* new local max=%v len=%v\n", maxVal, i)
		}
	}
	return maxVal
}

func findMaxV2(values []int, length int) int64 {
	startIndex := 0
	maxIndex := -1
	maxVal := -1
	endIndex := (len(values) - length) + 1
	//endIndex := length
	picks := make([]int, length)
	for i := 0; i < length; i++ {
		slice := values[startIndex:endIndex]
		maxIndex, maxVal = findLocalMax(slice)
		fmt.Printf("slice=%v / pos=%v / maxIndex=%v maxValue=%v si=%v ei=%v\n", slice, i, maxIndex, maxVal, startIndex, endIndex)
		picks[i] = maxVal
		startIndex = startIndex + maxIndex + 1
		if endIndex < len(values) {
			endIndex++
		}
		fmt.Printf("* pos=%v max=%v si=%v ei=%v\n", i, maxVal, startIndex, endIndex)
	}
	fmt.Printf("picks=%v\n", picks)
	fmt.Printf("picksNum=%v\n", arrayToNumber(picks))
	return arrayToNumber(picks)
}

func arrayToNumber(array []int) int64 {
	result := int64(0)
	for i := 0; i < len(array); i++ {
		result = (result * 10) + int64(array[i])
	}
	return result
}

// P2 : 17384857711727 too low
