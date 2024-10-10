package helloworld

import "github.com/gin-gonic/gin"

type HelloEndpoint struct {
}

func InitHelloEndpoint() (*HelloEndpoint, error) {
	return &HelloEndpoint{}, nil
}

func (he *HelloEndpoint) RegisterEndpoint(router *gin.Engine) {
	router.GET("/helloworld", he.HelloWorld)
}

func (he *HelloEndpoint) HelloWorld(c *gin.Context) {
	c.JSON(200, "Hello World")
}
