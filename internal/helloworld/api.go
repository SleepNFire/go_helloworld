package helloworld

import "github.com/gin-gonic/gin"

type HelloEndpoint struct {
}

func InitHelloEndpoint() (*HelloEndpoint, error) {
	return &HelloEndpoint{}, nil
}

func (he *HelloEndpoint) RegisterEndpoint(router *gin.Engine) {
	router.GET("/turing", he.HelloWorld)
	router.GET("/alan-turing", he.HelloWorld)
	router.GET("/AlanTuring", he.HelloWorld)
	router.GET("/alanturing", he.HelloWorld)
	router.GET("/reponse", he.NotResponse)
}

func (he *HelloEndpoint) HelloWorld(c *gin.Context) {
	c.JSON(200, "Félicitations, vous avez terminé l'initiation ! J'espère que le jeu de piste vous a plu.")
}

func (he *HelloEndpoint) NotResponse(c *gin.Context) {
	c.JSON(200, "Il faut changer le mot reponse par le bon nom")
}
