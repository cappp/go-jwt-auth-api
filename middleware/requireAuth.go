package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	jwtTokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":   http.StatusUnauthorized,
			"mensagem": "Você precisa estar autenticado para entrar aqui.",
		})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, _ := jwt.Parse(jwtTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["HS256"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":   http.StatusInternalServerError,
				"mensagem": "O seu token expirou, logue novamente.",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":   http.StatusInternalServerError,
			"mensagem": "O token dado não é daqui não satanico.",
		})
		ctx.AbortWithStatus(http.StatusUnauthorized)

	}
}
