package main

import (
	"github.com/heroku/go-love-notes/database"
	"github.com/heroku/go-love-notes/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

func init() {
	database.SetupDatabase()
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/notes", func(c *gin.Context) {
		notes, err := service.GetAllNotes()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": notes})
	})

	_ = router.Run(":" + port)
}
