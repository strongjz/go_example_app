package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {

	go func() {
		routerAdmin := gin.Default()

		routerAdmin.GET("/admin", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Admin Sections",
			})
		})
		routerAdmin.Run(":8090")
	}()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Default Page",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	r.GET("/secret", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You should not see this",
		})
	})

	r.GET("/data", func(c *gin.Context) {
		db := CreateCon()

		err := db.Ping()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "DB is not connected",
			})

		}else{
			c.JSON(200, gin.H{
				"message": "Database Connected",
			})
		}
	})

	r.Run()
}

/*Create sql database connection*/
func CreateCon() *sql.DB {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v?sslmode=disable",user,pass,host,port)

	fmt.Printf("Database Connection String: %v \n",connStr)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	return db
}