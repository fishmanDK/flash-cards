package handlers

import (
	"net/http"

	anki "github.com/fishmanDK/anki_telegram"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handlers) GetAllQuestions(c *gin.Context) {
	userId, errUserId := c.Get("user_id")
	if !errUserId {
		NewErrorResponse(c, http.StatusBadRequest, "ошибка при получении id пользователя (GetAllCategories)")
		return
	}

	categoryName := c.Param("categoryName")

	questions, err := h.Service.Question.GetAllQuestions(userId.(string), categoryName)
	if err != nil{
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, questions)
}

func (h *Handlers) GetQuestion(c *gin.Context) {
	userId, errUserId := c.Get("user_id")
	if !errUserId {
		NewErrorResponse(c, http.StatusBadRequest, "ошибка при получении id пользователя (GetQuestion)")
		return
	}
	questionName := c.Param("questionName")
	res, err := h.Service.Question.GetQuestion(userId.(string), questionName)
	if err != nil{
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handlers) CreateQuestion(c *gin.Context) {
	var input anki.CreateQuestion
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, errUserId := c.Get("user_id")
	if !errUserId {
		NewErrorResponse(c, http.StatusBadRequest, "ошибка при получении id пользователя (GetAllCategories)")
		return
	}

	cateoryName := c.Param("categoryName")
	

	res, err := h.Service.Question.CreateQuestion(userId.(string), cateoryName, input)
	if err != nil{
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handlers) UpdateQustion(c *gin.Context) {
	var input anki.UpdateQuestion
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, errUserId := c.Get("user_id")
	if !errUserId {
		NewErrorResponse(c, http.StatusBadRequest, "ошибка при получении id пользователя (GetAllCategories)")
		return
	}

	cateoryName := c.Param("categoryName")
	questionName := c.Param("questionName")

	err := h.Service.Question.UpdateQustion(userId.(string), cateoryName, questionName, input)
	if err != nil{
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, bson.M{
		"status": "изменения прошли успешно",
	})
}

func (h *Handlers) DeleteQuestion(c *gin.Context) {
	userId, errUserId := c.Get("user_id")
	if !errUserId {
		NewErrorResponse(c, http.StatusBadRequest, "ошибка при получении id пользователя (GetAllCategories)")
		return
	}

	categoryName := c.Param("categoryName")
	questionName := c.Param("questionName")

	err := h.Service.DeleteQuestion(userId.(string), categoryName, questionName)
	if err != nil{
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, bson.M{
		"status": "удаление прошло успешно",
	})
}
