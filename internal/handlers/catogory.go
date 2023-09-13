package handlers

import (
	"net/http"

	anki "github.com/fishmanDK/anki_telegram"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type UpdateTitle struct {
	OldTitle string
	NewTitle string `json:"new-title"`
}

func (h *Handlers) GetAllCategories(c *gin.Context) {
	userId, errUserId := c.Get("user_id")
	if !errUserId {
		NewErrorResponse(c, http.StatusBadRequest, "ошибка при получении id пользователя (GetAllCategories)")
		return
	}

	userCategories, _ := h.Service.Category.GetCategories(userId.(string))
	c.JSON(http.StatusOK, bson.M{
		"category-tytles": userCategories.TitleCategories,
	})
}

func (h *Handlers) GetCategory(c *gin.Context) {
	userId, errUserId := c.Get("user_id")
	if !errUserId {
		NewErrorResponse(c, http.StatusBadRequest, "ошибка при получении id пользователя (GetCategory)")
		return
	}

	categoryName := c.Param("categoryName")

	category, err := h.Service.Category.GetCategoryById(userId.(string), categoryName)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	categoryJSON := bson.M{
		"tytle":     category.Title,
		"questions": category.QuestionsText,
	}

	c.JSON(http.StatusOK, categoryJSON)
}

func (h *Handlers) CreateCategory(c *gin.Context) {
	userId, errUserId := c.Get("user_id")

	if !errUserId {
		NewErrorResponse(c, http.StatusBadRequest, "ошибка при получении id пользователя (CreateCategory)")
		return
	}

	var input anki.CreateCategory
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	categoryId, err := h.Service.Category.CreateCategory(userId.(string), input.Title)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, bson.M{
		"created-id": categoryId,
	})
}

func (h *Handlers) UpdateTitleCategory(c *gin.Context) {
	userId, errUserId := c.Get("user_id")

	if !errUserId {
		NewErrorResponse(c, http.StatusBadRequest, "ошибка при получении id пользователя (UpdateTitleCategory)")
		return
	}

	var input UpdateTitle
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	categoryName := c.Param("categoryName")
	input.OldTitle = categoryName

	err := h.Service.Category.UpdateTitle(userId.(string), input.OldTitle, input.NewTitle)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusCreated, bson.M{
		"status": "изменения прошли успешно",
	})
}

func (h *Handlers) DeleteCategory(c *gin.Context) {
	userId, errUserId := c.Get("user_id")

	if !errUserId {
		NewErrorResponse(c, http.StatusBadRequest, "ошибка при получении id пользователя (UpdateTitleCategory)")
		return
	}

	categoryName := c.Param("categoryName")

	err := h.Service.Category.DeleteCategory(userId.(string), categoryName)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, bson.M{
		"status": "удалени прошло успешно",
	})
}
