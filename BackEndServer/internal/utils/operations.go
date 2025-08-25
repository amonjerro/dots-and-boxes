package utils

// Generic reduce operation for iterables
func Reduce[T any, R any](slice []T, reducer func(R, T) R, initialValue R) R {
	accumulator := initialValue
	for _, item := range slice {
		accumulator = reducer(accumulator, item)
	}
	return accumulator
}
