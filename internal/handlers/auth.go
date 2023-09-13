package handlers

import (
	"log"
	"net/http"

	anki "github.com/fishmanDK/anki_telegram"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func (h *Handlers) signIn(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.HTML(http.StatusOK, "entrance.html", nil)
}

func (h *Handlers) PsignIn(c *gin.Context) {
	var input User

	err := c.BindJSON(&input)
	if err != nil {
		log.Println(input)
		NewErrorResponse(c, http.StatusBadRequest, "поля при аутентификации не заполненны")
		return
	}

	accessToken, err := h.Service.Authentication(input.Email, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, bson.M{
		"accessToken": accessToken,
	})
}

func (h *Handlers) signUp(c *gin.Context) {
	c.Header("Cache-Control", "no-cache")

	c.HTML(http.StatusOK, "registration.html", gin.H{})
}

func (h *Handlers) PsignUp(c *gin.Context) {

	var input anki.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	report, err := h.Service.Validate.ValidateRegistration(input)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errors": report,
		})
		return
	}

	err = h.Service.Autorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
