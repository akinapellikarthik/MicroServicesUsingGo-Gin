package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	addFn func(int, int) int
	subFn func(int, int) int
)

type groceryServicesMock struct{}

func (g groceryServicesMock) Add(i int, i2 int) int {
	return addFn(i, i2)
}

func (g groceryServicesMock) Sub(i int, i2 int) int {
	return subFn(i, i2)
}

func TestAdd(t *testing.T) {
	addFn = func(i int, i2 int) int {
		fmt.Println("Inside the test mock method of add")
		return i + i2
	}
	GService = &groceryServicesMock{}

	sum := GService.Add(10, 20)

	assert.Equal(t, 30, sum)

}
