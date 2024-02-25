package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Private(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"mensagem": "Você descobriu a imagem perdida de Naldo, Kanye West e Will Smith: https://i.imgur.com/FSvDetw.png",
	})
}
