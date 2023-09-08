package handlers

import (
	"github.com/fishmanDK/anki_telegram/internal/service"
	"github.com/gin-gonic/gin"
)

var (
	templatesPwd = "/Users/denissvecnikov/golang/anki/internal/template/*.html"
	frontPath    = "/Users/denissvecnikov/golang/anki/internal/front/"
)

type Handlers struct {
	Service service.Service
}

func NewHandlers(service service.Service) *Handlers {
	return &Handlers{
		Service: service,
	}
}

func (h *Handlers) InitRouts() *gin.Engine {

	router := gin.Default()

	router.Static("/assets/", frontPath)
	router.LoadHTMLGlob(templatesPwd)

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	auth := router.Group("/auth")
	{
		auth.GET("/signIn", h.signIn)
		auth.POST("/signIn", h.PsignIn)

		auth.GET("/signUp", h.signUp)
		auth.POST("/signUp", h.PsignUp)
		
	}


	api := router.Group("/api", h.Authentication)
	{
		api.GET("/main", h.GetGeneralInfo)
		api.GET("/category", h.GetCategories)
		// api.GET("/quesiton", h.)
	}

	return router
}

