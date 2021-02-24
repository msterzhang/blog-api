package repository

import "blog-api/api/models"

type TagRepository interface {
	Save(models.Tag) (models.Tag, error)
	FindAll(page int,size int) ([]models.Tag, error)
	FindByID(string) (models.Tag, error)
	UpdateByID(string, models.Tag) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string) ([]models.Tag, error)
}

