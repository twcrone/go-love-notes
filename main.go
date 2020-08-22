package main

import (
	"encoding/json"
	"github.com/heroku/go-getting-started/database"
	"github.com/heroku/go-getting-started/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
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
		notesJson, err := json.Marshal(notes)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, notesJson)
	})

	router.Run(":" + port)
}
