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
	file, err := os.ReadFile("day02/in.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = strings.Split(string(file), ",")
	start := time.Now()
	part1Result := Part1(input)
	fmt.Println(fmt.Sprintf("part1 result=%d (%dms)", part1Result, time.Since(start).Milliseconds()))
	start = time.Now()
	part2Result := Part2(input)
	fmt.Println(fmt.Sprintf("part2 result=%d (%dms)", part2Result, time.Since(start).Milliseconds()))
}

func Part1(input []string) int64 {
	symetricIds := []int64{}
	var startId int64
	var endId int64
	for _, idRange := range input {
		ids := strings.Split(idRange, "-")
		//startId, _ = strconv.Atoi(ids[0])
		startId, _ = strconv.ParseInt(ids[0], 10, 64)
		//endId, _ = strconv.Atoi(ids[1])
		endId, _ = strconv.ParseInt(ids[1], 10, 64)
		//fmt.Printf("------ start=%v end=%v\n", startId, endId)
		for i := startId; i <= endId; i++ {
			if isSymmetric(i) {
				symetricIds = append(symetricIds, i)
			}
		}
	}
	return sumArray(symetricIds)
}

func isSymmetric(id int64) bool {
	//idString := strconv.Itoa(id)
	idString := strconv.FormatInt(id, 10)
	idLeft := idString[0 : len(idString)/2]
	idRight := idString[len(idString)/2:]
	//fmt.Printf("idLeft=%v idRight=%v\n", idLeft, idRight)
	paramIsSymmetric := idLeft == idRight
	if paramIsSymmetric {
		//fmt.Printf("symmetric! %v\n", id)
	}
	return paramIsSymmetric
}

func sumArray(numbers []int64) int64 {
	var result int64 = 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func Part2(input []string) int64 {
	symetricIds := []int64{}
	var startId int64
	var endId int64
	for _, idRange := range input {
		ids := strings.Split(idRange, "-")
		//startId, _ = strconv.Atoi(ids[0])
		startId, _ = strconv.ParseInt(ids[0], 10, 64)
		//endId, _ = strconv.Atoi(ids[1])
		endId, _ = strconv.ParseInt(ids[1], 10, 64)
		//fmt.Printf("------ start=%v end=%v\n", startId, endId)
		for i := startId; i <= endId; i++ {
			if isRepeating(i) {
				symetricIds = append(symetricIds, i)
			}
		}
	}
	return sumArray(symetricIds)
}

func isRepeating(id int64) bool {
	idString := strconv.FormatInt(id, 10)
	for patternLength := 1; patternLength <= (len(idString) / 2); patternLength++ {
		//if len(idString)%patternLength != 0 {
		//return false
		//break
		//}
		pattern := idString[0:patternLength]
		//for i := patternLength; i < len(idString); i = i+patternLength {
		//	subPatternMatched := idString[i:i+patternLength] == pattern
		//	if !subPatternMatched {
		//		break
		//	}
		//}
		if isRepeatingForPattern(idString, pattern) {
			//fmt.Printf("!!! repeating p=%v, id=%v\n", pattern, idString)
			return true
		}
	}
	return false
	//if len(idString)%2 != 0 {
	//	return false
	//}
	//for patternLength := 0; patternLength < len(idString)/2; patternLength += 1 {
	//
	//}
}

func isRepeatingForPattern(idString string, pattern string) bool {
	patternLength := len(pattern)
	if len(idString)%patternLength != 0 {
		return false
	}
	for i := patternLength; i+patternLength <= len(idString); i = i + patternLength {
		//fmt.Printf("check p=%v, i=%v id=%v\n", patternLength, i, idString)
		subPatternMatched := idString[i:i+patternLength] == pattern
		if !subPatternMatched {
			return false
		}
	}
	return true
}

//12121212 p=12
//  23
//    45
//      67
