package service

import "GIN-DEMO/utils/sort"

func SortService(elements []int) []int {

	if len(elements) <= 10000 {
		return sort.BubbleSortAscendingOrder(elements)
	}
	return sort.SortDefaultProvided(elements)

}
