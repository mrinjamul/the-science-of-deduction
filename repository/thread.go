package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/the-science-of-deduction/models"
	"gorm.io/gorm"
)

type Thread interface {
	// CreateThread creates a new Thread
	CreateThread(ctx *gin.Context, thread *models.Thread) error
	// GetThreadByID returns a Thread
	GetThreadByID(ctx *gin.Context, id int) (models.Thread, error)
	// GetThreads returns a list of Threads
	GetThreads(ctx *gin.Context) ([]models.Thread, error)
	// UpdateThread updates a Thread
	UpdateThread(ctx *gin.Context, thread *models.Thread) error
	// DeleteThread deletes a Thread
	DeleteThread(ctx *gin.Context, id int) error
}

type thread struct {
	db gorm.DB
}

// CreateThread creates a new Thread
func (p *thread) CreateThread(ctx *gin.Context, thread *models.Thread) error {
	return p.db.Create(thread).Error
}

// GetThreadByID returns a Thread
func (p *thread) GetThreadByID(ctx *gin.Context, id int) (models.Thread, error) {
	var thread models.Thread
	err := p.db.First(&thread, uint(id)).Error
	return thread, err
}

// GetThreads returns a list of Threads
func (p *thread) GetThreads(ctx *gin.Context) ([]models.Thread, error) {
	var threads []models.Thread
	err := p.db.Order("created_at desc").Find(&threads).Error
	return threads, err
}

// UpdateThread updates a Thread
func (p *thread) UpdateThread(ctx *gin.Context, thread *models.Thread) error {
	return p.db.Save(thread).Error
}

// DeleteThread deletes a Thread
func (p *thread) DeleteThread(ctx *gin.Context, id int) error {
	return p.db.Delete(&models.Thread{}, uint(id)).Error
}

// NewThread returns a new repository
func NewThread(db *gorm.DB) Thread {
	return &thread{
		db: *db,
	}
}
