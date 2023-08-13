package handlers

import (
	"log"
	"net/http"

	anki "github.com/fishmanDK/anki_telegram"
	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int
	Username string
	Email    string
	Password string
}

type Details struct {
	Email    string
	Password string
}

func (h *Handlers) signIn(c *gin.Context) {
	c.HTML(200, "entrance.html", gin.H{})
}

func (h *Handlers) PsignIn(c *gin.Context) {
	var input anki.User

	if err := c.BindJSON(&input); err != nil {
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

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	report, err := h.service.Validate.ValidateRegistration(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": report,
		})
		return
	}

	err = h.service.Autorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Redirect to another page in case of successful processing of the POST request
	c.Redirect(http.StatusOK, "/auth/signIn")
}

// func (h *Handlers) PsignUp(c *gin.Context) {
// 	var input anki.User

// 	if err := c.BindJSON(&input); err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	_, err := h.service.Validate.ValidateRegistration(input)
// 	log.Println("report")
// 	if err != nil{
// 		c.JSON(http.StatusBadRequest, "report")
// 		return
// 	}

// 	err = h.service.Autorization.CreateUser(input)
// 	if err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	// В случае успешной обработки POST-запроса, можно выполнить редирект на другую страницу
// 	c.Redirect(http.StatusOK, "/auth/signIn")
// }

// func (h *Handlers) PsignUp(c *gin.Context) {
// 	var input anki.User

// 	if err := c.BindJSON(&input); err != nil {
// 		NewErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	err := h.service.Autorization.CreateUser(input)
// 	if err != nil {
// 		report, _ := h.service.Autorization.db.ValidateRegistration(input.Email, input.Username, input.Password, input.RepeatPassword)
// 		c.HTML(http.StatusOK, "registration.html", gin.H{
// 			"errorMessage": "Произошла ошибка",
// 			"data": gin.H{
// 				"Report": report,
// 			},
// 		})
// 		return
// 	}

// 	// В случае успешной обработки POST-запроса, можно выполнить редирект на другую страницу
// 	c.Redirect(http.StatusSeeOther, "/auth/signIn")
// }
