package service

import "testing"

func TestSortService(t *testing.T) {
	elements := []int {10,9,8,7,6,5,4,3,2,1}
	sortedElements := SortService(elements)
	if 1 != sortedElements[0] {
		t.Error("elements are not sorted")
	}
}


