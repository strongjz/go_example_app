package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	r := gin.Default()

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

		fmt.Printf("Error: %v/n", err)

		if err != nil {
			fmt.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "db is not connected",
			})

		}else{
			c.JSON(200, gin.H{
				"message": "Database Connected",
			})
		}
	})

	r.Run()
}

/*Create mysql connection*/
func CreateCon() *sql.DB {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user,pass,host,port)

	fmt.Printf("Database Connection String: %v",conn)

	db, err := sql.Open("mysql", conn)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()


	return db
}