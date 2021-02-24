package repository

import "blog-api/api/models"

type PostRepository interface {
	Save(models.Post) (models.Post, error)
	FindAll(page int,size int) ([]models.Post, error)
	FindByID(string) (models.Post, error)
	UpdateByID(string, models.Post) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string) ([]models.Post, error)
}

