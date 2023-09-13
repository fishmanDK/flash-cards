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
		// api.GET("/main", h.GetGeneralInfo)

		category := api.Group("/categories")
		{
			category.GET("/", h.GetAllCategories)
			category.GET("/:categoryName", h.GetCategory)
			category.POST("/", h.CreateCategory)
			category.PATCH("/:categoryName", h.UpdateTitleCategory)
			category.DELETE("/:categoryName", h.DeleteCategory)

			question := category.Group(":categoryName/questions")
			{
				question.GET("/", h.GetAllQuestions)
				question.GET("/:questionName", h.GetQuestion)
				question.POST("/", h.CreateQuestion)
				question.PATCH("/:questionName", h.UpdateQustion)
				question.DELETE("/:questionName", h.DeleteQuestion)
			}
		}

	}

	return router
}
