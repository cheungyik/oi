package algo

// Combinations generates all possible combinations of k elements from a range of n elements (0 to n-1).
// If k is 0 or greater than n, it returns an empty slice.
//
// Parameters:
// - n: the size of the range (0 to n-1).
// - k: the number of elements in each combination.
//
// Returns:
// - A 2D slice containing all combinations, where each combination is represented as a slice of integers.
//
// Example:
// n := [0, 1, 2, 3, 4]
// k := 3
// result := Combinations(n, k)
// Output:
// [
//
//	[0, 1, 2],
//	[0, 1, 3],
//	[0, 1, 4],
//	[0, 2, 3],
//	[0, 2, 4],
//	[0, 3, 4],
//	[1, 2, 3],
//	[1, 2, 4],
//	[1, 3, 4],
//	[2, 3, 4]
//
// ]
func Combinations[T any](elements []T, k int) [][]T {
	if k == 0 || k > len(elements) {
		return [][]T{}
	}
	var result [][]T
	var comb []T
	var combine func(start, depth int)
	combine = func(start, depth int) {
		if depth == k {
			// Make a copy of the combination and add it to the result.
			result = append(result, append([]T(nil), comb...))
			return
		}
		for i := start; i < len(elements); i++ {
			comb = append(comb, elements[i])
			combine(i+1, depth+1)
			comb = comb[:len(comb)-1]
		}
	}
	combine(0, 0)
	return result
}
