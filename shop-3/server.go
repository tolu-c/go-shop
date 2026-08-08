package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tolu-c/go-shop/handlers"
	"github.com/tolu-c/go-shop/middleware"
	"github.com/tolu-c/go-shop/service"
)

var (
	videoService service.VideoService   = service.New()
	videoHandler handlers.VideoHandlers = handlers.New(videoService)
)

func main() {
	s := gin.New()
	s.Static("/css", "./templates/css")
	s.LoadHTMLGlob("templates/*.html")

	s.Use(gin.Recovery(), middleware.Logger())

	apiRoutes := s.Group("/api")
	apiRoutes.Use(middleware.BasicAuth())

	apiRoutes.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoHandler.FindAll())
	})
	apiRoutes.POST("/videos", func(ctx *gin.Context) {
		err := videoHandler.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Video Created",
			})
		}

	})

	viewRoutes := s.Group("/view")
	viewRoutes.GET("/videos", videoHandler.ShowAll)
	
	s.Run(":4000")
}
