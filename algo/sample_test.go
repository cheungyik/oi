package algo

import (
	"fmt"
	"testing"
)

func TestSamples(t *testing.T) {

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	samples := Samples(nums, 5)

	fmt.Println(nums, samples)
}
