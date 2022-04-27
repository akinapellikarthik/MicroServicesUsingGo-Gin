package utils

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGetElements(t *testing.T) {
	input := 10
	expectedOutput := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	testOutCome := GetElements(input)
	assert.Equal(t, expectedOutput, testOutCome)
}
