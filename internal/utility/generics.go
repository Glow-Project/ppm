package utility

// get the index of a certain item in a string slice
func IndexOf[T comparable](target T, data []T) int {
	for index, value := range data {
		if target == value {
			return index
		}
	}
	return -1
}

// check whether a certain item exists in a string slice
func SliceContains[T comparable](target T, data []T) bool {
	for _, v := range data {
		if v == target {
			return true
		}
	}

	return false
}

// get the first predicate match of the slice
//
// returns nil if none of the items match
func GetFirstMatch[T any](arr []T, predicate func(T, int) bool) *T {
	for i, value := range arr {
		if predicate(value, i) {
			return &value
		}
	}

	return nil
}

// returns a new slice containing all the items on which the predicate returns true
func Filter[T any](arr []T, predicate func(T, int) bool) []T {
	newArr := []T{}
	for i, item := range arr {
		if predicate(item, i) {
			newArr = append(newArr, item)
		}
	}

	return newArr
}

func Map[T any, N any](arr []T, predicate func(T, int) N) []N {
	newArr := []N{}
	for i, item := range arr {
		newArr = append(newArr, predicate(item, i))
	}

	return newArr
}

func Some[T any](arr []T, predicate func(T, int) bool) bool {
	for i, item := range arr {
		if predicate(item, i) {
			return true
		}
	}

	return false
}
