package db

import "github.com/jorgeAM/bloGrpc/models"

// Handler handles database operations
type Handler interface {
	NewBlog(blog models.Blog) (*models.Blog, error)
	ReadBlog(id string) (*models.Blog, error)
	ListBlogs() ([]*models.Blog, error)
	UpdateBlog(blog models.Blog) (*models.Blog, error)
	DeleteBlog(id string) error
}
