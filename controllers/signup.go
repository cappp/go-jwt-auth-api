package controllers

import (
	"database/sql"
	"go-jwt-auth-api/initializers"
	"go-jwt-auth-api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var user User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   http.StatusBadRequest,
			"mensagem": "Tá errado esse JSON aí viu.",
		})
		return
	}

	dbUsernameRow := initializers.DB.QueryRow(
		"select username from users where username like ?",
		user.Username,
	)

	var username string

	switch err := dbUsernameRow.Scan(&username); err {
	case sql.ErrNoRows:
		break
	case nil:
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   http.StatusBadRequest,
			"mensagem": "Já tem esse username aí rapaz!",
		})
		return
	default:
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"mensagem": "Ocorreu um erro ao criar sua conta.",
		})
		return
	}

	hashedUserPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"mensagem": "Ocorreu um erro ao criar sua conta.",
		})
		return
	}

	_, err = initializers.DB.Exec(
		"insert into users(name, username, password) values(?, ?, ?)",
		user.Name, user.Username, hashedUserPassword,
	)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"mensagem": "Ocorreu um erro ao criar sua conta.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"mensagem": "Sua conta foi criada.",
	})
}
