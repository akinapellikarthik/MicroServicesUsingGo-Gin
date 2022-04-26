package controller

import (
	"fmt"
	"testing"

	"GIN-DEMO/service"
)

//TODO: Need to mock the service class because main business logic is being called
//TODO: Need to understand below
type DemoServiceMock struct {
	HelloServiceFn func(custName string) string
}

//TODO: Need to understand below
func (mock DemoServiceMock) HelloService(custName string) string{
	return DemoServiceMock{}.HelloServiceFn(custName)
}

func TestHelloWorld(t *testing.T) {
	serviceMock := DemoServiceMock{}
	serviceMock.HelloServiceFn = func(custName string) string {
		return fmt.Sprintf("%s %s", "Test Method Hello",custName)
	}

	service.DemoService = serviceMock
	//res := service.HelloService("testCust")
	//if res != "Hello Customer testCust" {
	//	t.Error("Error")
	//}
}
