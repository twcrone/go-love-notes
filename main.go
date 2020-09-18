package main

import (
	"github.com/heroku/go-love-notes/database"
	"github.com/heroku/go-love-notes/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/itsjamie/gin-cors"
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
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
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
