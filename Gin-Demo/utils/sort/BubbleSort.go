package sort

import (
	"sort"
)

//Bubble Sort to sort elements
func BubbleSortAscendingOrder(elements []int) []int {
	keepWorking := true
	for keepWorking {
		//when this line will be commented channels will work and test case will fail
		keepWorking = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepWorking = true
			}
		}
	}

	return elements

}

//Sorts the input using inbuilt goLang sort library
func SortDefaultProvided(elements []int) []int {
	sort.Ints(elements)
	return elements
}
