package main

import (
	 "fmt"
	 "math/rand"
	 "time"
)

const arrayLength = 11
var sortedArray = [arrayLength]int{1,2,3,4,5,6,7,8,9,10,11}

func isEqual(a, b [arrayLength]int) bool {
	if a == b {
		return true
	}
	return false
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func shuffleArray(input [arrayLength]int) [arrayLength]int {
	rand.Shuffle(len(input), func(i, j int) { input[i], input[j] = input[j], input[i] })
	return input
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var arrayToSort = sortedArray

	var startTime = time.Now()

	arrayToSort = shuffleArray(arrayToSort)
	fmt.Printf("Original: %v\n", arrayToSort)

	var isSorted bool = false
	var iterations int = 0
	
	for !isSorted {
		arrayToSort = shuffleArray(arrayToSort)
		if isEqual(sortedArray, arrayToSort) {
			isSorted = true
			var endTime = time.Now()
			var elapsedTime = endTime.Sub(startTime)
			fmt.Printf("Result: %v\n", arrayToSort)
			fmt.Printf("Iterations: %v\n", iterations)
			fmt.Printf("Time to complete: %v\n", elapsedTime)
		}
		iterations++
	}
}
