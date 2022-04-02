package middleware

import (
	"BCC_Intership_BuildUlang/helper"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func MiddlewareJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		bearerToken := strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := jwt.Parse(bearerToken, ekstractToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, helper.HelperMessage(false, err.Error(), nil))
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := uint(claims["id"].(float64))
			c.Set("loginberhasil", userId)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.HelperMessage(false, err.Error(), nil))
		}
	}
}

func ekstractToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte("klinikhewan"), nil
}
