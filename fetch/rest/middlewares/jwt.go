package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ramailh/backend/fetch/props"
)

const (
	tokenExp = 1 * time.Hour
)

type claims struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Timestamp int    `json:"timestamp"`
}

func JWTAuth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "HS512" {
			return nil, fmt.Errorf("signing algorithm not matched")
		}
		return []byte(props.Secret), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "status": false})
		c.Abort()
		return
	}

	if err = token.Claims.Valid(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "status": false})
		c.Abort()
		return
	}

	tokenByte, _ := json.Marshal(token.Claims)

	var claim claims
	json.Unmarshal(tokenByte, &claim)

	claimsTimestamp := time.Unix(0, int64(claim.Timestamp)*int64(1000000)).Add(tokenExp).Unix()
	if time.Now().Unix() > claimsTimestamp {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "token has expired", "status": false})
		c.Abort()
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token", "status": false})
		c.Abort()
	}
}

func JWTAuthAdmin(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "HS512" {
			return nil, fmt.Errorf("signing algorithm not matched")
		}
		return []byte(props.Secret), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "status": false})
		c.Abort()
		return
	}

	if err = token.Claims.Valid(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "status": false})
		c.Abort()
		return
	}

	tokenByte, _ := json.Marshal(token.Claims)

	var claim claims
	json.Unmarshal(tokenByte, &claim)

	claimsTimestamp := time.Unix(0, int64(claim.Timestamp)*int64(1000000)).Add(tokenExp).Unix()
	if time.Now().Unix() > claimsTimestamp {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "token has expired", "status": false})
		c.Abort()
		return
	}

	if claim.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized role", "status": false})
		c.Abort()
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token", "status": false})
		c.Abort()
	}
}
