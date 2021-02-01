package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ramailh/backend/fetch/rest/controller"
	"github.com/ramailh/backend/fetch/rest/middlewares"
)

func NewRouter() *gin.Engine {
	rtr := gin.Default()

	fetch := rtr.Group("/fetch")
	{
		fetch.GET("/with-usd", middlewares.JWTAuth, controller.GetDataWithUSD)
		fetch.GET("/verify-token", middlewares.JWTAuth, controller.Verify)
		fetch.GET("/aggregate", middlewares.JWTAuthAdmin, controller.GetAggregateData)
	}

	return rtr
}
