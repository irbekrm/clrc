package ch1

// Sorts a slice passed as argument in descending order.
func SortDec(unsorted []int) {
	var i, j, key int
	for j = 1; j < len(unsorted); j++ {
		key = unsorted[j]
		for i = j - 1; i >= 0 && unsorted[i] < key; i-- {
			unsorted[i+1] = unsorted[i]
		}
		unsorted[i+1] = key
	}
}

func SortInc(unsorted []int) {
	var i, j, key int
	for j = 1; j < len(unsorted); j++ {
		key = unsorted[j]
		for i = j - 1; i >= 0 && unsorted[i] > key; i-- {
			unsorted[i+1] = unsorted[i]
		}
		unsorted[i+1] = key
	}
}
