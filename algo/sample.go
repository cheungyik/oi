package algo

import (
	"math/rand"
)

// Samples returns N random unique items from collection.
func Samples[T any](collection []T, count int) []T {
	size := len(collection)
	tmp := append([]T{}, collection...)
	res := make([]T, 0)
	for i := 0; i < size && i < count; i++ {
		length := size - i
		index := rand.Intn(size - i)
		res = append(res, tmp[index])
		tmp[index] = tmp[length-1]
		tmp = tmp[:length-1]
	}
	return res
}
