package handlers

import (
	"log"
	"net/http"

	anki "github.com/fishmanDK/anki_telegram"
	"github.com/gin-gonic/gin"
)

type User struct{
	Id 		 int
	Username string
	Email 	 string
	Password string
}

type Details struct{
	Email    string
	Password string
}

func (h *Handlers) signIn(c *gin.Context) {
	c.HTML(200, "entrance.html", gin.H{})
}

func (h *Handlers) PsignIn(c *gin.Context) {
	var input anki.User
	
	if err := c.BindJSON(&input); err != nil{
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}


	// err := h.service.Autorization.CreateUser(input)
	// if err != nil{
	// 	NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// }

	log.Println(input)
	c.String(200, "succses")
}

func (h *Handlers) signUp(c *gin.Context) {
	c.HTML(http.StatusOK, "registration.html", gin.H{})
}


func (h *Handlers) PsignUp(c *gin.Context) {
	var input anki.User

	if err := c.BindJSON(&input); err != nil{
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Autorization.CreateUser(input)
	if err != nil{
		c.HTML(http.StatusOK, "index.html", gin.H{})
		return
	}

	c.String(200, "succses")
}