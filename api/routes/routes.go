package routes

import (
	"embed"
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

	// Home Page
	routes.GET("/", func(c *gin.Context) {
		svc.View().Index(c)
	})
	// Case Files
	routes.GET("/case-files", func(c *gin.Context) {
		svc.View().CaseFiles(c)
	})
	// Forum
	routes.GET("/forum", func(c *gin.Context) {
		svc.View().Forum(c)
	})
	// Hidden Messages
	routes.GET("/hidden-messages", func(c *gin.Context) {
		svc.View().Index(c)
	})
	routes.NoRoute(func(c *gin.Context) {
		svc.View().NotFound(c)
	})

	// Backend API
	// health check
	routes.GET("/api/health", func(c *gin.Context) {
		svc.HealthCheckService().HealthCheck(c, StartTime, BootTime)
	})
	//	api := routes.Group("/api")

}
