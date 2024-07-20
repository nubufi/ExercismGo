package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

// Keep returns a new slice containing all elements of the input slice that satisfy the predicate f.
func Keep[T interface{}](collection []T, f func(T) bool) []T {
	var keeped []T

	for _, i := range collection {
		if f(i) {
			keeped = append(keeped, i)
		}
	}

	return keeped
}

func Discard[T interface{}](collection []T, f func(T) bool) []T {
	var keeped []T

	for _, i := range collection {
		if !f(i) {
			keeped = append(keeped, i)
		}
	}

	return keeped
}
