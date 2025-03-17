package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func echoHandlerGin(c *gin.Context) {
	// Set Content-Type header to match the request
	if ct := c.Request.Header.Get("Content-Type"); ct != "" {
		c.Header("Content-Type", ct)
	}

	// Log the request method and URL
	log.Printf("Received %s request for %s", c.Request.Method, c.Request.URL.Path)

	// Copy request body to response
	if c.Request.Body != nil {
		defer c.Request.Body.Close()
		_, err := io.Copy(c.Writer, c.Request.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error reading request body")
			return
		}
	} else {
		// If no request body, return empty content with 200 OK
		c.Status(http.StatusOK)
	}
}

func setupRouter() *gin.Engine {
	// Create a default gin router with Logger and Recovery middleware
	router := gin.Default()

	// Handle all routes with the echo handler
	router.Any("/*path", echoHandlerGin)

	return router
}

func main() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Get the router
	router := setupRouter()

	// Start the server on port 8089
	addr := ":8089"
	log.Printf("Starting Gin echo server on port %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
