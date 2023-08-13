package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
