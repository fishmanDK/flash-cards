package handlers

import (
	// "html/template"
	"log"
	"net/http"

	anki "github.com/fishmanDK/anki_telegram"
	"github.com/gin-gonic/gin"
)

// type User struct {
// 	Id       int
// 	Username string
// 	Email    string
// 	Password string
// }

// type Details struct {
// 	Email    string
// 	Password string
// }

func (h *Handlers) signIn(c *gin.Context) {
	c.HTML(200, "entrance.html", gin.H{})
}

func (h *Handlers) PsignIn(c *gin.Context) {
	var input anki.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Autorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	log.Println(input)
	c.String(200, "succses")
}

func (h *Handlers) signUp(c *gin.Context) {
	c.Header("Cache-Control", "no-cache")

	c.HTML(http.StatusOK, "registration.html", gin.H{})
}

// func (h *Handlers) PsignUp(c *gin.Context) {
// 	var input anki.User
// 	if err := c.BindJSON(&input); err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
// 	report, err := h.service.Validate.ValidateRegistration(input)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": report,
// 		})
// 		return
// 	}
// 	err = h.service.Autorization.CreateUser(input)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}
//  	c.Redirect(302, "/auth/signIn")
// }

func (h *Handlers) PsignUp(c *gin.Context) {

	var input anki.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	report, err := h.service.Validate.ValidateRegistration(input)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errors": report,
		})
		return
	}

	err = h.service.Autorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

