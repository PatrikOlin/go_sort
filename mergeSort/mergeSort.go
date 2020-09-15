package mergeSort

func Sort(input []int) []int {
	// makin a new slice and copying the contents of the input into it
	// this slice is going to be subsequently mutated so that
	// we get the desired order of elements
	sorted := make([]int, len(input))
	copy(sorted, input)

	// calling the private subSort function which can sort a sub-slice
	// by taking two more arguments: the start and the end index
	subSort(sorted, 0, len(input)-1)
	return sorted
}

// this function is going to call itself recursively with the left
// and right halves of the input slice (the divide)
// then call the merge function (the conquest)
func subSort(sorted []int, leftStart, rightEnd int) {
	// stop the recursion if there is nothing to divide
	if leftStart >= rightEnd {
		return
	}

	// calculatin the middle element so we can divide our slice
	middle := (leftStart + rightEnd) / 2
	// calling itself recursively with both halves
	subSort(sorted, leftStart, middle)
	subSort(sorted, middle+1, rightEnd)
	// merging the two sorted halves
	merge(sorted, leftStart, rightEnd)
}

func merge(sorted []int, leftStart, rightEnd int) {
	// creating a temporary slice, as we cant easily do it in place
	temp := make([]int, len(sorted))
	copy(temp, sorted)

	// the end of the left sub-slice will end in the middle
	leftEnd := (leftStart + rightEnd) / 2
	// the start of the right slice will be right after the left end
	rightStart := leftEnd + 1

	left := leftStart
	right := rightStart
	index := leftStart

	// this is the loop where the actual sorting happens
	// we iterate until either the left or the right sub-slice
	// runs out of elements
	for left <= leftEnd && right <= rightEnd {
		// here we start adding elements to the temporary
		// slice we created aboice
		// we choose the smaller element every time
		if sorted[left] < sorted[right] {
			temp[index] = sorted[left]
			left++
		} else {
			temp[index] = sorted[right]
			right++
		}
		index++
	}

	// here we append to the temp slice the remaining elements
	// that were not included in the loop above
	// first we do it for the left and then for the right sub-slice
	// one of them will not contain any remaining elements
	// and will therefore make no changes
	copy(temp[index:rightEnd+1], sorted[right:rightEnd+1])
	copy(temp[index:rightEnd+1], sorted[left:leftEnd+1])
	// finally we store the sorted numbers from the temporary slices
	// into our osrted slice
	copy(sorted, temp)
}
