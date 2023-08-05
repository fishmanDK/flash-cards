package handlers

import (

	"github.com/fishmanDK/anki_telegram/internal/service"
	"github.com/gin-gonic/gin"
)

var(
	templatesPwd = "/Users/denissvecnikov/golang/anki/internal/templates/*.html"
	frontPath = "/Users/denissvecnikov/golang/anki/internal/front/"
)

type Handlers struct{
	service service.Service
}

func NewHandlers(service service.Service) *Handlers{
	return &Handlers{
		service: service,
	}
}



func (h *Handlers) InitRouts() *gin.Engine{
	router := gin.New()	

	router.Static("/assets/", frontPath)
	router.LoadHTMLGlob(templatesPwd)

	auth := router.Group("/auth")
	{
		auth.GET("/signIn", h.signIn)
	}
	
	return router
}