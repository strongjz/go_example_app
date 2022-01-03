package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/strongjz/go_example_app/app"
	"os"
)

func main() {

	go func() {
		adPort := os.Getenv("AD_PORT")
		routerAdmin := gin.Default()

		routerAdmin.GET("/admin", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Admin Sections",
			})
		})
		routerAdmin.Run(adPort)
	}()

	app := app.New()
	fmt.Println("Starting App")
	app.Engine()
	app.Start()
}
