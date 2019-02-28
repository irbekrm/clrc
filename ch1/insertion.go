package ch1

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(unsorted []int) {
	var i, j, key int
	for j = 1; j < len(unsorted); j++ {
		key = unsorted[j]
		for i = j - 1; i >= 0 && unsorted[i] < key; i-- {
			unsorted[i+1] = unsorted[i]
		}
		unsorted[i+1] = key
	}
}

func TestInsertionSort() {
	arr := make([]int, 100)
	t := time.Now().UnixNano()
	s := rand.NewSource(t)
	r := rand.New(s)

	// Generating an array of 100 random integers [1 - 1000]
	for i := 0; i < 100; i++ {
		arr[i] = r.Intn(1000)
	}
	start := time.Now()
	fmt.Printf("\nUnsorted array: \n%v\n\n\n", arr)
	insertionSort(arr)
	timeTaken := time.Since(start)
	fmt.Printf("Sorted array: \n%v\n\n\n", arr)
	fmt.Printf("Time taken to sort: %v\n", timeTaken)
}
