package handlers

import (
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


func (h *Handlers) signIn(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func (h *Handlers) signUp(c gin.Context) {
	var input anki.User

	if err := c.BindJSON(&input); err != nil{
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// h.service.Autorization
}