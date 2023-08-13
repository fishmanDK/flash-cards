package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) mainPage(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", gin.H{})
}