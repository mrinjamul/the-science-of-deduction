package services

import (
	"github.com/mrinjamul/the-science-of-deduction/api/controllers"
)

type Services interface {
	HealthCheckService() controllers.HealthCheck
}

type services struct {
	healthCheck controllers.HealthCheck
}

func (svc *services) HealthCheckService() controllers.HealthCheck {
	return svc.healthCheck
}

// NewServices initializes services
func NewServices() Services {
	// db := database.GetDB()
	return &services{
		healthCheck: controllers.NewHealthCheck(),
	}
}
