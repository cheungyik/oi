package algo

// Combinations generates all possible combinations of k elements from a given slice of any type.
//
// This function uses recursion to explore all possible selections of k elements from the input slice `elements`.
// It supports any type of elements, making it versatile for combinations of strings, numbers, structs, etc.
//
// T represents the generic type, allowing the function to handle slices of any data type.
//
// Parameters:
//   - elements: A slice of any type T from which combinations will be generated.
//   - k: The number of elements to include in each combination.
//
// Returns:
//   - A slice of slices, where each inner slice represents a combination of k elements.
//
// Example:
//
//	// Example 1: Combinations of strings
//	strings := []string{"apple", "banana", "cherry"}
//	result := Combinations(strings, 2)
//	fmt.Println(result)
//	// Output: [["apple", "banana"], ["apple", "cherry"], ["banana", "cherry"]]
//
//	// Example 2: Combinations of integers
//	integers := []int{1, 2, 3, 4}
//	result := Combinations(integers, 3)
//	fmt.Println(result)
//	// Output: [[1, 2, 3], [1, 2, 4], [1, 3, 4], [2, 3, 4]]
//
//	// Example 3: Combinations of custom structs
//	type Point struct { X, Y int }
//	points := []Point{{1, 2}, {3, 4}, {5, 6}}
//	result := Combinations(points, 2)
//	fmt.Println(result)
//	// Output: [[{1 2} {3 4}], [{1 2} {5 6}], [{3 4} {5 6}]]
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
