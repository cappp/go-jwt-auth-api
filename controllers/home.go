package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	if _, err := ctx.Cookie("Authorization"); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"mensagem": "E aí piva. Você não tá autenticado.",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"mensagem": "E aí piva. Você tá autenticado.",
	})
}
