package service

import "fmt"

type groceryServices interface {
	Add(int, int) int
	Sub(int, int) int
}

type GrocerServices struct{}


var (
	GService groceryServices = &GrocerServices{}
)

func (g *GrocerServices) Add(i int, i2 int) int {
	fmt.Println("Inside real implementation")
	return i + i2
}

func (g *GrocerServices) Sub(i int, i2 int) int {
	return i - i2
}
