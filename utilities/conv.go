package utilities

import "strconv"

// SliceStringsToIntsLike takes a slice of strings and attempts to convert
// to a slice of an integer-like type.
func SliceStringsToIntsLike[T ~int](numStrs []string) ([]T, error) {
	if numStrs == nil {
		return nil, nil
	}

	var n int
	var err error
	nums := make([]T, len(numStrs))

	for idx, s := range numStrs {
		if n, err = strconv.Atoi(s); err != nil {
			return nil, err
		}
		nums[idx] = T(n)
	}

	return nums, nil
}

// SliceStringsToInts takes a slice of strings and attempts to convert to a
// slice of integers.
func SliceStringsToInts(numStrs []string) ([]int, error) {
	return SliceStringsToIntsLike[int](numStrs)
}
