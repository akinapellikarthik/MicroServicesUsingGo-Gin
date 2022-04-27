package sort

import (
	"GIN-DEMO/utils"
	"github.com/stretchr/testify/assert"

	//"github.com/go-playground/assert/v2"
	"testing"
)

func TestBubbleSortNoError(t *testing.T) {
	elements := []int{9, 8, 7, 6, 5, 3, 1, 0, 11}
	expectedSortedElements := []int{0, 1, 3, 5, 6, 7, 8, 9, 11}
	sortedElements := BubbleSortAscendingOrder(elements)
	assert.Equal(t, expectedSortedElements, sortedElements)

	if sortedElements[0] != 0 {
		t.Error("Not a valid sort")
	}
}

func BenchmarkBubbleSortAscendingOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSortAscendingOrder(utils.GetElements(10.000))
	}
}

func BenchmarkSortDefaultProvided(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortDefaultProvided(utils.GetElements(10.000))
	}
}
