package services

import (
	"github.com/mrinjamul/the-science-of-deduction/api/controllers"
	"github.com/mrinjamul/the-science-of-deduction/database"
	"github.com/mrinjamul/the-science-of-deduction/repository"
)

type Services interface {
	HealthCheckService() controllers.HealthCheck
	View() controllers.Template
}

type services struct {
	healthCheck controllers.HealthCheck
	view        controllers.Template
}

func (svc *services) HealthCheckService() controllers.HealthCheck {
	return svc.healthCheck
}

func (svc *services) View() controllers.Template {
	return svc.view
}

// NewServices initializes services
func NewServices() Services {
	db := database.GetDB()
	return &services{
		healthCheck: controllers.NewHealthCheck(),
		view: controllers.NewTemplate(
			repository.NewCaseFiles(db),
			repository.NewThread(db),
			repository.NewComment(db),
		),
	}
}
