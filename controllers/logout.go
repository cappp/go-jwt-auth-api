package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	if _, err := ctx.Cookie("Authorization"); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":   http.StatusUnauthorized,
			"mensagem": "Você não tá logado.",
		})
		return
	}
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"mensagem": "Você foi deslogado com sucesso.",
	})
}
