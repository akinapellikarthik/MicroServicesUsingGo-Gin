package service

import "fmt"

//Step 1: Create a interface
type IDemoService interface {
	HelloService(custName string) string
}

//Step 2: Implement the interface
type demoServiceImpl struct {}

var (
	DemoService IDemoService = &demoServiceImpl{}
)

func (service *demoServiceImpl)HelloService(custName string) string {
	fmt.Println("Inside the service function....")
	return fmt.Sprintf("%s %s", "Hello Customer", custName)
}
