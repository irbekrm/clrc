package main

import (
	"fmt"

	"github.com/irbekrm/clrc/ch1"
)

func main() {
	// test insertion sort
	x := []int{4, 9, 34, 0, 2}
	ch1.InsertionSort(x)
	fmt.Println(x)
}
