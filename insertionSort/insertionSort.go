package insertionSort

func Sort(a []int) []int {
	// creating a new slice of a size equal to the input slice
	result := make([]int, len(a))
	// copying elements from the orignal slice to the new slice
	copy(result, a)
	// iteration through all elements of the slice starting with the second
	for i := 1; i < len(result); i++ {
		// key holds the current value
		key := result[i]
		// j holds the index of the previous value
		j := i - 1
		// iterate backwards starting with the prevoius value and compare it
		// with the current value
		for j >= 0 && result[j] > key {
			// if the previous value is greater, then we move it
			// forward one position
			result[j+1] = result[j]
			// we then decrement the index of the previous value so that
			// we check the next value backwards
			j--
		}
		// at the end of the inner loop we want to potentially move
		// the current calue of the main loop to a new position (if
		// anything has shifted above )
		result[j+1] = key
	}
	return result
}
