package utilities

import "strconv"

func AlphaToIntLike[T ~int](s string) (T, error) {
	i, err := strconv.Atoi(s)
	return T(i), err
}

// SliceStringsToIntsLike takes a slice of strings and attempts to convert
// to a slice of an integer-like type.
func SliceStringsToIntsLike[T ~int](numStrs []string) ([]T, error) {
	if numStrs == nil {
		return nil, nil
	}
	var err error
	nums := make([]T, len(numStrs))
	for idx, s := range numStrs {
		if nums[idx], err = AlphaToIntLike[T](s); err != nil {
			return nil, err
		}
	}
	return nums, nil
}

// SliceStringsToInts takes a slice of strings and attempts to convert to a
// slice of integers.
func SliceStringsToInts(numStrs []string) ([]int, error) {
	return SliceStringsToIntsLike[int](numStrs)
}
