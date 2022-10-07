package routes

import (
	"embed"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/the-science-of-deduction/api/services"
)

// ViewsFs for static files
var ViewsFs embed.FS

var (
	StartTime time.Time
	BootTime  time.Duration
)

func InitRoutes(routes *gin.Engine) {
	// Initialize services
	svc := services.NewServices()

	// Serve the frontend
	// This will ensure that the web pages are served correctly

	routes.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./templates/layouts/index.html")
		} else {
			c.File("./templates/layouts/" + path.Join(dir, file))
		}
	})

	// Backend API
	// health check
	routes.GET("/api/health", func(c *gin.Context) {
		svc.HealthCheckService().HealthCheck(c, StartTime, BootTime)
	})
	//	api := routes.Group("/api")

}
