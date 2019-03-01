package ch1

import (
	"fmt"
	"math/rand"
	"time"
)

func SortIncExample() {
	arr := generateRandomArray()
	start := time.Now()
	SortInc(arr)
	timeTaken := time.Since(start)
	printInfo("increasing", timeTaken, arr)
}

func SortDecExample() {
	arr := generateRandomArray()
	start := time.Now()
	SortDec(arr)
	timeTaken := time.Since(start)
	printInfo("decreasing", timeTaken, arr)
}

// Generates an array of 100 random integers [1 - 1000]
func generateRandomArray() []int {
	arr := make([]int, 100)
	t := time.Now().UnixNano()
	s := rand.NewSource(t)
	r := rand.New(s)
	for i := 0; i < 100; i++ {
		arr[i] = r.Intn(1000)
	}
	return arr
}

func printInfo(order string, timeTaken time.Duration, arr []int) {
	fmt.Printf("\nUnsorted array: \n%v\n\n\n", arr)
	fmt.Printf("Array sorted in %s order: \n%v\n\n\n", order, arr)
	fmt.Printf("Time taken to sort: %v\n", timeTaken)
}
