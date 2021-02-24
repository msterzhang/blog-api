package repository

import "blog-api/api/models"

type LinkRepository interface {
	Save(models.Link) (models.Link, error)
	FindAll(page int,size int) ([]models.Link, error)
	FindByID(string) (models.Link, error)
	UpdateByID(string, models.Link) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string) ([]models.Link, error)
}

