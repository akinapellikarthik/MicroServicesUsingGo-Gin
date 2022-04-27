package utils

func GetElements(n int) []int {
	tempArr := make([]int, n)
	j := 0
	for i := n; i > 0; i-- {
		tempArr[j] = i
		j++
	}

	return tempArr
}
