package repository

import "blog-api/api/models"

type CommentRepository interface {
	Save(models.Comment) (models.Comment, error)
	FindAll(page int,size int) ([]models.Comment, error)
	FindByID(string) (models.Comment, error)
	UpdateByID(string, models.Comment) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string) ([]models.Comment, error)
}

