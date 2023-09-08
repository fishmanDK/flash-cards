package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) Authentication(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "пользователь не авторизирован")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "ошибка авторизации")
		return
	}

	id, err := h.Service.Autorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "ошибка парсинга токена")
		return
	}
	c.Set("user_id", id)
	c.Next()
}
