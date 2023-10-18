package main

import (
	"net/http"

	docs "golang-swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Ping API
//	@version		1.0
//	@description	This is a sample API.

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Title = "ping API"
	r.GET("/ping", ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type ReturnModel struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Property int    `json:"property"`
}

// PingExample godoc
//
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ReturnModel	"ok"
//	@Router			/ping [get]
func ping(c *gin.Context) {
	resp := ReturnModel{Id: "2241441", Name: "wfffqw", Property: 12244555}
	c.JSON(http.StatusOK,
		&resp)
}
