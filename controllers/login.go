package controllers

import (
	"database/sql"
	"go-jwt-auth-api/initializers"
	"go-jwt-auth-api/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	if _, err := ctx.Cookie("Authorization"); err == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":   http.StatusUnauthorized,
			"mensagem": "Você já está autenticado!",
		})
		return
	}

	var login login

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   http.StatusBadRequest,
			"mensagem": "Tá errado esse JSON aí viu.",
		})
		return
	}

	dbUserRow := initializers.DB.QueryRow(
		"select name, username, password from users where username like ?",
		login.Username,
	)

	var dbUser User

	switch err := dbUserRow.Scan(&dbUser.Name, &dbUser.Username, &dbUser.Password); err {
	case sql.ErrNoRows:
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":   http.StatusUnauthorized,
			"mensagem": "Esse seu username pau no xibiu não existe não.",
		})
		return
	case nil:
		break
	default:
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"mensagem": "Ocorreu um erro ao entrar na sua conta.",
		})
		return
	}

	isTheSamePassword := utils.CheckPasswordHash(login.Password, dbUser.Password)
	if !isTheSamePassword {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":   http.StatusUnauthorized,
			"mensagem": "A senha tá errada viu mano.",
		})
		return
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": login.Username,
			"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
	jwtTokenString, err := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"mensagem": "Ocorreu um erro ao autenticar sua conta.",
		})
		return
	}

	ctx.SetCookie("Authorization", jwtTokenString, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"mensagem": "Você foi autenticado com sucesso!",
	})
}
