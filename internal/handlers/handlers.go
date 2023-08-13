package handlers

import (

	"github.com/fishmanDK/anki_telegram/internal/service"
	"github.com/gin-gonic/gin"
)

var(
	templatesPwd = "/Users/denissvecnikov/golang/anki/internal/template/*.html"
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
		auth.POST("/signIn", h.PsignIn)

		auth.GET("/signUp", h.signUp)
		auth.POST("/signUp", h.PsignUp)
	}
	
	api := router.Group("/api")
	{
		api.GET("/main", h.mainPage)
	}
	return router
}