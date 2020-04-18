package app

import "github.com/gin-gonic/gin"

type App struct {
	router            *gin.Engine
}

func New() *App {

	return &App{}
}

func (a *App) Start() *gin.Engine {

		// set server mode
		gin.SetMode(gin.DebugMode)

		r := gin.New()

		// Global middleware
		r.Use(gin.Logger())
		r.Use(gin.Recovery())

		r.GET("/", rootHandler)
		r.GET("/ping", pingHandler)
		r.GET("/data", dataHandler)
		r.GET("/secret", secretHandler)
		r.GET("/host", hostHandler)
		r.GET("/external", externalHandler)

		return r
}
