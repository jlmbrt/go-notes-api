package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jlmbrt/go-notes-api/pkg/storage"
)

// Init storage
var DB *storage.Storage = storage.New()

func main() {
	// init Gin engine
	api := gin.New()

	// register route
	api.GET("/notes", getHandler)
	api.POST("/notes", postHandler)

	// start api
	log.Fatal(api.Run())
}

// Get Notes Handler
func getHandler(c *gin.Context) {
	var notes []storage.Note

	// check if query as tag parameter
	if tag := c.Query("tag"); tag != "" {
		// tag found ! load from DB with tag filtering
		notes = DB.FindWithTag(tag)
	} else {
		// no tag found, load all items from DB
		notes = DB.FinAll()
	}

	// Send result
	c.JSON(http.StatusOK, notes)
}

// Create Note Handler
func postHandler(c *gin.Context) {
	var newNote storage.Note

	// Bind Request.body into newNote
	// Gin Bind engine automatically controle JSON structure
	// based on tag in struct definition (see Note definition in ./pkg/storage/storage.go)
	if err := c.ShouldBindJSON(&newNote); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Create note in DB
	DB.Create(newNote)

	// Send Ok
	c.Status(http.StatusCreated)
}
