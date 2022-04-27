package service

import (
	"testing"
	"time"
)

func TestSortService(t *testing.T) {
	commChan := make(chan bool, 1)
	defer close(commChan)
	elements := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var sortedElements []int
	go func() {
		sortedElements = SortService(elements)
		commChan <- false
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		commChan <- true
	}()

	if <-commChan {
		t.Error("Bubble Sort taken more than 50 ms to run and hence test case is aborted")
		return
	}

	if 1 != sortedElements[0] {
		t.Error("elements are not sorted")
	}
}
