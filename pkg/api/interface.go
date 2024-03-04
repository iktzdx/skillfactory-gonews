package api

import (
	"github.com/iktzdx/skillfactory-gonews/internal/app/posts"
)

type BoundaryPort interface {
	// Create(post Post) (Post, error)
	GetPostByID(id string) (posts.Post, error)
	// List(opts SearchOpts) (Posts, error)
	// Update(post Post) (Post, error)
	// Delete(id string) error
}