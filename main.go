package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/the-science-of-deduction/api/routes"
)

//go:embed templates/layouts/*
var files embed.FS

var (
	// startTime is the time wshen the server starts
	startTime time.Time = time.Now()
)

func main() {
	// Get port from env
	port := ":3000"
	_, present := os.LookupEnv("PORT")
	if present {
		port = ":" + os.Getenv("PORT")

	}

	// Set the router as the default one shipped with Gin
	server := gin.Default()
	templ := template.Must(template.New("").ParseFS(files, "templates/layouts/*.html", "templates/layouts/*.html"))
	server.SetHTMLTemplate(templ)
	static, err := fs.Sub(files, "templates/static")
	if err != nil {
		panic(err)
	}
	media, err := fs.Sub(files, "templates/media")
	if err != nil {
		panic(err)
	}
	server.StaticFS("/static", http.FS(static))
	server.StaticFS("/media", http.FS(media))

	// Initialize the routes
	routes.StartTime = startTime
	routes.InitRoutes(server)
	routes.BootTime = time.Since(startTime)
	// Start and run the server
	log.Fatal(server.Run(port))
}
