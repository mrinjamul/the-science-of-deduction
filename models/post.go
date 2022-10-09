package models

import (
	"database/sql"
	"time"
)

// Post
type Post struct {
	Id          uint         `json:"id" gorm:"primary_key,unique,not null"`
	UserId      uint         `json:"user_id"`
	Type        string       `json:"type"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Content     string       `json:"content"`
	IsArchived  *bool        `json:"archived,omitempty"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:""`
	CreatedAt   time.Time    `json:"created_at" gorm:""`
	DeletedAt   sql.NullTime `json:"deleted_at" gorm:"index"`
}

// Comment
type Comment struct {
	Id        uint         `json:"id" gorm:"primary_key,unique,not null"`
	UserId    uint         `json:"user_id"`
	PostId    uint         `json:"post_id"`
	ThreadId  uint         `json:"thread_id"`
	Type      string       `json:"type"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	CreatedAt time.Time    `json:"created_at" gorm:""`
	UpdatedAt time.Time    `json:"updated_at" gorm:""`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"index"`
}

// Thread
type Thread struct {
	Id        uint         `json:"id" gorm:"primary_key,unique,not null"`
	UserId    uint         `json:"user_id"`
	ThreadId  uint         `json:"thread_id"`
	Message   string       `json:"message"`
	CreatedAt time.Time    `json:"created_at" gorm:""`
	UpdatedAt time.Time    `json:"updated_at" gorm:""`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"index"`
}

// Tag
type Tag struct {
	Id          uint         `json:"id" gorm:"primary_key,unique,not null"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"not null"`
	DeletedAt   sql.NullTime `json:"deleted_at" gorm:"index"`
}

// Event
type Event struct {
	Id          uint         `json:"id" gorm:"primary_key,unique,not null"`
	UserId      uint         `json:"user_id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at" gorm:""`
	DeletedAt   sql.NullTime `json:"deleted_at" gorm:"index"`
}
