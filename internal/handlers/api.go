package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handlers) GetGeneralInfo(c *gin.Context) {
	id, _ := c.Get("user_id")

	// opts := nil

	userData, _ := h.Service.Api.GetDataAboutUser(id.(string))
	c.JSON(http.StatusOK, userData)
}

func (h *Handlers) GetCategories(c *gin.Context){
	id, _ := c.Get("user_id")

	// opts := []string{"Categories"}

	userCategories, _ := h.Service.Api.GetCategories(id.(string))

	c.JSON(http.StatusOK, userCategories)
}

// type Element struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }
