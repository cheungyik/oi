package algo_test

import (
	"testing"

	"oi/algo"
)

func TestCombinations(t *testing.T) {
	// Test the Combinations function.
	n := []int{0, 1, 2, 3, 4}
	k := 3
	result := algo.Combinations(n, k)
	expected := [][]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 1, 4},
		{0, 2, 3},
		{0, 2, 4},
		{0, 3, 4},
		{1, 2, 3},
		{1, 2, 4},
		{1, 3, 4},
		{2, 3, 4},
	}
	if len(result) != len(expected) {
		t.Fatalf("expected: %v, got: %v", expected, result)
	}
	for i := 0; i < len(result); i++ {
		if len(result[i]) != len(expected[i]) {
			t.Fatalf("expected: %v, got: %v", expected, result)
		}
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] != expected[i][j] {
				t.Fatalf("expected: %v, got: %v", expected, result)
			}
		}
	}
}
