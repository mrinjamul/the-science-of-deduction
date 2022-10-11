package repository

import (
	"github.com/mrinjamul/the-science-of-deduction/models"
	"gorm.io/gorm"
)

type Comment interface {
	// CreateComment creates a new Comment
	CreateComment(comment *models.Comment) error
	// GetCommentByID returns a Comment
	GetCommentByID(id uint) (models.Comment, error)
	// GetComments returns a list of Comments
	GetComments() ([]models.Comment, error)
	// GetCommentsByThreadID returns a list of Comments
	GetCommentsByThreadID(id uint) ([]models.Comment, error)
	// UpdateComment updates a Comment
	UpdateComment(comment *models.Comment) error
	// DeleteComment deletes a Comment
	DeleteComment(id uint) error
}

type comment struct {
	db gorm.DB
}

// CreateComment creates a new Comment
func (p *comment) CreateComment(comment *models.Comment) error {
	return p.db.Create(comment).Error
}

// GetCommentByID returns a Comment
func (p *comment) GetCommentByID(id uint) (models.Comment, error) {
	var comment models.Comment
	err := p.db.First(&comment, uint(id)).Error
	return comment, err
}

// GetComments returns a list of Comments
func (p *comment) GetComments() ([]models.Comment, error) {
	var comments []models.Comment
	err := p.db.Order("created_at desc").Find(&comments).Error
	return comments, err
}

// GetCommentsByThreadID returns a list of Comments
func (p *comment) GetCommentsByThreadID(id uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := p.db.Order("created_at").Where("thread_id = ?", id).Find(&comments).Error
	return comments, err
}

// UpdateComment updates a Comment
func (p *comment) UpdateComment(comment *models.Comment) error {
	return p.db.Save(comment).Error
}

// DeleteComment deletes a Comment
func (p *comment) DeleteComment(id uint) error {
	return p.db.Delete(&models.Comment{}, uint(id)).Error
}

// NewComment returns a new repository
func NewComment(db *gorm.DB) Comment {
	return &comment{
		db: *db,
	}
}
