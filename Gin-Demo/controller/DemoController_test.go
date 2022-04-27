package controller

import (
	"GIN-DEMO/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	HelloService func(custName string) string
)

type demoServiceMock struct{}

func (m *demoServiceMock) HelloService(custName string) string {
	return HelloService(custName)
}

type someRes struct {
	Message string `json:"message"`
}

func TestHelloWorld(t *testing.T) {
	var res someRes
	HelloService = func(custName string) string {
		return custName
	}
	service.DemoService = &demoServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Params = gin.Params{
	{Key: "cust", Value: "Karthik"},
	}
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	HelloWorld(c)
	err := json.Unmarshal(response.Body.Bytes(), &res)
	if err != nil {
		return 
	}

	assert.Equal(t, "Karthik", res.Message)
}
