package controller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ramailh/backend/fetch/props"
	"github.com/ramailh/backend/fetch/services"
)

func GetDataWithUSD(c *gin.Context) {
	datas, err := services.GetDataWithUSD()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": datas})
}

func GetAggregateData(c *gin.Context) {
	datas, err := services.GetAggregateData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": datas})
}

func Verify(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(props.Secret), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": token.Claims})
}
